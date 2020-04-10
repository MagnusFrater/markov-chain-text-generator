package markov

import "testing"

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
