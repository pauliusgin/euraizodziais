package euraizodziais

import (
	"errors"
	"testing"
)

type englishEuroAmountCase struct {
	name          string
	amountInCents int
	want          string
}

func TestEuroAmountToEnglishWordsAnchorCases(t *testing.T) {
	cases := []englishEuroAmountCase{
		{name: "zero", amountInCents: 0, want: "zero euros 00 ct"},
		{name: "one cent", amountInCents: 1, want: "zero euros 01 ct"},
		{name: "one euro", amountInCents: 100, want: "one euro 00 ct"},
		{name: "one euro one cent", amountInCents: 101, want: "one euro 01 ct"},
		{name: "two euros", amountInCents: 200, want: "two euros 00 ct"},
		{name: "eleven euros", amountInCents: 1100, want: "eleven euros 00 ct"},
		{name: "twenty one euros", amountInCents: 2100, want: "twenty-one euros 00 ct"},
		{name: "hundred twenty three euros forty five cents", amountInCents: 12345, want: "one hundred twenty-three euros 45 ct"},
		{name: "one thousand euros", amountInCents: 100000, want: "one thousand euros 00 ct"},
		{name: "one million euros", amountInCents: 100000000, want: "one million euros 00 ct"},
		{name: "million one euros", amountInCents: 100000100, want: "one million one euros 00 ct"},
		{
			name:          "maximum amount",
			amountInCents: 2147483647,
			want:          "twenty-one million four hundred seventy-four thousand eight hundred thirty-six euros 47 ct",
		},
	}

	assertEnglishEuroAmounts(t, cases)
}

func TestEuroAmountToEnglishWordsCentsRendering(t *testing.T) {
	cases := []englishEuroAmountCase{
		{name: "cents 00", amountInCents: 500, want: "five euros 00 ct"},
		{name: "cents 01", amountInCents: 501, want: "five euros 01 ct"},
		{name: "cents 09", amountInCents: 509, want: "five euros 09 ct"},
		{name: "cents 10", amountInCents: 510, want: "five euros 10 ct"},
		{name: "cents 99", amountInCents: 599, want: "five euros 99 ct"},
		{name: "cents without euros", amountInCents: 99, want: "zero euros 99 ct"},
	}

	assertEnglishEuroAmounts(t, cases)
}

func TestSelectEnglishEuroNoun(t *testing.T) {
	cases := []struct {
		name  string
		count int
		want  string
	}{
		{name: "zero takes plural", count: 0, want: "euros"},
		{name: "one takes singular", count: 1, want: "euro"},
		{name: "two takes plural", count: 2, want: "euros"},
		{name: "eleven takes plural", count: 11, want: "euros"},
		{name: "twenty one takes plural", count: 21, want: "euros"},
		{name: "hundred one takes plural", count: 101, want: "euros"},
	}

	for _, testCase := range cases {
		t.Run(testCase.name, func(t *testing.T) {
			if got := selectEnglishEuroNoun(testCase.count); got != testCase.want {
				t.Errorf("selectEnglishEuroNoun(%d) = %q, want %q", testCase.count, got, testCase.want)
			}
		})
	}
}

func TestEuroAmountToEnglishWordsRejectsNegativeAmount(t *testing.T) {
	for _, amountInCents := range []int{-1, -2147483648} {
		got, err := EuroAmountToEnglishWords(amountInCents)
		if got != "" {
			t.Errorf("EuroAmountToEnglishWords(%d) = %q, want empty string", amountInCents, got)
		}

		var negativeAmount ErrNegativeAmount
		if !errors.As(err, &negativeAmount) {
			t.Fatalf("EuroAmountToEnglishWords(%d) error = %v, want ErrNegativeAmount", amountInCents, err)
		}
		if negativeAmount.AmountInCents != amountInCents {
			t.Errorf("ErrNegativeAmount.AmountInCents = %d, want %d", negativeAmount.AmountInCents, amountInCents)
		}
	}
}

func TestEuroAmountToEnglishWordsRejectsTooLargeAmount(t *testing.T) {
	for _, amountInCents := range []int{2147483648, 9999999999} {
		got, err := EuroAmountToEnglishWords(amountInCents)
		if got != "" {
			t.Errorf("EuroAmountToEnglishWords(%d) = %q, want empty string", amountInCents, got)
		}

		var tooLarge ErrAmountTooLarge
		if !errors.As(err, &tooLarge) {
			t.Fatalf("EuroAmountToEnglishWords(%d) error = %v, want ErrAmountTooLarge", amountInCents, err)
		}
		if tooLarge.AmountInCents != amountInCents {
			t.Errorf("ErrAmountTooLarge.AmountInCents = %d, want %d", tooLarge.AmountInCents, amountInCents)
		}
	}
}

func assertEnglishEuroAmounts(t *testing.T, cases []englishEuroAmountCase) {
	t.Helper()
	for _, testCase := range cases {
		t.Run(testCase.name, func(t *testing.T) {
			got, err := EuroAmountToEnglishWords(testCase.amountInCents)
			if err != nil {
				t.Fatalf("EuroAmountToEnglishWords(%d) returned error %v", testCase.amountInCents, err)
			}
			if got != testCase.want {
				t.Errorf("EuroAmountToEnglishWords(%d) = %q, want %q", testCase.amountInCents, got, testCase.want)
			}
		})
	}
}
