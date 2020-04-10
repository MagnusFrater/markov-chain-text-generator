package markov

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const maxPrefixLength = 3
const maxSuffixLength = 3

// Chain is a Markov Chain Text Generator.
type Chain struct {
	chain        map[string][]string
	prefixLookup []string

	prefixLength int
	suffixLength int

	allowedRunes map[rune]struct{}
}

// New returns a new Chain.
func New(prefixLength, suffixLength int) *Chain {
	rand.Seed(time.Now().UnixNano())

	if prefixLength < 1 {
		prefixLength = 1
	} else if prefixLength > maxPrefixLength {
		prefixLength = maxPrefixLength
	}

	if suffixLength < 1 {
		suffixLength = 1
	} else if suffixLength > maxSuffixLength {
		suffixLength = maxSuffixLength
	}

	chain := &Chain{
		chain:        make(map[string][]string),
		prefixLookup: []string{},
		prefixLength: prefixLength,
		suffixLength: suffixLength,
	}
	chain.generateAllowedRunes()

	return chain
}

// Add adds the given text to the Chain.
func (c *Chain) Add(corpus string) {
	words := strings.Fields(corpus)
	for i := range words {
		// generate prefix
		prefix := []string{}
		for n := 0; n < c.prefixLength; n++ {
			delta := i + n
			if delta < len(words) {
				cleanedWord := c.cleanWord(words[delta])
				if cleanedWord == "" {
					break
				}

				prefix = append(prefix, cleanedWord)
			}
		}

		// malformed prefix due to 'unclean' word(s)
		// skip suffix to maintain chain health
		// TODO check if any non-zero length prefix is fine
		if len(prefix) != c.prefixLength {
			continue
		}

		// generate suffix
		suffix := []string{}
		for n := 0; n < c.suffixLength; n++ {
			delta := i + c.prefixLength + n
			if delta < len(words) {
				cleanedWord := c.cleanWord(words[delta])
				if cleanedWord == "" {
					break
				}

				suffix = append(suffix, cleanedWord)
			}
		}

		// add to chain
		prefixString := strings.Join(prefix, " ")
		suffixString := strings.Join(suffix, " ")

		if len(suffix) == 0 && len(suffixString) > 0 {
			fmt.Println("suffix: we have a problem here chief")
		}

		c.createPrefix(prefixString)

		if len(suffixString) > 0 {
			c.addSuffix(prefixString, suffixString)
		}
	}
}

// Generate generates text simulating the chain.
func (c *Chain) Generate(numWords int) string {
	// for prefix := range c.chain {
	// 	fmt.Printf("Prefix: '%s'\n", prefix)
	// 	fmt.Printf("Suffixes: %s\n\n", strings.Join(c.chain[prefix], " | "))
	// }
	// return ""

	corpus := []string{}

	// add prefix if there's room in the corpus
	prefix := c.randomPrefix()
	prefixParts := strings.Fields(prefix)
	for _, word := range prefixParts {
		if len(corpus) < numWords {
			corpus = append(corpus, word)
		}
	}

	for len(corpus) < numWords {
		// check if suffixes exist
		suffix := c.randomSuffix(prefix)
		suffixParts := strings.Fields(suffix)
		if suffix == "" {
			// no suffixes
			prefix = c.randomPrefix()
			continue
		}

		// add suffix if there's room in the passage
		if len(corpus) < numWords {
			corpus = append(corpus, suffix)
		}

		// create new prefix
		newPrefix := []string{}

		// new-prefix might contain the last bits of old-prefix depending on prefix-length
		for i := len(prefixParts) - c.prefixLength + len(suffixParts); i < len(prefixParts); i++ {
			newPrefix = append(newPrefix, prefixParts[i])
		}

		// new-prefix must always contain last-suffix
		for _, word := range suffixParts {
			newPrefix = append(newPrefix, word)
		}

		// reset prefix
		prefix = strings.Join(newPrefix, " ")
	}

	return strings.Join(corpus, " ")
}

func (c *Chain) generateAllowedRunes() {
	c.allowedRunes = map[rune]struct{}{
		'.': {}, '!': {}, '?': {},
		',': {}, '\'': {},
	}

	// lowercase letters
	for r := 'a'; r <= 'z'; r++ {
		c.allowedRunes[r] = struct{}{}
	}

	// uppercase letters
	for r := 'A'; r <= 'Z'; r++ {
		c.allowedRunes[r] = struct{}{}
	}
}

func (c *Chain) cleanWord(word string) string {
	return strings.ToLower(strings.Map(
		func(r rune) rune {
			if _, ok := c.allowedRunes[r]; ok {
				return r
			}

			return -1
		},
		word,
	))
}

func (c *Chain) createPrefix(prefix string) {
	if _, ok := c.chain[prefix]; !ok {
		c.chain[prefix] = []string{}
		c.prefixLookup = append(c.prefixLookup, prefix)
	}
}

func (c *Chain) addSuffix(prefix, suffix string) {
	c.chain[prefix] = append(c.chain[prefix], suffix)
}

func (c *Chain) randomPrefix() string {
	return c.prefixLookup[randNum(0, len(c.prefixLookup))]
}

func (c *Chain) randomSuffix(prefix string) string {
	suffixes := c.chain[prefix]

	if len(suffixes) == 0 {
		return ""
	}

	return suffixes[randNum(0, len(suffixes))]
}

func randNum(min, max int) int {
	return rand.Intn(max-min) + min
}
