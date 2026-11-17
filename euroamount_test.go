package euraizodziais

import (
	"errors"
	"testing"
)

func TestEuroAmountToLithuanianWordsAnchorCases(t *testing.T) {
	cases := []struct {
		name          string
		amountInCents int
		want          string
	}{
		{name: "zero", amountInCents: 0, want: "nulis eurų 00 ct"},
		{name: "one cent", amountInCents: 1, want: "nulis eurų 01 ct"},
		{name: "one euro", amountInCents: 100, want: "vienas euras 00 ct"},
		{name: "one euro one cent", amountInCents: 101, want: "vienas euras 01 ct"},
		{name: "two euros", amountInCents: 200, want: "du eurai 00 ct"},
		{name: "eleven euros", amountInCents: 1100, want: "vienuolika eurų 00 ct"},
		{name: "twenty one euros", amountInCents: 2100, want: "dvidešimt vienas euras 00 ct"},
		{name: "hundred twenty three euros forty five cents", amountInCents: 12345, want: "vienas šimtas dvidešimt trys eurai 45 ct"},
		{name: "one thousand euros", amountInCents: 100000, want: "vienas tūkstantis eurų 00 ct"},
		{name: "twenty thousand euros", amountInCents: 2000000, want: "dvidešimt tūkstančių eurų 00 ct"},
		{name: "one million euros", amountInCents: 100000000, want: "vienas milijonas eurų 00 ct"},
		{name: "twenty one million euros", amountInCents: 2100000000, want: "dvidešimt vienas milijonas eurų 00 ct"},
	}

	for _, testCase := range cases {
		t.Run(testCase.name, func(t *testing.T) {
			got, err := EuroAmountToLithuanianWords(testCase.amountInCents)
			if err != nil {
				t.Fatalf("EuroAmountToLithuanianWords(%d) returned error %v", testCase.amountInCents, err)
			}
			if got != testCase.want {
				t.Errorf("EuroAmountToLithuanianWords(%d) = %q, want %q", testCase.amountInCents, got, testCase.want)
			}
		})
	}
}

func TestEuroAmountToLithuanianWordsCentsRendering(t *testing.T) {
	cases := []struct {
		name          string
		amountInCents int
		want          string
	}{
		{name: "cents 00", amountInCents: 500, want: "penki eurai 00 ct"},
		{name: "cents 01", amountInCents: 501, want: "penki eurai 01 ct"},
		{name: "cents 09", amountInCents: 509, want: "penki eurai 09 ct"},
		{name: "cents 10", amountInCents: 510, want: "penki eurai 10 ct"},
		{name: "cents 45", amountInCents: 545, want: "penki eurai 45 ct"},
		{name: "cents 99", amountInCents: 599, want: "penki eurai 99 ct"},
		{name: "cents without euros", amountInCents: 99, want: "nulis eurų 99 ct"},
	}

	for _, testCase := range cases {
		t.Run(testCase.name, func(t *testing.T) {
			got, err := EuroAmountToLithuanianWords(testCase.amountInCents)
			if err != nil {
				t.Fatalf("EuroAmountToLithuanianWords(%d) returned error %v", testCase.amountInCents, err)
			}
			if got != testCase.want {
				t.Errorf("EuroAmountToLithuanianWords(%d) = %q, want %q", testCase.amountInCents, got, testCase.want)
			}
		})
	}
}

func TestEuroAmountToLithuanianWordsEuroNounForm(t *testing.T) {
	cases := []struct {
		name          string
		amountInCents int
		want          string
	}{
		{name: "zero euros takes genitive", amountInCents: 0, want: "nulis eurų 00 ct"},
		{name: "one euro takes singular", amountInCents: 100, want: "vienas euras 00 ct"},
		{name: "two euros take plural", amountInCents: 200, want: "du eurai 00 ct"},
		{name: "nine euros take plural", amountInCents: 900, want: "devyni eurai 00 ct"},
		{name: "ten euros take genitive", amountInCents: 1000, want: "dešimt eurų 00 ct"},
		{name: "eleven euros take genitive", amountInCents: 1100, want: "vienuolika eurų 00 ct"},
		{name: "twelve euros take genitive", amountInCents: 1200, want: "dvylika eurų 00 ct"},
		{name: "thirteen euros take genitive", amountInCents: 1300, want: "trylika eurų 00 ct"},
		{name: "fourteen euros take genitive", amountInCents: 1400, want: "keturiolika eurų 00 ct"},
		{name: "fifteen euros take genitive", amountInCents: 1500, want: "penkiolika eurų 00 ct"},
		{name: "sixteen euros take genitive", amountInCents: 1600, want: "šešiolika eurų 00 ct"},
		{name: "seventeen euros take genitive", amountInCents: 1700, want: "septyniolika eurų 00 ct"},
		{name: "eighteen euros take genitive", amountInCents: 1800, want: "aštuoniolika eurų 00 ct"},
		{name: "nineteen euros take genitive", amountInCents: 1900, want: "devyniolika eurų 00 ct"},
		{name: "twenty euros take genitive", amountInCents: 2000, want: "dvidešimt eurų 00 ct"},
		{name: "twenty one euros take singular", amountInCents: 2100, want: "dvidešimt vienas euras 00 ct"},
		{name: "twenty two euros take plural", amountInCents: 2200, want: "dvidešimt du eurai 00 ct"},
		{name: "thirty euros take genitive", amountInCents: 3000, want: "trisdešimt eurų 00 ct"},
		{name: "hundred euros take genitive", amountInCents: 10000, want: "vienas šimtas eurų 00 ct"},
		{name: "hundred one euros take singular", amountInCents: 10100, want: "vienas šimtas vienas euras 00 ct"},
		{name: "hundred eleven euros take genitive", amountInCents: 11100, want: "vienas šimtas vienuolika eurų 00 ct"},
		{name: "thousand euros take genitive", amountInCents: 100000, want: "vienas tūkstantis eurų 00 ct"},
		{name: "thousand one euros take singular", amountInCents: 100100, want: "vienas tūkstantis vienas euras 00 ct"},
	}

	for _, testCase := range cases {
		t.Run(testCase.name, func(t *testing.T) {
			got, err := EuroAmountToLithuanianWords(testCase.amountInCents)
			if err != nil {
				t.Fatalf("EuroAmountToLithuanianWords(%d) returned error %v", testCase.amountInCents, err)
			}
			if got != testCase.want {
				t.Errorf("EuroAmountToLithuanianWords(%d) = %q, want %q", testCase.amountInCents, got, testCase.want)
			}
		})
	}
}

