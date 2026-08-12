package euraizodziais

import "fmt"

const amountExpectationText = "expected 0 to 2147483647 cents, that is 0.00 to 21,474,836.47 EUR"

// ErrNegativeAmount reports an amount below zero cents.
type ErrNegativeAmount struct {
	AmountInCents int
}

func (e ErrNegativeAmount) Error() string {
	return fmt.Sprintf("negative amount %d cents: %s", e.AmountInCents, amountExpectationText)
}

// ErrAmountTooLarge reports an amount above the largest supported one.
type ErrAmountTooLarge struct {
	AmountInCents int
}

func (e ErrAmountTooLarge) Error() string {
	return fmt.Sprintf("amount %d cents is too large: %s", e.AmountInCents, amountExpectationText)
}
