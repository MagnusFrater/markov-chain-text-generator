package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	markovgenerator "github.com/MagnusFrater/markov-chain-text-generator"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Usage: ./cmd/generator path/to/input")
	}

	buf, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	generator := markovgenerator.New(2)
	generator.Add(string(buf))
	fmt.Println(generator.Generate(100))
}
