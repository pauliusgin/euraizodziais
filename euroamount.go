package euraizodziais

import (
	"fmt"
	"math"
)

// maxAmountInCents caps the amount at the int32 limit on every platform, so a rendered amount
// does not depend on how wide int happens to be. It is typed int64 because on a 32-bit
// platform the comparison below would otherwise be against the largest possible int and never
// fire. The test suite does not compile on a 32-bit platform: it names out-of-range amounts as
// untyped constants, which overflow int at compile time. Running the tests requires a 64-bit
// platform.
const maxAmountInCents int64 = math.MaxInt32

// EuroAmountToWords renders an amount of euro cents as Lithuanian words, in the form
// "<euros in words> <euro noun> <cents> ct" — for example 12345 gives
// "vienas šimtas dvidešimt trys eurai 45 ct". Cents are always two digits and are never
// spelled out.
//
// Supported amounts run from 0 to 2147483647 cents inclusive. Outside that range the
// returned string is empty and the error is an [ErrNegativeAmount] or an [ErrAmountTooLarge],
// both of which are types and so are matched with [errors.As] rather than [errors.Is].
func EuroAmountToWords(amountInCents int) (string, error) {
	if amountInCents < 0 {
		return "", ErrNegativeAmount{AmountInCents: amountInCents}
	}
	if int64(amountInCents) > maxAmountInCents {
		return "", ErrAmountTooLarge{AmountInCents: amountInCents}
	}

	euros := amountInCents / 100
	cents := amountInCents % 100

	return fmt.Sprintf(
		"%s %s %02d ct",
		integerToLithuanianWords(euros),
		selectNounForm(euros, euroNounForms),
		cents,
	), nil
}
