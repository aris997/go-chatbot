package actions

func Dispatcher(intentName string) string {
	switch intentName {
	case "IntentGreet":
		return IntentGreet()

	case "IntentThanks":
		return IntentThanks()

	case "Fallback":
		return Fallback()

	}
	return Fallback()
}
