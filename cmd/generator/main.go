package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	markov "github.com/MagnusFrater/markov-chain-text-generator"
)

func main() {
	file := flag.String("file", "", "path to text file")
	str := flag.String("str", "", "string to add")
	prefixLength := flag.Int("prefixLength", 2, "length of the Markov Chain's prefixes")
	suffixLength := flag.Int("suffixLength", 2, "length of the Markov Chain's suffixes")
	numWords := flag.Int("numWords", 100, "number of words to generate")
	flag.Parse()

	if *file == "" && *str == "" {
		flag.Usage()
		os.Exit(1)
	}

	var corpus string
	if *file != "" {
		buf, err := ioutil.ReadFile(*file)
		if err != nil {
			log.Fatal(err)
		}
		corpus = string(buf)
	} else if *str != "" {
		corpus = *str
	} else {
		flag.Usage()
		os.Exit(1)
	}

	generator := markov.New(*prefixLength, *suffixLength)
	generator.Add(corpus)
	fmt.Println(generator.Generate(*numWords))
}
