package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

var fullScan bool

func main() {

	fullScan = *flag.Bool("fullscan", false, "Whether to use all file content or just first 1 mb")

	flag.Parse()

	// 1st argument is the directory location
	arg1 := flag.Arg(0)
	fmt.Println("Destination:", arg1)
	arg2 := flag.Arg(1)
	fmt.Println("Source:", arg2)

	fmt.Println("Full scan:", fullScan)

	destinationPath, err := filepath.Abs(arg1)
	if err != nil {
		fmt.Println("Destination path is not valid")
		return
	}

	sourcePath, err := filepath.Abs(arg2)
	if err != nil {
		fmt.Println("Source path is not valid")
		return
	}

	/*
		var folderMeta []FileMeta

		folderMeta, err = createFolderMeta(absolutePath)
		if err != nil {
			fmt.Println(err)
		}

		//printFolderMeta(folderMeta)
		saveFolderMeta(folderMeta, absolutePath+string(os.PathSeparator)+"folderMeta.txt")
	*/

	var destinationMeta []FileMeta
	var sourceMeta []FileMeta

	if _, err := os.Stat(destinationPath + string(os.PathSeparator) + "folderMeta.txt"); err == nil {
		// path/to/whatever exists
		destinationMeta, err = loadFolderMeta(destinationPath + string(os.PathSeparator) + "folderMeta.txt")
		if err != nil {
			fmt.Println(err)
			return
		}
	} else if os.IsNotExist(err) {
		// path/to/whatever does *not* exist
		destinationMeta, err = createFolderMeta(destinationPath)
		if err != nil {
			fmt.Println(err)
			return
		}
	} else {
		// Schrodinger: file may or may not exist. See err for details.
		fmt.Println(err)
		return
	}

	if _, err := os.Stat(sourcePath + string(os.PathSeparator) + "folderMeta.txt"); err == nil {
		// path/to/whatever exists
		sourceMeta, err = loadFolderMeta(sourcePath + string(os.PathSeparator) + "folderMeta.txt")
		if err != nil {
			fmt.Println(err)
			return
		}
	} else if os.IsNotExist(err) {
		// path/to/whatever does *not* exist
		sourceMeta, err = createFolderMeta(sourcePath)
		if err != nil {
			fmt.Println(err)
			return
		}
	} else {
		// Schrodinger: file may or may not exist. See err for details.
		fmt.Println(err)
		return
	}

	compareFolders(destinationMeta, sourceMeta)
}
