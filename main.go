package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	//	"io/ioutil" //TODO: ADD FILE READING AND WRITING(like cat and echo)
	"github.com/djherbis/times"
	"github.com/fatih/color"
)

func main() {
	cliArgs := os.Args

	if len(cliArgs) < 2 {
		color.Red("Usage: cli-fm <expression> <name>")
		return
	}

	color.Cyan("CLI File Manager\n\n")

	path := "."
	if len(cliArgs) > 2 {
		path = cliArgs[2]
	}

	entries, _ := os.ReadDir(path)

	switch cli := cliArgs[1]; cli {
	case "ls":
		for _, entry := range entries {
			absPath := filepath.Join(path, entry.Name())
			t, err := times.Stat(absPath)
			if err != nil {
				log.Fatalf("Erro ao obter informações do arquivo: %v%s\n", err, color.RedString(entry.Name()))
			}
			if entry.IsDir() {
				fmt.Printf("%s/\t\t%s\n", color.BlueString(entry.Name()), t.BirthTime().Format("02/01/2006 15:04:05"))
			} else {
				fmt.Printf("%s\t%s\n", color.GreenString(entry.Name()), t.BirthTime().Format("02/01/2006 15:04:05"))
			}
		}
	case "rm":
		os.Remove(path)
		fmt.Println("File removed:", path)
	case "nwdr":
		os.Mkdir(path, 0755)
		fmt.Println("Directory created:", path)
	case "nw":
		os.Create(path)
		fmt.Println("File created:", path)
	case "help":
		color.Yellow("cli-fm list <path> - List all files and directories in the specified path\n")
		color.Yellow("cli-fm remove <path> - Remove the specified file\n")
		color.Yellow("cli-fm nwdir <path> - Create a new directory in the specified path\n")
		color.Yellow("cli-fm nw <path> - Create a new file in the specified path\n")
		//case "rname": //TODO: Add file renaming
		//case "mv": //TODO: Add file writing
		//case "rf": //TODO: Add file reading
		//case "wf": //TODO: Add file writing
	default:
		color.Red("Invalid command\nTo get help, type: cli-fm help")
	}
}
