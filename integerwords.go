package euraizodziais

import (
	"slices"
	"strings"
)

// integerToLithuanianWords spells a non-negative integer out in Lithuanian.
func integerToLithuanianWords(number int) string {
	if number == 0 {
		return zeroWord
	}

	chunks := splitIntoThreeDigitChunks(number)

	var words []string
	for position, chunk := range slices.Backward(chunks) {
		if chunk == 0 {
			continue
		}
		words = append(words, renderThreeDigitChunk(chunk)...)
		if position > 0 {
			words = append(words, selectNounForm(chunk, scaleNounForms[position]))
		}
	}

	return strings.Join(words, " ")
}

// splitIntoThreeDigitChunks breaks a number into groups of three digits, least significant
// group first, so a group's index is also its scale: 1 is thousands, 2 is millions.
func splitIntoThreeDigitChunks(number int) []int {
	var chunks []int
	for number > 0 {
		chunks = append(chunks, number%1000)
		number /= 1000
	}
	return chunks
}

func renderThreeDigitChunk(chunk int) []string {
	words := renderHundreds(chunk / 100)
	return append(words, renderTensAndUnits(chunk%100)...)
}

// renderHundreds spells the hundreds place. The unit word is kept even for a single hundred,
// giving "vienas šimtas" rather than "šimtas".
func renderHundreds(hundredsDigit int) []string {
	if hundredsDigit == 0 {
		return nil
	}
	if hundredsDigit == 1 {
		return []string{onesWords[1], hundredSingular}
	}
	return []string{onesWords[hundredsDigit], hundredPlural}
}

func renderTensAndUnits(tensAndUnits int) []string {
	tensDigit := tensAndUnits / 10
	unitsDigit := tensAndUnits % 10

	// The teens are single words, so the units digit is consumed by the tens word.
	if tensDigit == 1 {
		return []string{teenWords[unitsDigit]}
	}

	var words []string
	if tensDigit > 1 {
		words = append(words, tensWords[tensDigit])
	}
	if unitsDigit > 0 {
		words = append(words, onesWords[unitsDigit])
	}
	return words
}
