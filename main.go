package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var data map[string][]string

func loadWordMap() {
	data = make(map[string][]string)

	f, err := os.Open("assets/words.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	counter := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		word := strings.ToLower(scanner.Text())
		hashKey := fmt.Sprintf("%s-%d", word[0:1], len(word))
		_, exists := data[hashKey]
		if !exists {
			data[hashKey] = []string{}
		}

		data[hashKey] = append(data[hashKey], word)
		counter += 1
	}

	fmt.Printf("loaded %d words into memory", counter)
}

func comparePhraseWordAndSearchWord(phraseWord, searchWord string) {
	foundWord := true
	for i := 0; i < len(phraseWord); i++ {
		if phraseWord[i:i+1] == "-" {
			continue
		}

		if phraseWord[i:i+1] != searchWord[i:i+1] {
			foundWord = false
			break
		}
	}

	if foundWord {
		fmt.Println(searchWord)
	}
}

func knowFirstLetter(phraseWord string) {
	for _, searchWord := range data[fmt.Sprintf("%s-%d", phraseWord[0:1], len(phraseWord))] {
		comparePhraseWordAndSearchWord(phraseWord, searchWord)
	}
}

func unknownFirstLetter(phraseWord string) {
	for k := range data {
		keyWordCountStr := strings.Split(k, "-")[1]
		keyWordCount, _ := strconv.Atoi(keyWordCountStr)

		if keyWordCount == len(phraseWord) {
			for _, searchWord := range data[k] {
				comparePhraseWordAndSearchWord(phraseWord, searchWord)
			}
		}
	}
}

func main() {
	loadWordMap()
	fmt.Printf("Note: Use '-' for unknown letters\n\n")

	for true {
		var input string

		fmt.Printf("Enter a phrase: ")
		fmt.Scanln(&input)

		phraseWords := strings.Split(input, " ")

		for _, phraseWord := range phraseWords {
			phraseWord = strings.ToLower(phraseWord)

			fmt.Printf("\npossibilities for '%s':\n", phraseWord)

			if phraseWord[0:1] == "-" {
				unknownFirstLetter(phraseWord)
			} else {
				knowFirstLetter(phraseWord)
			}
		}
	}
}
