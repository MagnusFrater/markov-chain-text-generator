package markovgenerator

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const (
	maxSequenceLength = 3
)

// MarkovGenerator is a Markov Chain Text Generator.
type MarkovGenerator struct {
	dictionary     map[string][]string
	keyLookup      []string
	sequenceLength int
}

type connections struct {
	dictionary map[string]int
	keyLookup  []string
}

// New returns a new MarkovGenerator.
func New(sequenceLength int) *MarkovGenerator {
	rand.Seed(time.Now().UnixNano())

	if sequenceLength < 1 {
		sequenceLength = 1
	} else if sequenceLength > 3 {
		sequenceLength = 3
	}

	return &MarkovGenerator{
		dictionary:     make(map[string][]string),
		keyLookup:      []string{},
		sequenceLength: sequenceLength,
	}
}

// Add adds the given text to the Dictionary.
func (g *MarkovGenerator) Add(text string) {
	words := strings.Fields(text)
	for i, word := range words {
		g.addWord(word)

		followingWords := []string{}
		for j := 1; j <= g.sequenceLength; j++ {
			delta := i + j

			if delta < len(words) {
				followingWords = append(followingWords, words[delta])
			} else {
				break
			}
		}

		g.addConnection(word, strings.Join(followingWords, " "))
	}
}

// Generate generates text via the Dictionary.
func (g *MarkovGenerator) Generate(numWords int) string {
	word := g.randomWord()
	var passage string = fmt.Sprintf("%s ", word)

	for i := 0; i < numWords/g.sequenceLength; i++ {
		if len(g.dictionary[word]) == 0 {
			word = g.randomWord()
		} else {
			word = g.randomConnection(word)
		}

		passage += fmt.Sprintf("%s ", word)

		// split := strings.Fields(word)
		// word = split[len(split)-1]
	}

	return passage
}

func (g *MarkovGenerator) addWord(word string) {
	if _, ok := g.dictionary[word]; !ok {
		g.dictionary[word] = []string{}
		g.keyLookup = append(g.keyLookup, word)
	}
}

func (g *MarkovGenerator) addConnection(word, connection string) {
	g.dictionary[word] = append(g.dictionary[word], connection)
}

func (g *MarkovGenerator) randomWord() string {
	return g.keyLookup[randNum(0, len(g.keyLookup))]
}

func (g *MarkovGenerator) randomConnection(word string) string {
	return g.dictionary[word][randNum(0, len(g.dictionary[word]))]
}

func randNum(min, max int) int {
	return rand.Intn(max-min) + min
}
