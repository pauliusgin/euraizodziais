package euraizodziais

import "testing"

func TestSelectNounFormPicksEuroForm(t *testing.T) {
	cases := []struct {
		name  string
		count int
		want  string
	}{
		{name: "zero takes genitive", count: 0, want: "eurų"},
		{name: "one takes singular", count: 1, want: "euras"},
		{name: "two takes plural", count: 2, want: "eurai"},
		{name: "nine takes plural", count: 9, want: "eurai"},
		{name: "ten takes genitive", count: 10, want: "eurų"},
		{name: "eleven takes genitive", count: 11, want: "eurų"},
		{name: "twelve takes genitive", count: 12, want: "eurų"},
		{name: "nineteen takes genitive", count: 19, want: "eurų"},
		{name: "twenty takes genitive", count: 20, want: "eurų"},
		{name: "twenty one takes singular", count: 21, want: "euras"},
		{name: "twenty two takes plural", count: 22, want: "eurai"},
		{name: "ninety takes genitive", count: 90, want: "eurų"},
		{name: "ninety one takes singular", count: 91, want: "euras"},
		{name: "hundred takes genitive", count: 100, want: "eurų"},
		{name: "hundred one takes singular", count: 101, want: "euras"},
		{name: "hundred eleven takes genitive", count: 111, want: "eurų"},
		{name: "hundred twelve takes genitive", count: 112, want: "eurų"},
		{name: "hundred twenty two takes plural", count: 122, want: "eurai"},
		{name: "thousand takes genitive", count: 1000, want: "eurų"},
		{name: "thousand one takes singular", count: 1001, want: "euras"},
		{name: "thousand eleven takes genitive", count: 1011, want: "eurų"},
		{name: "million takes genitive", count: 1000000, want: "eurų"},
		{name: "twenty one million takes genitive", count: 21000000, want: "eurų"},
		{name: "maximum euros take plural", count: 21474836, want: "eurai"},
	}

	for _, testCase := range cases {
		t.Run(testCase.name, func(t *testing.T) {
			if got := selectNounForm(testCase.count, euroNounForms); got != testCase.want {
				t.Errorf("selectNounForm(%d, euroNounForms) = %q, want %q", testCase.count, got, testCase.want)
			}
		})
	}
}

func TestSelectNounFormPicksScaleForm(t *testing.T) {
	cases := []struct {
		name  string
		count int
		forms lithuanianNounForms
		want  string
	}{
		{name: "one thousand", count: 1, forms: scaleNounForms[1], want: "tūkstantis"},
		{name: "two thousands", count: 2, forms: scaleNounForms[1], want: "tūkstančiai"},
		{name: "ten thousands", count: 10, forms: scaleNounForms[1], want: "tūkstančių"},
		{name: "eleven thousands", count: 11, forms: scaleNounForms[1], want: "tūkstančių"},
		{name: "twenty thousands", count: 20, forms: scaleNounForms[1], want: "tūkstančių"},
		{name: "twenty one thousands", count: 21, forms: scaleNounForms[1], want: "tūkstantis"},
		{name: "hundred thousands", count: 100, forms: scaleNounForms[1], want: "tūkstančių"},
		{name: "one million", count: 1, forms: scaleNounForms[2], want: "milijonas"},
		{name: "two millions", count: 2, forms: scaleNounForms[2], want: "milijonai"},
		{name: "eleven millions", count: 11, forms: scaleNounForms[2], want: "milijonų"},
		{name: "hundred millions", count: 100, forms: scaleNounForms[2], want: "milijonų"},
	}

	for _, testCase := range cases {
		t.Run(testCase.name, func(t *testing.T) {
			if got := selectNounForm(testCase.count, testCase.forms); got != testCase.want {
				t.Errorf("selectNounForm(%d, %+v) = %q, want %q", testCase.count, testCase.forms, got, testCase.want)
			}
		})
	}
}
