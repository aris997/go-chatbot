package main

import (
	"fmt"
	"fulfillment/actions"
	"fulfillment/recognizer"

	"os"
)

func main() {
	if len(os.Args) == 0 {
		fmt.Println("Bye.")
	}

	userMessage := os.Args[1]
	intentName := recognizer.UserIntent(userMessage)
	fmt.Println(actions.Dispatcher(intentName))
}
