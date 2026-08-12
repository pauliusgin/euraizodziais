package euraizodziais

// lithuanianNounForms holds the three forms a counted noun takes in Lithuanian.
type lithuanianNounForms struct {
	Singular string // vienas euras
	Plural   string // du eurai
	Genitive string // dešimt eurų
}

// selectNounForm picks the form a noun takes after count.
func selectNounForm(count int, forms lithuanianNounForms) string {
	tensDigit := (count / 10) % 10
	unitsDigit := count % 10

	if tensDigit == 1 || unitsDigit == 0 {
		return forms.Genitive
	}
	if unitsDigit == 1 {
		return forms.Singular
	}
	return forms.Plural
}
