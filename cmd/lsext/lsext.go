package main

import (
	"archive/zip"
	"flag"
	"fmt"
	"log"

	"github.com/koron/go-zipext"
)

func lsext(name string) error {
	zr, err := zip.OpenReader(name)
	if err != nil {
		return err
	}
	defer zr.Close()
	for _, zf := range zr.File {
		ex := zipext.Parse(zf)
		fmt.Printf("%s\n", zf.Name)
		fmt.Printf("  ModTime: %s\n", zf.ModTime())
		fmt.Printf("  (EX)ModTime: %s\n", ex.ModTime())
		fmt.Printf("  (EX)AcTime:  %s\n", ex.AcTime())
		fmt.Printf("  (EX)CrTime:  %s\n", ex.CrTime())
	}
	return nil
}

func main() {
	flag.Parse()
	for _, s := range flag.Args() {
		if err := lsext(s); err != nil {
			log.Fatal(err)
		}
	}
}
