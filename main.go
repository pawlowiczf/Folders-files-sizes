package main

import (
	"flag"
	"fmt"
	"foldersize/utils"
	"log"
)

const (
	BiggestDir   = "biggest-dirs"
	BiggestFiles = "biggest-files"
)

func main() {
	dirP := flag.String("path", ".", "Directory path to look for biggest files")
	amountP := flag.Int("amount", 5, "Get 'amount' of biggest files in dir")
	modeP := flag.String("mode", "biggest-dirs", "Get list of biggest directories within a directory or biggest files within directory")
	flag.Parse()

	fsys := utils.OpenDir(*dirP)

	if *modeP == BiggestFiles {
		fmt.Printf("Biggest files in %s directory. Listing %d files\n", *dirP, *amountP)

		entries, err := utils.GetBiggestFilesSorted(fsys)
		if err != nil {
			log.Fatalln("error utils.GetDirSize", err)
		}

		for a := 0; a < min(*amountP, len(entries)); a++ {
			fmt.Printf("%+v\n", entries[a])
		}
	} else {
		fmt.Printf("Biggest folders in %s directory. Listing %d folders\n", *dirP, *amountP)

		entries, err := utils.GetBiggestDirSorted(fsys)
		if err != nil {
			log.Fatalln("error utils.GetBiggestDirSorted", err)
		}
		for a := 0; a < min(*amountP, len(entries)); a++ {
			fmt.Printf("%+v\n", entries[a])
		}
	}

}
