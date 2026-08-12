package euraizodziais

const zeroWord = "nulis"

const (
	hundredSingular = "šimtas"
	hundredPlural   = "šimtai"
)

// onesWords holds the unit words, indexed by digit. Index 0 is unused: a zero
// unit contributes no word at all.
var onesWords = [10]string{
	1: "vienas",
	2: "du",
	3: "trys",
	4: "keturi",
	5: "penki",
	6: "šeši",
	7: "septyni",
	8: "aštuoni",
	9: "devyni",
}

// teenWords holds 10 through 19, indexed by the units digit.
var teenWords = [10]string{
	0: "dešimt",
	1: "vienuolika",
	2: "dvylika",
	3: "trylika",
	4: "keturiolika",
	5: "penkiolika",
	6: "šešiolika",
	7: "septyniolika",
	8: "aštuoniolika",
	9: "devyniolika",
}

// tensWords holds the round tens 20 through 90, indexed by the tens digit. Indices 0 and 1
// are unused: a zero tens digit contributes no word, and 10 through 19 come from teenWords.
var tensWords = [10]string{
	2: "dvidešimt",
	3: "trisdešimt",
	4: "keturiasdešimt",
	5: "penkiasdešimt",
	6: "šešiasdešimt",
	7: "septyniasdešimt",
	8: "aštuoniasdešimt",
	9: "devyniasdešimt",
}

// scaleNounForms holds the noun a three-digit chunk takes, indexed by the chunk's position
// counted from the least significant end. Position 0 is the bare units chunk and carries no
// scale word, so its entry is unused. The table stops at milijonas because the supported
// range tops out at twenty one million euros.
var scaleNounForms = [3]lithuanianNounForms{
	1: {Singular: "tūkstantis", Plural: "tūkstančiai", Genitive: "tūkstančių"},
	2: {Singular: "milijonas", Plural: "milijonai", Genitive: "milijonų"},
}

var euroNounForms = lithuanianNounForms{Singular: "euras", Plural: "eurai", Genitive: "eurų"}
