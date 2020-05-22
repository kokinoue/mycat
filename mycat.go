package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func main() {
	// -nオプションで番号つける
	var n = flag.Bool("n", false, "行番号")
	flag.Parse()

	var (
		files = flag.Args()
		line  = 0
	)

	for x := 0; x < len(files); x++ {
		sf, err := os.Open(files[x])
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		} else {
			scanner := bufio.NewScanner(sf)
			for scanner.Scan() {
				line++
				if *n {
					fmt.Printf("%v: ", line)
				}
				fmt.Println(scanner.Text())
			}
		}
	}
}
