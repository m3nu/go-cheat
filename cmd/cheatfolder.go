package cmd

import (
	"fmt"
	"log"
	"io/ioutil"
	"os"
	"path"
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
