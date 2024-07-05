package matcher

import (
	"testing"
)

func TestDistance(t *testing.T) {
	frase := "Tanto va la gatta al lardo che ci lascia lo zampino"
	for idx := 0; idx <= len(frase); idx++ {
		if Distance(frase, frase[:idx]) != len(frase)-idx {
			t.Error("something did not work")
		}
	}
}
