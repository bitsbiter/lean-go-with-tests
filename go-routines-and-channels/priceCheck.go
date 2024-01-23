package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

// everything above this price is not a deal
const MAX_CHICKEN_PRICE float32 = 10

var lowestChickenPrice float32

func main() {
	var chickenChannel = make(chan map[string]float32)
	var websites = []string{"kip.nl", "kip.be", "kip.com"}

	lowestChickenPrice = MAX_CHICKEN_PRICE

	for i := range websites {
		go checkChickenPrices(websites[i], chickenChannel)
	}

	sendMessage(chickenChannel)
}

func checkChickenPrices(website string, chickenChannel chan map[string]float32) {
	for {
		time.Sleep(time.Second * 1)
		var chickenPrice = rand.Float32() * 20

		if chickenPrice <= lowestChickenPrice {
			lowestChickenPrice = chickenPrice
			chickenChannel <- map[string]float32{website: chickenPrice}
		}
	}
}

func sendMessage(chickenChannel chan map[string]float32) {
	for {
		result := <-chickenChannel

		for website, price := range result {
			if price < 0.00 {
				fmt.Printf("\nFound free chicken! at: %s", website)
				os.Exit(0)
			}

			fmt.Printf("\nFound a deal on chicken at: %s, for %.2f euro", website, price)
		}
	}
}
