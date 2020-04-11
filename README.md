# Markov Chain Text Generator

Automatically generates text that mimics the style of the input text.

## Example

```go
func generate() {
  generator := markov.New(2, 1)
  generator.Add("now he is gone she said he is gone for good")
  fmt.Println(generator.Generate(8))
  // sample output: "gone she said he is gone she said"
}
```

## Test it out!

`go run ./cmd/generate -file=path/to/file`

Flag Options:
* `-file=path/to/file`
  * path to text file to add to the Markov Chain
  * **(required if no -str)**
  * string
* `-str="According to all known laws of aviation, ..."`
  * string to add to the Markov Chain
  * **(required if no -file)**
  * string
* `-prefixLength=1`
  * length of the Markov Chain prefixes
  * integer
  * range: [1,3]
  * default: 2
* `-suffixLength=3`
  * length of the Markov Chain suffixes
  * integer
  * range: [1,3]
  * default: 2
* `-numWords=69`
  * number of words to generate
  * integer
  * range: [1,infinity]
  * default: 100
