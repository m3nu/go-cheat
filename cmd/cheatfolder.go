package cmd

import (
	"bufio"
	"fmt"
	"log"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"sync"
)

type cheatfolder struct {
	path string
}

func NewCheatFolder(path string) (cheatfolder, error) {
	f := cheatfolder{path: path}
	if f.folderExists() {
		return f, nil
	} else {
		return f, fmt.Errorf("folder doesn't exist")
	}
}

func (f *cheatfolder) folderExists() bool {
	_, err := os.Stat("userCheatFolder")
	return os.IsNotExist(err)
}

// Return cheatsheet fo given name
func (f *cheatfolder) getCheatsheet(cmd string) (string, error) {
	cheatSheetPath := path.Join(f.path, cmd)

	cheatSheetBytes, err := ioutil.ReadFile(cheatSheetPath)
	if err != nil {
		return "", err
	}
	return string(cheatSheetBytes), nil
}

//List all cheat sheets in this folder
func (f *cheatfolder) listCheatSheets() {
	files, err := ioutil.ReadDir(f.path)
	if err != nil {
		log.Fatal(err)
	}

	for _, sheet := range files {
		fmt.Printf("%-25v%v\n", sheet.Name(), f.path)
	}
}

// look for a specific search term inside the cheatsheet
func (f *cheatfolder) search(term string) {
	files, _ := ioutil.ReadDir(f.path)
	var wg sync.WaitGroup
	for _, fp := range files {
		filename := fp.Name()
		wg.Add(1)
		go func() {
			lineRes := []string{}
			f, _ := os.Open(path.Join(f.path, filename))
			defer f.Close()
			scanner := bufio.NewScanner(f)

			for scanner.Scan() {
				if strings.Contains(scanner.Text(), term) {
					lineRes = append(lineRes, scanner.Text())
				}
			}
			if len(lineRes) > 0 {
				out := fmt.Sprintf("%v:\n", filename)
				for _, line := range lineRes {
					out += fmt.Sprintf("  %v\n", line)
				}
				fmt.Println(out)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}
