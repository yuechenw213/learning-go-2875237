package main

import (
	"bufio"
	"fmt"
	"os"
)

const aConst int = 64

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	input, _ := reader.ReadString('\n')
	fmt.Println("You entered:", input)
}
