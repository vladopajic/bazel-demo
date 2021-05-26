package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	filename := flag.String("out", "", "the name of version file")
	flag.Parse()

	f, err := os.Create(*filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	fileContent := strings.ReplaceAll(fileTemplate, "<VERSION>", time.Now().String())
	_, err2 := f.WriteString(fileContent)
	if err2 != nil {
		log.Fatal(err2)
	}

	fmt.Println("Version Generated")
}

var fileTemplate = `package main

const VERSION = "<VERSION>"

`
