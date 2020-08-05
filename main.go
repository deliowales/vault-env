package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

//filerDir is a wraper of WalkFunc that removes directoies from the output
func filterDir(files *[]string) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatal(err)
		}

		if info.IsDir() {
			return nil
		}

		*files = append(*files, path)
		return nil
	}
}

//outputEnviromentVariable takes the content of the projected secret and outputs to STDOUT
func outputEnviromentVariable(file *string) {
	fileName := strings.Split(*file, "/")
	content, err := os.Open(*file)
	if err != nil {
		log.Fatal(err)
	}
	defer content.Close()

	b, err := ioutil.ReadAll(content)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s=%s", strings.ToUpper(fileName[len(fileName)-1]), b)
}

func main() {
	var files []string

	secretPath := os.Getenv("SECRETPATH")

	err := filepath.Walk(secretPath, filterDir(&files))
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		outputEnviromentVariable(&file)
	}
}
