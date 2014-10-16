package main

import (
	"bufio"
	"fmt"
	"os"
)

func getInput(question string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("%s: ", question)
	text, _ := reader.ReadString('\n')
	return text
}
