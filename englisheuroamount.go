package euraizodziais

import "fmt"

// EuroAmountToEnglishWords renders an amount of euro cents as English words, in the form
// "<euros in words> <euro noun> <cents> ct" — for example 12345 gives
// "one hundred twenty-three euros 45 ct". Cents are always two digits and are never
// spelled out.
//
// Supported amounts run from 0 to 2147483647 cents inclusive. Outside that range the
// returned string is empty and the error is an [ErrNegativeAmount] or an [ErrAmountTooLarge],
// both of which are types and so are matched with [errors.As] rather than [errors.Is].
func EuroAmountToEnglishWords(amountInCents int) (string, error) {
	if err := validateAmountInCents(amountInCents); err != nil {
		return "", err
	}

	euros := amountInCents / 100
	cents := amountInCents % 100

	return fmt.Sprintf(
		"%s %s %02d ct",
		integerToEnglishWords(euros),
		selectEnglishEuroNoun(euros),
		cents,
	), nil
}

// selectEnglishEuroNoun picks the euro noun after count. English counts zero as plural,
// giving "zero euros".
func selectEnglishEuroNoun(count int) string {
	if count == 1 {
		return englishEuroSingular
	}
	return englishEuroPlural
}
