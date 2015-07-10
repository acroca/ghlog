package main

import (
	"fmt"
	"os"
)

func main() {
	gh := NewGhWrapper(os.Getenv("GH_TOKEN"))
	events := gh.GetEvents()

	for _, event := range events {
		if event != nil {
			fmt.Println("-----------------------------------------")
			fmt.Println(event.GetEventBody())
		}
	}
}
