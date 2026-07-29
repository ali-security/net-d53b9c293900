package main

import (
	"fmt"
	"os"

	"golang.org/x/mod/module"
	"golang.org/x/mod/zip"
)

func main() {
	if len(os.Args) != 5 {
		fmt.Fprintln(os.Stderr, "usage: create_mod <module-path> <version> <dir> <out.zip>")
		os.Exit(1)
	}
	m := module.Version{Path: os.Args[1], Version: os.Args[2]}
	out, err := os.Create(os.Args[4])
	if err != nil {
		fmt.Fprintln(os.Stderr, "create:", err)
		os.Exit(1)
	}
	defer out.Close()
	if err := zip.CreateFromDir(out, m, os.Args[3]); err != nil {
		fmt.Fprintln(os.Stderr, "CreateFromDir:", err)
		os.Exit(1)
	}
	fmt.Println("wrote", os.Args[4])
}
