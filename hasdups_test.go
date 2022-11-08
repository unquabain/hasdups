package hasdups_test

import (
	"hasdups"
	"testing"
)

type TestCase struct {
	Text        string
	MinLength   int
	MinCount    int
	Expectation bool
}

func (tc TestCase) Test(tb testing.TB) {
	tb.Helper()
	actual := hasdups.HasDups(
		tc.Text,
		tc.MinLength,
		tc.MinCount,
	)
	if actual != tc.Expectation {
		tb.Fatalf(`HasDups(text: %q, minLength: %d, minCount: %d) = %v; expected %v`,
			tc.Text,
			tc.MinLength,
			tc.MinCount,
			actual,
			tc.Expectation,
		)
	}
}

var TestCases = []TestCase{
	{`abracadabra`, 1, 4, true},
	// Minimum length is too long
	{`abracadabra`, 2, 4, false},
	// When matches are longer
	{`aabraacaadaabraa`, 2, 4, true},

	{`example text is an example`, 7, 2, true},
	// Minimum length too long
	{`example text is an example`, 8, 2, false},
	// Expect more matches than there are
	{`example text is an example`, 7, 3, false},

	{`bingo bongo bango`, 3, 3, true},
	// More matches than there are
	{`bingo bongo bango`, 3, 4, false},

	// Doesn't error on empty string.
	{``, 2, 3, false},

	// It'll find overlapping matches.
	// Not sure if that's a problem or not.
	{`bbbbbb`, 4, 3, true},
	// It does only find unique matches, even
	// if they overlap.
	{`bbbbbb`, 4, 4, false},
}

func TestHasDups(t *testing.T) {
	for _, tc := range TestCases {
		tc.Test(t)
	}
}
