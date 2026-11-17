package euraizodziais

import "testing"

type englishIntegerWordsCase struct {
	name   string
	number int
	want   string
}

func TestIntegerToEnglishWordsUnitsAndTeens(t *testing.T) {
	cases := []englishIntegerWordsCase{
		{name: "zero", number: 0, want: "zero"},
		{name: "one", number: 1, want: "one"},
		{name: "five", number: 5, want: "five"},
		{name: "nine", number: 9, want: "nine"},
		{name: "ten", number: 10, want: "ten"},
		{name: "eleven", number: 11, want: "eleven"},
		{name: "twelve", number: 12, want: "twelve"},
		{name: "thirteen", number: 13, want: "thirteen"},
		{name: "fifteen", number: 15, want: "fifteen"},
		{name: "eighteen", number: 18, want: "eighteen"},
		{name: "nineteen", number: 19, want: "nineteen"},
	}

	assertEnglishIntegerWords(t, cases)
}

func TestIntegerToEnglishWordsRoundTensAndCompounds(t *testing.T) {
	cases := []englishIntegerWordsCase{
		{name: "twenty", number: 20, want: "twenty"},
		{name: "twenty one", number: 21, want: "twenty-one"},
		{name: "twenty nine", number: 29, want: "twenty-nine"},
		{name: "thirty", number: 30, want: "thirty"},
		{name: "forty four", number: 44, want: "forty-four"},
		{name: "fifty five", number: 55, want: "fifty-five"},
		{name: "sixty", number: 60, want: "sixty"},
		{name: "seventy seven", number: 77, want: "seventy-seven"},
		{name: "eighty", number: 80, want: "eighty"},
		{name: "ninety nine", number: 99, want: "ninety-nine"},
	}

	assertEnglishIntegerWords(t, cases)
}

func TestIntegerToEnglishWordsHundreds(t *testing.T) {
	cases := []englishIntegerWordsCase{
		{name: "one hundred", number: 100, want: "one hundred"},
		{name: "hundred one", number: 101, want: "one hundred one"},
		{name: "hundred ten", number: 110, want: "one hundred ten"},
		{name: "hundred twenty three", number: 123, want: "one hundred twenty-three"},
		{name: "five hundred", number: 500, want: "five hundred"},
		{name: "nine hundred ninety nine", number: 999, want: "nine hundred ninety-nine"},
	}

	assertEnglishIntegerWords(t, cases)
}

func TestIntegerToEnglishWordsThousandsAndMillions(t *testing.T) {
	cases := []englishIntegerWordsCase{
		{name: "one thousand", number: 1000, want: "one thousand"},
		{name: "thousand one", number: 1001, want: "one thousand one"},
		{name: "twenty one thousand", number: 21000, want: "twenty-one thousand"},
		{name: "hundred thousand", number: 100000, want: "one hundred thousand"},
		{name: "largest six digits", number: 999999, want: "nine hundred ninety-nine thousand nine hundred ninety-nine"},
		{name: "one million", number: 1000000, want: "one million"},
		{name: "million one skips empty thousands", number: 1000001, want: "one million one"},
		{name: "million thousand", number: 1001000, want: "one million one thousand"},
		{
			name:   "maximum euros",
			number: 21474836,
			want:   "twenty-one million four hundred seventy-four thousand eight hundred thirty-six",
		},
	}

	assertEnglishIntegerWords(t, cases)
}

func assertEnglishIntegerWords(t *testing.T, cases []englishIntegerWordsCase) {
	t.Helper()
	for _, testCase := range cases {
		t.Run(testCase.name, func(t *testing.T) {
			if got := integerToEnglishWords(testCase.number); got != testCase.want {
				t.Errorf("integerToEnglishWords(%d) = %q, want %q", testCase.number, got, testCase.want)
			}
		})
	}
}
