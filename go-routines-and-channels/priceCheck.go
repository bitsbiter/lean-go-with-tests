package main

import (
	"fmt"
	"math/rand"
	"time"
)

const MAX_CHICKEN_PRICE float32 = 3

func main() {
	var chickenChannel = make(chan map[string]float32)
	var websites = []string{"kip.nl", "kip.be", "kip.com"}

	for i := range websites {
		go checkChickenPrices(websites[i], chickenChannel)
	}

	sendMessage(chickenChannel)
}

func checkChickenPrices(website string, chickenChannel chan map[string]float32) {
	for {
		time.Sleep(time.Second * 1)
		var chickenPrice = rand.Float32() * 20

		if chickenPrice <= MAX_CHICKEN_PRICE {
			chickenChannel <- map[string]float32{website: chickenPrice}
		}
	}
}

func sendMessage(chickenChannel chan map[string]float32) {
	for {
		result := <-chickenChannel

		for website, price := range result {
			fmt.Printf("\nFound a deal on chicken at: %s, for %.2f euro", website, price)
		}
	}
}
