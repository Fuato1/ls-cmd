/*
	copy of the "list files in directory" command for UNIX
*/

package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/fatih/color"
)

func main() {
	files, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if file.Mode().IsDir() {
			fn := file.Name()
			cl := color.New(color.FgCyan)

			cl.Printf("%s     ", fn)
		} else if file.Mode().IsRegular() {
			fmt.Printf("%s     ", "file")
		}
	}
	fmt.Println()
}
