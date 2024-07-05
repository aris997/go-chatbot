package recognizer

import (
	"fmt"
	"fulfillment/matcher"
	"os"

	"gopkg.in/yaml.v3"
)

type Intent struct {
	name    string
	phrases []string
	score   int
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
	fmt.Println(intents)
	fmt.Println("END OF LOADING FILE")
	return intents
}

func UserIntent(userMessage string) string {
	intents := LoadIntentPhrases()

	for _, intent := range intents {
		for _, phrase := range intent.phrases {
			intent.score += matcher.Distance(userMessage, phrase)
		}
		fmt.Println(intent.name, intent.score)
	}
	return "Fallback"
}
