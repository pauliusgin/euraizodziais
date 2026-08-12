// Package euraizodziais renders a euro amount, given as an integer number of cents,
// as Lithuanian words.
//
// The single entry point is [EuroAmountToWords]. It spells the euro part out in words,
// picks the grammatically correct form of the euro noun, and appends the cents as two
// digits followed by the literal "ct" — cents are never spelled out.
//
// euraizodziais.EuroAmountToWords(12345) // "vienas šimtas dvidešimt trys eurai 45 ct"
//
// # Supported range
//
// The largest supported amount is 2147483647 cents, roughly 21.4 million EUR — the int32
// limit, so the range is the same on a 32-bit and a 64-bit platform. Anything outside returns an
// [ErrNegativeAmount] or an [ErrAmountTooLarge] together with the zero string. Both
// are types, so callers match them with [errors.As], not [errors.Is].
package euraizodziais
