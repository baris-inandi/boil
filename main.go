package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

var definedLanguages = map[string]string{}
var configPath = "~/.config/boil/cauldron"

func constructLanguage(f os.FileInfo, path string) {
	extensions := strings.Split(f.Name(), ".")[1:]
	if len(extensions) >= 1 {
		for i := range extensions {
			definedLanguages[extensions[i]] = path
		}
	}
}

func getDefinitionPath() string {
	dirname, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(configPath, "~", dirname, 1)
}

func sliceContains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
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

func exists(name string) (bool, error) {
	_, err := os.Stat(name)
	if err == nil {
		return true, nil
	}
	if errors.Is(err, os.ErrNotExist) {
		return false, nil
	}
	return false, err
}

func displayHelp(args []string) {
	if len(args) == 0 || sliceContains(args, "--help") || sliceContains(args, "-h") {
		fmt.Println(`
boil
a touch wrapper to generate boilerplate files.
    `)
		touchHelp, _ := exec.Command("touch", "--help").Output()
		fmt.Println(strings.Replace(string(touchHelp), "touch", "boil", 1))
		os.Exit(0)
	}
}

func copyBoilerplate(src, dst string) (nonEmpty bool, err error) {
	currentPathContents, err := ioutil.ReadFile(dst)
	if string(currentPathContents) != "" {
		splitPath := strings.Split(dst, "/")
		fmt.Println(splitPath[len(splitPath)-1], "is not an empty file, skipping")
		return true, nil
	}
	in, err := os.Open(src)
	if err != nil {
		return
	}
	defer in.Close()
	out, err := os.Create(dst)
	if err != nil {
		return
	}
	defer func() {
		cerr := out.Close()
		if err == nil {
			err = cerr
		}
	}()
	if _, err = io.Copy(out, in); err != nil {
		return
	}
	err = out.Sync()
	return
}

func createFiles(files []string) {
	cmd := exec.Command("touch", files...)
	stderr, _ := cmd.StderrPipe()
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(stderr)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	// add boilerplate if file exists
	for i := range files {
		if !strings.HasPrefix(files[i], "-") { // if not argument
			workingDir, _ := os.Getwd()
			currentPath := workingDir + "/" + files[i]
			isExistent, _ := exists(currentPath)
			if isExistent {
				splitFilename := strings.Split(files[i], ".")
				extension := splitFilename[len(splitFilename)-1]
				splitPath := strings.Split(currentPath, "/")
				nonEmpty, err := copyBoilerplate(definedLanguages[extension], currentPath)
				if err == nil && !nonEmpty {
					fmt.Println(
						"\033[92m"+ // green
							"boil"+
							"\033[0m", // default color
						"generated boilerplate:",
						splitPath[len(splitPath)-1],
					)
				}
			}
		}
	}
}

func main() {
	readLanguages()
	toBeCreated := os.Args[1:]
	displayHelp(toBeCreated)
	createFiles(toBeCreated)
}
