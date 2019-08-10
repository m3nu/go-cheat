package cmd

import (
	"fmt"
	"github.com/alecthomas/chroma/formatters"
	"github.com/alecthomas/chroma/lexers"
	"github.com/alecthomas/chroma/styles"
	"github.com/spf13/pflag"
	"os"
	"os/user"
	"path"
	"github.com/gobuffalo/packr/v2"
	"strings"
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
		for _, cheatFolder := range cheatFolders {
			cheatFolder.search(doSearchFlag)
		}
	}

	if listCheatFoldersFlag {
		for _, f := range cheatFolders {
			fmt.Println(f.path)
		}
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
				if os.Getenv("CHEAT_COLORS") == "true" {
					printColoredCheatsheet(cheatStr)
				} else {
					fmt.Println(cheatStr)
				}
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

	// Add folders found in CHEAT_PATH
	cheatPath := os.Getenv("CHEAT_PATH")
	if cheatPath != "" {
		for _, p := range strings.Split(cheatPath, ":") {
			f, err := NewCheatFolder(p)
			if err == nil {
				folders = append(folders, f)
			}
		}
	}

	return folders
}

func printColoredCheatsheet(s string) {
	colorscheme := os.Getenv("CHEAT_COLORSCHEME")
	if colorscheme == "" {
		colorscheme = "pygments"
	}

	style := styles.Get(colorscheme)
	lexer := lexers.Get("bash")

	formatter := formatters.Get("terminal256")
	if formatter == nil {
		formatter = formatters.Fallback
	}

	iterator, _ := lexer.Tokenise(nil, s)
	formatter.Format(os.Stdout, style, iterator)
}
