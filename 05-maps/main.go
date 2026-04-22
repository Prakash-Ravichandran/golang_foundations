package main

import "fmt"

func main() {
    bikes := map[string]string {
		"yamaha": "mt-07",
		"ducati": "panigale",
		"bmw":    "s1000rr",
	}

	fmt.Println(bikes)

	cars := make(map[string]string)

	cars["audi"] = "a4"
	cars["mercedes"] = "c63"
	cars["porsche"] = "911"

	fmt.Println(cars)

	delete(cars, "mercedes")

	fmt.Println("modified cars:", cars)

	var aircrafts map[string]string

	aircrafts = make(map[string]string)
	aircrafts["boeing"] = "747"
	aircrafts["airbus"] = "a380"
	aircrafts["embraer"] = "e195"

	fmt.Println(aircrafts)
}