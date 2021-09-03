package main

import (
	"bufio"
	"fmt"
	"os"
)

// func revertString(words string) string {
// 	wordsSlice := strings.Split(words, " ")
// 	length := len(wordsSlice)
// 	revertedWordsSlice := []string{}
// 	for i := length - 1; i >= 0; i-- {
// 		revertedWordsSlice = append(revertedWordsSlice, wordsSlice[i])
// 	}
// 	revertedWords := strings.Join(revertedWordsSlice, " ")
// 	return revertedWords
// }

// input: "I am Chinese"
// output: "Chinese am I"

func revertUnit(word []byte) {
	length := len(word)
	for i := 0; i <= (length-1)/2; i++ {
		word[i], word[length-1-i] = word[length-1-i], word[i]
	}
}

func revertString(word []byte) []byte {
	length := len(word)
	if length == 1 || length == 0 {
		return word
	}
	spliter := byte(' ')
	var (
		first, rear int
	)
	first = 0
	rear = first + 1
	for ; rear < length; rear++ {
		if word[rear] == spliter {
			revertUnit(word[first:rear])
			first = rear + 1
			rear = first
			continue
		}
	}
	revertUnit(word)
	return word
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	word, _, _ := reader.ReadLine()
	fmt.Println(string(revertString(word)))
}
