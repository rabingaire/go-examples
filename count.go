package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readFromFile(filePath string) []byte {
	dat, err := ioutil.ReadFile(filePath)
	check(err)
	return dat
}

func convertStringToArray(splitText *[]string, filePath string) {
	splitArray := strings.Split(string(readFromFile(filePath)), "\n")
	for i := 0; i < len(splitArray); i++ {
		splitMore := strings.Split(splitArray[i], " ")
		for j := 0; j < len(splitMore); j++ {
			*splitText = append(*splitText, splitMore[j])
		}
	}
}

func findStringPosition(position *[]int, splitText []string, arg string) {
	for i := 0; i < len(splitText); i++ {
		if splitText[i] == arg {
			*position = append(*position, i+1)
		}
	}
}

func main() {
	filePath := os.Args[1]
	arg := os.Args[2]
	position := make([]int, 0)
	splitText := make([]string, 0)
	convertStringToArray(&splitText, filePath)
	findStringPosition(&position, splitText, arg)
	fmt.Println(position)
}
