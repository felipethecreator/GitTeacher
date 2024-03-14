package main

import (
	"fmt"
	"os"
	"strings"
	"bufio"
)
func main(){
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Insert a git command: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	command, ok := CommandMap[Name(input)]
	if !ok {
		fmt.Printf("REPROVADO! SR-!!!!! VOCÊ REALMENTE ACHA QUE %v EXISTE? PODE TRANCAR!!!!!!\n", input)
		return
	}
	
	
	fmt.Printf("%+v\n", command.description)
	
	
}
