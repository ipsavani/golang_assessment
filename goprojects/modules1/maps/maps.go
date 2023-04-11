package maps

import "fmt"

var nameAgeMap map[string]int

func Maps() {
	nameAgeMap = map[string]int{
		"James": 50,
		"Ali":   39,
	}
	fmt.Println("Print the age of James: ", nameAgeMap["James"])

	//We can range through the map and print each value:
	for key, value := range nameAgeMap {
		fmt.Printf("%v is %d years old\n", key, value)
	}

	// adding to map
	nameAgeMap["Pranshu"] = 26
	fmt.Println("new entry:", nameAgeMap)

	// remove from map
	delete(nameAgeMap, "James")
	fmt.Println("entry removed:", nameAgeMap)

	// replacing one entry with another
	nameAgeMap["Ali"] = 30
	fmt.Println("updated entry:", nameAgeMap)

}

func Nested_maps() {
	currency := map[string]map[string]int{
		"Great Britain Pound": {"GBP": 1},
		"Euro":                {"EUR": 2},
		"USA Dollar":          {"USD": 3},
	}

	for key, value := range currency {
		fmt.Printf("Currency Name: %v\n", key)
		for k, v := range value {
			fmt.Printf("\t Currency Code: %v\n\t\t\t Ranking: %v\n\n", k, v)
		}
	}
}
