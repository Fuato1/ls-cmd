/*
	copy of the "list files in directory" command for UNIX
*/

package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/fatih/color"
)

func main() {
	arg := flag.String("o", "all", "Lists only directories (d) or files (f)")
	flag.Parse()

	files := readFiles(".")

	ls(files, *arg)
}

func readFiles(dir string) []os.FileInfo {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	return files
}

func ls(files []os.FileInfo, arg string) {
	for _, file := range files {
		if strings.Split(file.Name(), "")[0] == "." {
			continue
		} else if arg != "all" {
			if arg == "d" && file.Mode().IsDir() {
				fn := file.Name()
				cl := color.New(color.FgCyan)

				cl.Printf("%s     ", fn)
			} else if arg == "f" && file.Mode().IsRegular() {
				fmt.Printf("%s     ", file.Name())
			}
		} else if arg == "all" {
			if file.Mode().IsDir() {
				fn := file.Name()
				cl := color.New(color.FgCyan)

				cl.Printf("%s     ", fn)
			} else if file.Mode().IsRegular() {
				fmt.Printf("%s     ", file.Name())
			}
		}
	}

	fmt.Println()
}
