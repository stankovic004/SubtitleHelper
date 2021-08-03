package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
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

	//geting the path
	fmt.Print("enter your file directory: ")
	var path, _ = reader.ReadString('\n')
	path = strings.TrimSpace(path)

	//input
	input, inputErr := ioutil.ReadFile(path)

	fmt.Printf("opening file  %s \n", path)
	var file, err = os.OpenFile(path, os.O_RDWR, 0644)

	output := bytes.Replace(input, []byte("hellloou"), []byte("muhuhuhuhuhuaa"), -1)

	if inputErr = ioutil.WriteFile(path, output, 0666); inputErr != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if isError(err) {
		return
	}

	defer file.Close()
}
