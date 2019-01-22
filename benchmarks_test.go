package snakecase

import (
	"testing"
)

func BenchmarkSnakecaseBorrowedLong(b *testing.B) {
	var s = "invite_your_customers_add_invites"
	b.ReportAllocs()
	b.ResetTimer()
	b.SetBytes(int64(len(s)))
	for n := 0; n < b.N; n++ {
		Snakecase(s)
	}
}

func BenchmarkSnakecaseBorrowedSimple(b *testing.B) {
	var s = "sample_text"
	b.ReportAllocs()
	b.ResetTimer()
	b.SetBytes(int64(len(s)))
	for n := 0; n < b.N; n++ {
		Snakecase(s)
	}
}

func BenchmarkSnakecaseOwnedUnicode(b *testing.B) {
	var s = "ß_ƒ_foo"
	b.ReportAllocs()
	b.ResetTimer()
	b.SetBytes(int64(len(s)))
	for n := 0; n < b.N; n++ {
		Snakecase(s)
	}
}
func BenchmarkSnakecaseOwnedLong(b *testing.B) {
	var s = "inviteYourCustomersAddInvites"
	b.ReportAllocs()
	b.ResetTimer()
	b.SetBytes(int64(len(s)))
	for n := 0; n < b.N; n++ {
		Snakecase(s)
	}
}

func BenchmarkSnakecaseOwnedLongSpecialChars(b *testing.B) {
	var s = "FOO:BAR$BAZ__Sample    Text___"
	b.ReportAllocs()
	b.ResetTimer()
	b.SetBytes(int64(len(s)))
	for n := 0; n < b.N; n++ {
		Snakecase(s)
	}
}

func BenchmarkSnakecaseOwnedSimple(b *testing.B) {
	var s = "sample text"
	b.ReportAllocs()
	b.ResetTimer()
	b.SetBytes(int64(len(s)))
	for n := 0; n < b.N; n++ {
		Snakecase("sample text")
	}
}

func BenchmarkSnakecaseOwnedUnicode2(b *testing.B) {
	var s = "ẞ•¶§ƒ˚foo˙∆˚¬"
	b.ReportAllocs()
	b.ResetTimer()
	b.SetBytes(int64(len(s)))
	for n := 0; n < b.N; n++ {
		Snakecase(s)
	}
}
