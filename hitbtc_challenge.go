package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type Currency struct {
	ID          string `json:"id"`
	FullName    string `json:"fullName"`
	Ask         string `json:"ask"`
	Bid         string `json:"bid"`
	Last        string `json:"last"`
	Open        string `json:"open"`
	Low         string `json:"low"`
	High        string `json:"high"`
	FeeCurrency string `json:"feeCurrency"`
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// Store the currency data in a concurrent-safe map
var currencyData sync.Map

// func (cur *Currency, id, fullname) UpdateSymbolData() {
// 	cur.ID = id
// 	cur.FullName = fullname
// }

// func (cur *Currency, ask, bid, last, low, high, open) UpdatetickData() {
// 	cur.Ask = ask
// 	cur.Bid = bid
// 	cur.Last = last
// 	cur.Low = low
// 	cur.High = high
// 	cur.Open = open
// }

func main() {
	go func() {
		// // Connect to the HitBTC WebSocket
		// c, _, err := websocket.DefaultDialer.Dial("wss://api.hitbtc.com/api/3/ws", nil)
		// if err != nil {
		// 	panic(err)
		// }
		dialer := websocket.Dialer{
			HandshakeTimeout: 45 * time.Second,
		}

		c, resp, err := dialer.Dial("wss://api.hitbtc.com/api/2/ws/public", nil)
		if err != nil {
			fmt.Println("Error during handshake:", err)
			if resp != nil {
				fmt.Println("HTTP Response Status:", resp.Status)
				fmt.Println("HTTP Response Headers:", resp.Header)
			}
			return
		}
		// fmt.Println("handshake succesful")
		defer c.Close()

		// Subscribe to the ticker for each currency pair
		for _, pair := range []string{"BTCUSD", "ETHBTC"} {
			subscribeMessage := map[string]interface{}{
				"method": "subscribeTicker",
				"params": map[string]string{
					"symbol": pair,
				},
				// "id": 123,
			}
			c.WriteJSON(subscribeMessage)
		}

		// Handle incoming messages
		for {
			_, message, err := c.ReadMessage()
			// fmt.Println(message)
			if err != nil {
				panic(err)
			}

			// Parse the JSON message into a Currency object
			var result map[string]interface{}
			json.Unmarshal(message, &result)
			fmt.Println(result)
			params, ok := result["params"].(map[string]interface{})
			if !ok {
				// Log the error, ignore this message, or handle the error in some other way.
				fmt.Println("Failed to parse params")
				continue // Skip to the next iteration of the loop.
			}
			currency := Currency{
				ID:   params["symbol"].(string),
				Ask:  params["ask"].(string),
				Bid:  params["bid"].(string),
				Last: params["last"].(string),
				Open: params["open"].(string),
				Low:  params["low"].(string),
				High: params["high"].(string),
				// FeeCurrency: params["feeCurrency"].(string),
			}

			// Store the currency data in the map
			currencyData.Store(currency.ID, currency)
		}
	}()

	r := gin.Default()
	r.GET("/currency/:symbol", handleCurrencySymbol)
	r.GET("/currency/all", handleCurrencyAll)

	r.Run() // listens and serves on 0.0.0.0:8080
}

func handleCurrencySymbol(c *gin.Context) {
	symbol := c.Param("symbol")
	// fmt.Println(symbol)

	// symboldata := map[string]interface{}{
	// 	"method": "getSymbol",
	// 	"params": map[string]string{
	// 		"symbol": symbol,
	// 	},
	// 	"id": 123,
	// }
	// c.WriteJSON(symboldata)

	// Fetch the currency data from the map
	value, ok := currencyData.Load(symbol)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid symbol"})
		return
	}

	currency := value.(Currency)
	c.JSON(http.StatusOK, currency)
}

func handleCurrencyAll(c *gin.Context) {
	currencies := make([]Currency, 0)

	currencyData.Range(func(key, value interface{}) bool {
		currencies = append(currencies, value.(Currency))
		return true
	})

	c.JSON(http.StatusOK, currencies)
}
