package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {

	//getting the input
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter location of file: ")
	path, _ := reader.ReadString('\n')
	path = strings.TrimSpace(path)

	fmt.Printf("opening file  %s \n", path)

	read, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
		return
	}

	normalized := strings.Map(normalize, string(read))
	fmt.Println(normalized)

	writingErr := ioutil.WriteFile(path, []byte(normalized), 0644)

	if writingErr != nil {
		log.Fatal(writingErr)
	}
}

func normalize(in rune) rune {
	switch in {
	case '':
		return 's'

	case 'ð':
		return 'd'

	case '':
		return 'z'

	case 'æ':
		return 'c'

	case 'è':
		return 'c'

	default:
		return in
	}
}
