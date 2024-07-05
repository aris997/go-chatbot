package recognizer

import (
	"fmt"
	"go-chatbot/matcher"
	"os"

	"gopkg.in/yaml.v3"
)

type Intent struct {
	name    string
	phrases []string
	score   float64
}

func LoadIntentPhrases() []Intent {
	f, err := os.ReadFile("recognizer/intents.yml")
	if err != nil {
		fmt.Println(err)
	}
	yf := make(map[string][]string)
	err = yaml.Unmarshal(f, &yf)
	if err != nil {
		fmt.Println(err)
	}

	intents := []Intent{}
	for name, phrases := range yf {
		var intent Intent
		intent.name = name
		intent.phrases = phrases
		intents = append(intents, intent)
	}
	// fmt.Println(intents)
	// fmt.Println("END OF LOADING FILE")
	return intents
}

func UserIntent(userMessage string) string {
	intents := LoadIntentPhrases()
	var score float64
	intentName := "Fallback"

	for _, intent := range intents {
		for _, phrase := range intent.phrases {
			intent.score += float64(matcher.Distance(userMessage, phrase))
		}
		intent.score = float64(len(intent.phrases)) / intent.score
		// fmt.Println(intent.name, intent.score)
		if intent.score > score {
			score = intent.score
			intentName = intent.name
		}
	}
	if score < 0.1 {
		return "Fallback"
	}
	return intentName
}
