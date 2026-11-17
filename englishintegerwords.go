package euraizodziais

import (
	"slices"
	"strings"
)

// integerToEnglishWords spells a non-negative integer out in English, without "and" and with
// hyphenated compound tens: 123 gives "one hundred twenty-three".
func integerToEnglishWords(number int) string {
	if number == 0 {
		return englishZeroWord
	}

	chunks := splitIntoThreeDigitChunks(number)

	var words []string
	for position, chunk := range slices.Backward(chunks) {
		if chunk == 0 {
			continue
		}
		words = append(words, renderEnglishThreeDigitChunk(chunk)...)
		if position > 0 {
			words = append(words, englishScaleWords[position])
		}
	}

	return strings.Join(words, " ")
}

func renderEnglishThreeDigitChunk(chunk int) []string {
	words := renderEnglishHundreds(chunk / 100)
	tensAndUnits := renderEnglishTensAndUnits(chunk % 100)
	if tensAndUnits == "" {
		return words
	}
	return append(words, tensAndUnits)
}

func renderEnglishHundreds(hundredsDigit int) []string {
	if hundredsDigit == 0 {
		return nil
	}
	return []string{englishOnesWords[hundredsDigit], englishHundredWord}
}

// renderEnglishTensAndUnits returns a single word, since compound tens are hyphenated
// ("twenty-one"); an empty string means the chunk ends in 00.
func renderEnglishTensAndUnits(tensAndUnits int) string {
	tensDigit := tensAndUnits / 10
	unitsDigit := tensAndUnits % 10

	if tensDigit == 0 {
		return englishOnesWords[unitsDigit]
	}
	if tensDigit == 1 {
		return englishTeenWords[unitsDigit]
	}
	if unitsDigit == 0 {
		return englishTensWords[tensDigit]
	}
	return englishTensWords[tensDigit] + "-" + englishOnesWords[unitsDigit]
}