func TestEuroAmountToLithuanianWordsChunkBoundaries(t *testing.T) {
	cases := []struct {
		name          string
		amountInCents int
		want          string
	}{
		{name: "999 euros", amountInCents: 99900, want: "devyni šimtai devyniasdešimt devyni eurai 00 ct"},
		{name: "1000 euros", amountInCents: 100000, want: "vienas tūkstantis eurų 00 ct"},
		{name: "1001 euros", amountInCents: 100100, want: "vienas tūkstantis vienas euras 00 ct"},
		{name: "999999 euros", amountInCents: 99999900, want: "devyni šimtai devyniasdešimt devyni tūkstančiai devyni šimtai devyniasdešimt devyni eurai 00 ct"},
		{name: "1000000 euros", amountInCents: 100000000, want: "vienas milijonas eurų 00 ct"},
		{name: "1000001 euros", amountInCents: 100000100, want: "vienas milijonas vienas euras 00 ct"},
		{name: "maximum amount", amountInCents: 2147483647, want: "dvidešimt vienas milijonas keturi šimtai septyniasdešimt keturi tūkstančiai aštuoni šimtai trisdešimt šeši eurai 47 ct"},
	}

	for _, testCase := range cases {
		t.Run(testCase.name, func(t *testing.T) {
			got, err := EuroAmountToLithuanianWords(testCase.amountInCents)
			if err != nil {
				t.Fatalf("EuroAmountToLithuanianWords(%d) returned error %v", testCase.amountInCents, err)
			}
			if got != testCase.want {
				t.Errorf("EuroAmountToLithuanianWords(%d) = %q, want %q", testCase.amountInCents, got, testCase.want)
			}
		})
	}
}

func TestEuroAmountToLithuanianWordsRejectsNegativeAmount(t *testing.T) {
	for _, amountInCents := range []int{-1, -100, -2147483648} {
		got, err := EuroAmountToLithuanianWords(amountInCents)
		if got != "" {
			t.Errorf("EuroAmountToLithuanianWords(%d) = %q, want empty string", amountInCents, got)
		}

		var negativeAmount ErrNegativeAmount
		if !errors.As(err, &negativeAmount) {
			t.Fatalf("EuroAmountToLithuanianWords(%d) error = %v, want ErrNegativeAmount", amountInCents, err)
		}
		if negativeAmount.AmountInCents != amountInCents {
			t.Errorf("ErrNegativeAmount.AmountInCents = %d, want %d", negativeAmount.AmountInCents, amountInCents)
		}
	}
}

func TestEuroAmountToLithuanianWordsRejectsTooLargeAmount(t *testing.T) {
	for _, amountInCents := range []int{2147483648, 9999999999} {
		got, err := EuroAmountToLithuanianWords(amountInCents)
		if got != "" {
			t.Errorf("EuroAmountToLithuanianWords(%d) = %q, want empty string", amountInCents, got)
		}

		var tooLarge ErrAmountTooLarge
		if !errors.As(err, &tooLarge) {
			t.Fatalf("EuroAmountToLithuanianWords(%d) error = %v, want ErrAmountTooLarge", amountInCents, err)
		}
		if tooLarge.AmountInCents != amountInCents {
			t.Errorf("ErrAmountTooLarge.AmountInCents = %d, want %d", tooLarge.AmountInCents, amountInCents)
		}
	}
}

func TestErrorMessagesCarryTheOffendingAmount(t *testing.T) {
	cases := []struct {
		name string
		err  error
		want string
	}{
		{
			name: "negative amount",
			err:  ErrNegativeAmount{AmountInCents: -5},
			want: "negative amount -5 cents: expected 0 to 2147483647 cents, that is 0.00 to 21,474,836.47 EUR",
		},
		{
			name: "amount too large",
			err:  ErrAmountTooLarge{AmountInCents: 2147483648},
			want: "amount 2147483648 cents is too large: expected 0 to 2147483647 cents, that is 0.00 to 21,474,836.47 EUR",
		},
	}

	for _, testCase := range cases {
		t.Run(testCase.name, func(t *testing.T) {
			if got := testCase.err.Error(); got != testCase.want {
				t.Errorf("Error() = %q, want %q", got, testCase.want)
			}
		})
	}
}
