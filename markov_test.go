package markov

import (
	"strings"
	"testing"
)

func TestAdd(t *testing.T) {
	var testChain = map[string][]string{
		"now he":   {"is"},
		"he is":    {"gone", "gone"},
		"is gone":  {"she", "for"},
		"gone she": {"said"},
		"she said": {"he"},
		"said he":  {"is"},
		"gone for": {"good"},
		"for good": {},
	}

	chain := New(2, 1)
	chain.Add("now he is gone she said he is gone for good")

	if len(chain.chain) != len(testChain) {
		t.Errorf("Unequal number of prefixes |\tExpected: %v\tActual: %v\n", len(testChain), len(chain.chain))
	}

	for testPrefix, testSuffix := range testChain {
		suffix, ok := chain.chain[testPrefix]
		if !ok {
			t.Errorf("Missing prefix |\tExpected: %v\n", testPrefix)
		}

		if len(suffix) != len(testSuffix) {
			t.Errorf(
				"Unequal number of suffixes (for prefix: %v) |\tExpected: %v\tActual: %v\n",
				testPrefix, len(testSuffix), len(suffix),
			)
		}

		for i := 0; i < len(testSuffix); i++ {
			if suffix[i] != testSuffix[i] {
				t.Errorf("Incorrect suffix order (index: %v) |\tExpected: %v\tActual: %v\n", i, testSuffix[i], suffix[i])
			}
		}
	}
}

func TestGenerate_numWords(t *testing.T) {
	var testCases = []struct {
		numWords         int
		expectedNumWords int
	}{
		// standard usage
		{numWords: 1, expectedNumWords: 1},
		{numWords: 5, expectedNumWords: 5},
		{numWords: 100, expectedNumWords: 100},

		// below 1 word
		{numWords: 0, expectedNumWords: 1},
		{numWords: -1, expectedNumWords: 1},
	}

	chain := New(2, 1)
	chain.Add("now he is gone she said he is gone for good")

	for _, tc := range testCases {
		var corpus = chain.Generate(tc.numWords)
		if len(strings.Fields(corpus)) != tc.expectedNumWords {
			t.Errorf("Expected: %v\tActual: %v\n", tc.numWords, tc.expectedNumWords)
		}
	}
}
