package snakecase

import (
	"testing"
)

func BenchmarkUnchangedLong(b *testing.B) {
	var s = "invite_your_customers_add_invites"
	b.SetBytes(int64(len(s)))
	for n := 0; n < b.N; n++ {
		Snakecase(s)
	}
}

func BenchmarkUnchangedSimple(b *testing.B) {
	var s = "sample_text"
	b.SetBytes(int64(len(s)))
	for n := 0; n < b.N; n++ {
		Snakecase(s)
	}
}

func BenchmarkModifiedUnicode(b *testing.B) {
	var s = "ß_ƒ_foo"
	b.SetBytes(int64(len(s)))
	for n := 0; n < b.N; n++ {
		Snakecase(s)
	}
}
func BenchmarkModifiedLong(b *testing.B) {
	var s = "inviteYourCustomersAddInvites"
	b.SetBytes(int64(len(s)))
	for n := 0; n < b.N; n++ {
		Snakecase(s)
	}
}

func BenchmarkModifiedLongSpecialChars(b *testing.B) {
	var s = "FOO:BAR$BAZ__Sample    Text___"
	b.SetBytes(int64(len(s)))
	for n := 0; n < b.N; n++ {
		Snakecase(s)
	}
}

func BenchmarkModifiedSimple(b *testing.B) {
	var s = "sample text"
	b.SetBytes(int64(len(s)))
	for n := 0; n < b.N; n++ {
		Snakecase("sample text")
	}
}

func BenchmarkModifiedUnicode2(b *testing.B) {
	var s = "ẞ•¶§ƒ˚foo˙∆˚¬"
	b.SetBytes(int64(len(s)))
	for n := 0; n < b.N; n++ {
		Snakecase(s)
	}
}

func BenchmarkLeadingUnderscoresDigitUpper(b *testing.B) {
	var s = "_5TEst"
	b.SetBytes(int64(len(s)))
	for n := 0; n < b.N; n++ {
		Snakecase(s)
	}
}

func BenchmarkDigitUpper(b *testing.B) {
	var s = "5TEst"
	b.SetBytes(int64(len(s)))
	for n := 0; n < b.N; n++ {
		Snakecase(s)
	}
}

func BenchmarkDigitUpper2(b *testing.B) {
	var s = "lk0B@bFmjrLQ_Z6YL"
	b.SetBytes(int64(len(s)))
	for n := 0; n < b.N; n++ {
		Snakecase(s)
	}
}
