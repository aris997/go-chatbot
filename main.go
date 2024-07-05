package main

import (
	"fmt"
	"go-chatbot/actions"
	"go-chatbot/recognizer"
)

func main() {
	var userMessage string
	for {
		userMessage = ""
		fmt.Print("YOU: ")
		fmt.Scanf("%v", &userMessage)
		if userMessage == "exit" {
			return
		}
		intentName := recognizer.UserIntent(userMessage)
		fmt.Println("GoBOT: ", actions.Dispatcher(intentName))

	}
}
