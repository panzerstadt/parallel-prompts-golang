package main

import (
	"fmt"
)

const (
	seed        int     = 999
	temperature float32 = 0.1
	categories  string  = "cafe, travel, utilities, daily, subscriptions, phone, misc, hobbies"
)

func main() {

	rows := []string{
		"Akasaka Biz Tower",
		"Ueshima Coffee Akasaka Hitotsugi",
		"AIA IPOS ONLINE (Country: MY)",
		"Rakuten Mobile Upgrade",
		"AEON SMKT-METRO PRIM (Country: MY)",
		"WREN CLIMATE PBC (Country: US)",
		"Rakuten Mobile Communication Fee",
		"Rakuten Cafe 9F",
		"TRAWICK TRAVEL INSUR (Country: US)",
		"135065 WING TAK",
		"093830 PATREON",
		"231067 AMAZON C (AIRTAGS)",
		"660174 GOOGLE C",
		"469885 AMAZON C (TIME BLOCK PLANNER)",
		"464149 APPLE.CO",
		"611165 KAI ANJI",
		"698172 JR EAST",
		"640420 PAYPAL (PRINTFUL)",
		"807276 STEAMGAM",
		"037358 UESHIMAK",
		"096420 RAMEN KI",
		"330439 MICROSOFT",
		"502471 NEBULA S",
		"934042 STEAM PU",
		"502357 AIRASIA",
		"Excel Shiori",
		"DHL JAPAN INC.",
	}

	ch := make(chan string)
	for i, description := range rows {
		go guessCategory(description, i, ch)
	}
	result := make([]string, len(rows))

	for i := range result {
		result[i] = <-ch
		fmt.Println(result[i])
	}
}

func guessCategory(description string, idx int, ch chan string) {
	answer, err := label(description)
	if err != nil {
		fmt.Println(err)
		ch <- "errored."
	}

	ch <- fmt.Sprint(idx) + "-> " + description + ": " + answer
}
