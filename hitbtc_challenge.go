package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
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

var currencies map[string]Currency
var mutex = &sync.Mutex{}

func main() {
	currencies = make(map[string]Currency)
	// Supported symbols
	// symbols := []string{"BTCUSD", "ETHBTC"}

	u := url.URL{Scheme: "wss", Host: "api.hitbtc.com", Path: "/api/2/ws"}

	// Initialize and maintain WebSocket connection
	go func() {
		var feeCurrencyy string
		var idd string
		var fullNamee string
		c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
		fmt.Printf("c = %T\n", c)

		if err != nil {
			log.Fatal("dial:", err)
		}
		defer c.Close()

		// get symbol
		symbolMessage := map[string]interface{}{
			"method": "getSymbol",
			"params": map[string]string{
				"symbol": "ETHBTC",
			},
			"id": time.Now().Unix(),
		}
		c.WriteJSON(symbolMessage)

		_, r1message, r1err := c.ReadMessage()
		if r1err != nil {
			log.Println("read:", err)
			return
		}

		var r1 map[string]interface{}
		json.Unmarshal(r1message, &r1)
		res, ok := r1["result"].(map[string]interface{})
		if !ok {
			fmt.Println("failed to parse params 1 eth")
		}
		if res != nil {
			mutex.Lock()
			feeCurrencyy = res["feeCurrency"].(string)
			mutex.Unlock()
		}
		//*******************************************************************
		// get currency
		currencyMessage := map[string]interface{}{
			"method": "getCurrency",
			"params": map[string]string{
				"currency": res["baseCurrency"].(string),
			},
			"id": time.Now().Unix(),
		}
		c.WriteJSON(currencyMessage)

		_, r2message, r2err := c.ReadMessage()
		if r2err != nil {
			log.Println("read:", err)
			return
		}
		var r2 map[string]interface{}
		json.Unmarshal(r2message, &r2)
		res2, ok := r2["result"].(map[string]interface{})
		if !ok {
			fmt.Println("failed to parse params 2 eth")
		}
		if res != nil {
			mutex.Lock()
			idd = res2["id"].(string)
			fullNamee = res2["fullName"].(string)
			mutex.Unlock()
		}
		//**********************************************************************
		// Subscribe to the market data for each symbol
		subscriptionMessage := map[string]interface{}{
			"method": "subscribeTicker",
			"params": map[string]string{
				"symbol": "ETHBTC",
			},
			"id": time.Now().Unix(),
		}
		c.WriteJSON(subscriptionMessage)

		// Store real-time data in memory
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}

			var result map[string]interface{}
			json.Unmarshal(message, &result)

			params, ok := result["params"].(map[string]interface{})
			if !ok {
				fmt.Println("failed to parse params 3 eth")
				continue
			}
			if params != nil {
				mutex.Lock()
				currencies[params["symbol"].(string)] = Currency{
					ID:          idd,
					FullName:    fullNamee,
					Ask:         params["ask"].(string),
					Bid:         params["bid"].(string),
					Last:        params["last"].(string),
					Open:        params["open"].(string),
					Low:         params["low"].(string),
					High:        params["high"].(string),
					FeeCurrency: feeCurrencyy,
				}
				mutex.Unlock()
			}
		}
	}()

	//****************************************************************************************************************************************
	go func() {
		var feeCurrencyy string
		var idd string
		var fullNamee string
		c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
		fmt.Printf("c = %T\n", c)

		if err != nil {
			log.Fatal("dial:", err)
		}
		defer c.Close()

		// get symbol
		symbolMessage := map[string]interface{}{
			"method": "getSymbol",
			"params": map[string]string{
				"symbol": "BTCUSD",
			},
			"id": time.Now().Unix(),
		}
		c.WriteJSON(symbolMessage)

		_, r1message, r1err := c.ReadMessage()
		if r1err != nil {
			log.Println("read:", err)
			return
		}

		var r1 map[string]interface{}
		json.Unmarshal(r1message, &r1)
		res, ok := r1["result"].(map[string]interface{})
		if !ok {
			fmt.Println("failed to parse params 1")
		}
		if res != nil {
			mutex.Lock()
			feeCurrencyy = res["feeCurrency"].(string)
			mutex.Unlock()
		}
		//*******************************************************************
		// get currency
		currencyMessage := map[string]interface{}{
			"method": "getCurrency",
			"params": map[string]string{
				"currency": res["baseCurrency"].(string),
			},
			"id": time.Now().Unix(),
		}
		c.WriteJSON(currencyMessage)

		_, r2message, r2err := c.ReadMessage()
		if r2err != nil {
			log.Println("read:", err)
			return
		}
		var r2 map[string]interface{}
		json.Unmarshal(r2message, &r2)
		res2, ok := r2["result"].(map[string]interface{})
		if !ok {
			fmt.Println("failed to parse params 2")
		}
		if res != nil {
			mutex.Lock()
			idd = res2["id"].(string)
			fullNamee = res2["fullName"].(string)
			mutex.Unlock()
		}
		//**********************************************************************
		// Subscribe to the market data for each symbol
		subscriptionMessage := map[string]interface{}{
			"method": "subscribeTicker",
			"params": map[string]string{
				"symbol": "BTCUSD",
			},
			"id": time.Now().Unix(),
		}
		c.WriteJSON(subscriptionMessage)

		// Store real-time data in memory
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}

			var result map[string]interface{}
			json.Unmarshal(message, &result)

			params, ok := result["params"].(map[string]interface{})
			if !ok {
				fmt.Println("failed to parse params 3")
				continue
			}
			if params != nil {
				mutex.Lock()
				currencies[params["symbol"].(string)] = Currency{
					ID:          idd,
					FullName:    fullNamee,
					Ask:         params["ask"].(string),
					Bid:         params["bid"].(string),
					Last:        params["last"].(string),
					Open:        params["open"].(string),
					Low:         params["low"].(string),
					High:        params["high"].(string),
					FeeCurrency: feeCurrencyy,
				}
				mutex.Unlock()
			}
		}
	}()

	router := gin.Default()
	router.GET("/currency/:symbol", getCurrency)
	router.GET("/currency/all", getAllCurrencies)

	router.Run(":8080")
}

func getCurrency(c *gin.Context) {
	symbol := c.Param("symbol")
	mutex.Lock()
	currency, ok := currencies[symbol]
	mutex.Unlock()
	if !ok {
		c.JSON(404, gin.H{"message": fmt.Sprintf("Currency with symbol %s not found", symbol)})
		return
	}

	c.JSON(200, currency)
}

func getAllCurrencies(c *gin.Context) {
	mutex.Lock()
	allCurrencies := make([]Currency, 0, len(currencies))
	for _, currency := range currencies {
		allCurrencies = append(allCurrencies, currency)
	}
	mutex.Unlock()

	c.JSON(200, gin.H{"currencies": allCurrencies})
}
