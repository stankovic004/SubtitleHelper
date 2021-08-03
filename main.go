package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}
	return (err != nil)
}

func main() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("enter your file directory: ")
	var path, _ = reader.ReadString('\n')

	path = strings.TrimSpace(path)

	fmt.Printf("opening file... %s", path)
	var file, err = os.OpenFile(path, os.O_RDWR, 0644)

	if isError(err) {
		return
	}

	defer file.Close()
}
