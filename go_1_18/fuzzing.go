package go118

import (
	"errors"
	"unicode/utf8"
)

// go fuzzing: https://go.dev/doc/fuzz/
// fuzz tutorial: https://go.dev/doc/tutorial/fuzz

// With fuzzing, random data is run against your test in an attempt to find vulnerabilities or crash-causing inputs.
// Some examples of vulnerabilities that can be found by fuzzing are SQL injection, buffer overflow, denial of service and cross-site scripting attacks.

// Reverse accept a string, loop over it a byte at a time, and return the reversed string at the end.
// This code is based on the stringutil.Reverse function within golang.org/x/example.
func Reverse(s string) string {
	b := []byte(s)
	for i, j := 0, len(b)-1; i < len(b)/2; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return string(b)
}

// The entire seed corpus used strings in which every character was a single byte.
// However, characters such as 泃 can require several bytes. Thus, reversing the string byte-by-byte will invalidate multi-byte characters.
// The key difference is that Reverse is now iterating over each rune in the string, rather than each byte.
func ReverseFixed(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func ReverseFixedWithErrorReturned(s string) (string, error) {
	if !utf8.ValidString(s) {
		return s, errors.New("input is not valid UTF-8")
	}
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r), nil
}
