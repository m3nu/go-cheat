package cmd

import (
	"fmt"
	"github.com/spf13/pflag"
	"os"
	"os/user"
	"path"
	"github.com/gobuffalo/packr/v2"
)

var (
	doSearchFlag string
	listCheatsFlag bool
	listCheatFoldersFlag bool
	printVersionFlag bool

	cheatFolders []cheatfolder
	version string
)

func Main() {
	pflag.Parse()
	cheatFolders = collectCheatFolders()
	includedCheatsheets := packr.New("Included cheatsheets", "../assets/cheatsheets")

	// TODO: Validate CLI args

	if doSearchFlag != "" {
		fmt.Println("Searching for", doSearchFlag)
	}

	if listCheatFoldersFlag {
		fmt.Println(cheatFolders)
	}

	if listCheatsFlag {
		for _, f := range cheatFolders {
			f.listCheatSheets()
		}
	}

	if printVersionFlag {
		fmt.Println("Version: ", version)
	}

	if len(pflag.Args()) == 1 {
		cmd := pflag.Args()[0]
		for _, f := range cheatFolders {
			cheatStr, err := f.getCheatsheet(cmd)
			if err == nil {
				fmt.Println(cheatStr)
				os.Exit(0)
			}
		}

		// Try built-in cheatsheets
		cheatStr, err := includedCheatsheets.FindString(cmd)
		if err == nil {
			fmt.Println(cheatStr)
			os.Exit(0)
		}

		// At this point we didn't find any cheat sheet.
		fmt.Println("No cheatsheet found for", cmd)
	}



}

func init() {
	// example with short version for long flag
	pflag.StringVarP(&doSearchFlag, "search", "s", "", "Search cheatsheets for <keyword>")
	pflag.BoolVarP(&listCheatsFlag, "list", "l", false, "List cheatsheets")
	pflag.BoolVarP(&listCheatFoldersFlag, "directories", "d", false, "List directories on $CHEAT_PATH")
	pflag.BoolVarP(&printVersionFlag, "version", "v", false, "Print the version number")
}

// Find the cheatsheet folders available on the system
func collectCheatFolders() []cheatfolder {
	var folders []cheatfolder

	// First add user cheat folder. May be in ~/.cheat or env var CHEAT_USER_DIR
	userCheatFolderPath := os.Getenv("CHEAT_USER_DIR")

	if userCheatFolderPath == "" {
		currentUser, _ := user.Current()
		userCheatFolderPath = path.Join(currentUser.HomeDir, ".cheat")
	}

	f, err := NewCheatFolder(userCheatFolderPath)
	if err == nil {
		folders = append(folders, f)
	}

	// TODO: Deal with CHEAT_PATH
	//cheatPath := os.Getenv("CHEAT_PATH")

	return folders
}
