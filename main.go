package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func PrintFile(filename string) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	line := ""
	aline := ""
	for i := 0; i < len(bytes); i++ {
		b := bytes[i]

		// ASCII文字以外は★を表示
		c := string(b)
		if b < 32 || b > 126 {
			c = "★"
		}
		aline += c

		m := i % 16
		if m == 0 {
			line += fmt.Sprintf("%08d: ", i)
		}
		line += fmt.Sprintf("%02x", b)
		switch m {
		case 15:
			fmt.Println(line + "|" + aline)
			line = ""
			aline = ""
		default:
			line += " "
		}
	}

	if line != "" {
		fmt.Printf("%-53s|%s\n", line, aline)
	}

}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("The number of arguments specified is incorrect.")
		os.Exit(1)
	} else {
		file := os.Args[1]
		ext := filepath.Ext(file)
		if ext != ".txt" {
			fmt.Println("Please specify *.txt.")
			os.Exit(1)
		}

		PrintFile(file)
	}
}
