package euraizodziais

import "testing"

func TestIntegerToLithuanianWordsUnitsAndTeens(t *testing.T) {
	cases := []struct {
		number int
		want   string
	}{
		{number: 0, want: "nulis"},
		{number: 1, want: "vienas"},
		{number: 2, want: "du"},
		{number: 3, want: "trys"},
		{number: 4, want: "keturi"},
		{number: 5, want: "penki"},
		{number: 6, want: "šeši"},
		{number: 7, want: "septyni"},
		{number: 8, want: "aštuoni"},
		{number: 9, want: "devyni"},
		{number: 10, want: "dešimt"},
		{number: 11, want: "vienuolika"},
		{number: 12, want: "dvylika"},
		{number: 13, want: "trylika"},
		{number: 14, want: "keturiolika"},
		{number: 15, want: "penkiolika"},
		{number: 16, want: "šešiolika"},
		{number: 17, want: "septyniolika"},
		{number: 18, want: "aštuoniolika"},
		{number: 19, want: "devyniolika"},
	}

	assertIntegerWords(t, cases)
}

func TestIntegerToLithuanianWordsRoundTensAndCompounds(t *testing.T) {
	cases := []struct {
		number int
		want   string
	}{
		{number: 20, want: "dvidešimt"},
		{number: 21, want: "dvidešimt vienas"},
		{number: 29, want: "dvidešimt devyni"},
		{number: 30, want: "trisdešimt"},
		{number: 33, want: "trisdešimt trys"},
		{number: 40, want: "keturiasdešimt"},
		{number: 44, want: "keturiasdešimt keturi"},
		{number: 50, want: "penkiasdešimt"},
		{number: 55, want: "penkiasdešimt penki"},
		{number: 60, want: "šešiasdešimt"},
		{number: 66, want: "šešiasdešimt šeši"},
		{number: 70, want: "septyniasdešimt"},
		{number: 77, want: "septyniasdešimt septyni"},
		{number: 80, want: "aštuoniasdešimt"},
		{number: 88, want: "aštuoniasdešimt aštuoni"},
		{number: 90, want: "devyniasdešimt"},
		{number: 99, want: "devyniasdešimt devyni"},
	}

	assertIntegerWords(t, cases)
}

func TestIntegerToLithuanianWordsHundreds(t *testing.T) {
	cases := []struct {
		number int
		want   string
	}{
		{number: 100, want: "vienas šimtas"},
		{number: 200, want: "du šimtai"},
		{number: 300, want: "trys šimtai"},
		{number: 400, want: "keturi šimtai"},
		{number: 500, want: "penki šimtai"},
		{number: 600, want: "šeši šimtai"},
		{number: 700, want: "septyni šimtai"},
		{number: 800, want: "aštuoni šimtai"},
		{number: 900, want: "devyni šimtai"},
		{number: 101, want: "vienas šimtas vienas"},
		{number: 110, want: "vienas šimtas dešimt"},
		{number: 115, want: "vienas šimtas penkiolika"},
		{number: 123, want: "vienas šimtas dvidešimt trys"},
		{number: 999, want: "devyni šimtai devyniasdešimt devyni"},
	}

	assertIntegerWords(t, cases)
}

func TestIntegerToLithuanianWordsThousandsAndMillions(t *testing.T) {
	cases := []struct {
		number int
		want   string
	}{
		{number: 1000, want: "vienas tūkstantis"},
		{number: 1001, want: "vienas tūkstantis vienas"},
		{number: 2000, want: "du tūkstančiai"},
		{number: 10000, want: "dešimt tūkstančių"},
		{number: 11000, want: "vienuolika tūkstančių"},
		{number: 20000, want: "dvidešimt tūkstančių"},
		{number: 21000, want: "dvidešimt vienas tūkstantis"},
		{number: 22000, want: "dvidešimt du tūkstančiai"},
		{number: 100000, want: "vienas šimtas tūkstančių"},
		{number: 999999, want: "devyni šimtai devyniasdešimt devyni tūkstančiai devyni šimtai devyniasdešimt devyni"},
		{number: 1000000, want: "vienas milijonas"},
		{number: 1000001, want: "vienas milijonas vienas"},
		{number: 1000500, want: "vienas milijonas penki šimtai"},
		{number: 1001000, want: "vienas milijonas vienas tūkstantis"},
		{number: 2000000, want: "du milijonai"},
		{number: 11000000, want: "vienuolika milijonų"},
		{number: 21000000, want: "dvidešimt vienas milijonas"},
		{number: 100000000, want: "vienas šimtas milijonų"},
	}

	assertIntegerWords(t, cases)
}

func assertIntegerWords(t *testing.T, cases []struct {
	number int
	want   string
}) {
	t.Helper()

	for _, testCase := range cases {
		if got := integerToLithuanianWords(testCase.number); got != testCase.want {
			t.Errorf("integerToLithuanianWords(%d) = %q, want %q", testCase.number, got, testCase.want)
		}
	}
}
