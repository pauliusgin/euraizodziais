package euraizodziais

const englishZeroWord = "zero"

const englishHundredWord = "hundred"

const (
	englishEuroSingular = "euro"
	englishEuroPlural   = "euros"
)

// englishOnesWords holds the unit words, indexed by digit. Index 0 is the empty string: a zero
// unit contributes no word at all.
var englishOnesWords = [10]string{
	1: "one",
	2: "two",
	3: "three",
	4: "four",
	5: "five",
	6: "six",
	7: "seven",
	8: "eight",
	9: "nine",
}

// englishTeenWords holds 10 through 19, indexed by the units digit.
var englishTeenWords = [10]string{
	0: "ten",
	1: "eleven",
	2: "twelve",
	3: "thirteen",
	4: "fourteen",
	5: "fifteen",
	6: "sixteen",
	7: "seventeen",
	8: "eighteen",
	9: "nineteen",
}

// englishTensWords holds the round tens 20 through 90, indexed by the tens digit. Indices 0
// and 1 are unused: a zero tens digit contributes no word, and 10 through 19 come from
// englishTeenWords.
var englishTensWords = [10]string{
	2: "twenty",
	3: "thirty",
	4: "forty",
	5: "fifty",
	6: "sixty",
	7: "seventy",
	8: "eighty",
	9: "ninety",
}

// englishScaleWords holds the scale word a three-digit chunk takes, indexed by the chunk's
// position counted from the least significant end. English scale words do not inflect with
// the count, unlike Lithuanian. The table stops at million for the same range reason as
// scaleNounForms.
var englishScaleWords = [3]string{
	1: "thousand",
	2: "million",
}
