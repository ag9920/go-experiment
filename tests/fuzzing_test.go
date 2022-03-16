package tests

import (
	"testing"
	"unicode/utf8"

	go118 "github.com/ag9920/go-experiment/go_1_18"
)

// The unit test has limitations, namely that each input must be added to the test by the developer.
// One benefit of fuzzing is that it comes up with inputs for your code, and may identify edge cases that the test cases you came up with didn’t reach.
func TestReverse(t *testing.T) {
	testcases := []struct {
		in, want string
	}{
		{"Hello, world", "dlrow ,olleH"},
		{" ", " "},
		{"!12345", "54321!"},
	}
	for _, tc := range testcases {
		rev := go118.Reverse(tc.in)
		if rev != tc.want {
			t.Errorf("Reverse: %q, want %q", rev, tc.want)
		}
	}
}

// You can keep unit tests, benchmarks, and fuzz tests in the same *_test.go file

// fuzz test for reverse
// func FuzzReverse(f *testing.F) {
// 	testcases := []string{"Hello, world", " ", "!12345"}
// 	for _, tc := range testcases {
// 		f.Add(tc) // Use f.Add to provide a seed corpus
// 	}
// 	f.Fuzz(func(t *testing.T, orig string) {
// 		rev := go118.Reverse(orig)
// 		doubleRev := go118.Reverse(rev)
// 		if orig != doubleRev {
// 			t.Errorf("Before: %q, after: %q", orig, doubleRev)
// 		}
// 		if utf8.ValidString(orig) && !utf8.ValidString(rev) {
// 			t.Errorf("Reverse produced invalid UTF-8 string %q", rev)
// 		}
// 	})
// }

// run go test without fuzzing to check if the seed inputs pass
// go test -v -run FuzzReverse

// Run FuzzReverse with fuzzing, to see if any randomly generated string inputs will cause a failure. This is executed using go test with a new flag, -fuzz.
// go test -v -fuzz FuzzReverse
/*
=== FUZZ  FuzzReverse
fuzz: elapsed: 0s, gathering baseline coverage: 0/3 completed
fuzz: elapsed: 0s, gathering baseline coverage: 3/3 completed, now fuzzing with 12 workers
fuzz: minimizing 37-byte failing input file
fuzz: elapsed: 0s, minimizing
--- FAIL: FuzzReverse (0.09s)
    --- FAIL: FuzzReverse (0.00s)
        fuzzing_test.go:43: Reverse produced invalid UTF-8 string "0\x97\x83\xe50"

    Failing input written to testdata/fuzz/FuzzReverse/d01e791d5baa1131d5593a9730b600b704b14d25bcdc756094f64808ccd17943
    To re-run:
    go test -run=FuzzReverse/d01e791d5baa1131d5593a9730b600b704b14d25bcdc756094f64808ccd17943
FAIL
exit status 1
*/

// A failure occurred while fuzzing, and the input that caused the problem is written to a seed corpus file that will be run the next time go test is called, even without the -fuzz flag.
// To view the input that caused the failure, open the corpus file written to the testdata/fuzz/FuzzReverse directory in a text editor.
// Your seed corpus file may contain a different string, but the format will be the same.
// like the following content
// go test fuzz v1
// string("0僗0")

// The first line of the corpus file indicates the encoding version.
// Each following line represents the value of each type making up the corpus entry.
// Since the fuzz target only takes 1 input, there is only 1 value after the version.

func FuzzReverse(f *testing.F) {
	testcases := []string{"Hello, world", " ", "!12345"}
	for _, tc := range testcases {
		f.Add(tc) // Use f.Add to provide a seed corpus
	}
	f.Fuzz(func(t *testing.T, orig string) {
		rev, err1 := go118.ReverseFixedWithErrorReturned(orig)
		if err1 != nil {
			return
		}
		doubleRev, err2 := go118.ReverseFixedWithErrorReturned(rev)
		if err2 != nil {
			return
		}
		if orig != doubleRev {
			t.Errorf("Before: %q, after: %q", orig, doubleRev)
		}
		if utf8.ValidString(orig) && !utf8.ValidString(rev) {
			t.Errorf("Reverse produced invalid UTF-8 string %q", rev)
		}
	})
}

// Rather than returning, you can also call t.Skip() to stop the execution of that fuzz input.
// Fuzz it with go test -fuzz=Fuzz -fuzztime 30s which will fuzz for 30 seconds before exiting if no failure was found.

// go test -fuzz=FuzzReverse -fuzztime 30s
/*

go test -fuzz=FuzzReverse -fuzztime 30s
fuzz: elapsed: 0s, gathering baseline coverage: 0/7 completed
fuzz: elapsed: 0s, gathering baseline coverage: 7/7 completed, now fuzzing with 12 workers
fuzz: elapsed: 3s, execs: 170223 (56711/sec), new interesting: 33 (total: 40)
fuzz: elapsed: 6s, execs: 457368 (95753/sec), new interesting: 36 (total: 43)
fuzz: elapsed: 9s, execs: 834000 (125550/sec), new interesting: 36 (total: 43)
fuzz: elapsed: 12s, execs: 1155857 (107291/sec), new interesting: 36 (total: 43)
fuzz: elapsed: 15s, execs: 1404915 (82970/sec), new interesting: 37 (total: 44)
fuzz: elapsed: 18s, execs: 1665566 (86932/sec), new interesting: 37 (total: 44)
fuzz: elapsed: 21s, execs: 1943627 (92685/sec), new interesting: 37 (total: 44)
fuzz: elapsed: 24s, execs: 2222022 (92800/sec), new interesting: 38 (total: 45)
fuzz: elapsed: 27s, execs: 2496828 (91605/sec), new interesting: 38 (total: 45)
fuzz: elapsed: 30s, execs: 2761636 (88271/sec), new interesting: 38 (total: 45)
fuzz: elapsed: 30s, execs: 2761636 (0/sec), new interesting: 38 (total: 45)
PASS
ok      github.com/ag9920/go-experiment/tests   30.534s

*/
