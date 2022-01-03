package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type language struct {
	filepath   string
	extensions []string
}

var definedLanguages = []language{}
var configPath = "~/code/boil/cauldron"

func constructLanguage(f os.FileInfo, path string) {
	extensions := strings.Split(f.Name(), ".")[1:]
	if len(extensions) >= 1 {
		definedLanguages = append(definedLanguages, language{
			filepath:   path,
			extensions: extensions,
		})
	}
}

func getDefinitionPath() string {
	dirname, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(configPath, "~", dirname, 1)
}

func readLanguages() {
	err := filepath.Walk(getDefinitionPath(), func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err)
			return err
		}
		constructLanguage(info, path)
		return nil
	})
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	readLanguages()
	toBeCreated := os.Args[1:]
	if len(toBeCreated) == 0 {
		fmt.Println(`
  boil
  Create boilerplate files.

  USAGE:
  boil <filename>
    `)
	}
	for i := range toBeCreated {
		fmt.Println(toBeCreated[i])
	}
}
