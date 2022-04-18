package main

/*
import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

type FileMeta struct {
	FileName string
	Path     string
	Hash     string
	Size     int64
	ModTime  int64
}

var folderMeta []FileMeta

func getFileMeta(path string) (FileMeta, error) {

	meta := FileMeta{}
	absPath, err := filepath.Abs(path)

	if err != nil {
		return meta, err
	}

	meta.FileName = filepath.Base(path)
	meta.Path = absPath

	//fmt.Println(path)
	infile, err := os.Open(path)

	if err != nil {
		return meta, err
	}

	stat, err := infile.Stat()
	if err != nil {
		return meta, err
	}

	hash := md5.New()
	io.Copy(hash, infile)

	marker := make([]byte, 2)
	_, err = infile.ReadAt(marker, 0)

	if err != nil {
		fmt.Println(err)
	}

	meta.Hash = hex.EncodeToString(hash.Sum([]byte("")))
	meta.Size = stat.Size()
	meta.ModTime = stat.ModTime().Unix()

	return meta, nil
}

func processFile(path string, f os.FileInfo, err error) error {
	if !f.IsDir() {

		meta, err := getFileMeta(path)

		if err != nil {
			fmt.Println(err)
		}

		folderMeta = append(folderMeta, meta)
	}

	return nil
}

func saveFolderMeta(folderMeta []FileMeta, path string) {
	f, err := os.OpenFile(path, os.O_CREATE, 0777)
	defer f.Close()

	if err != nil {
		fmt.Println(err)
	}

	for i := range folderMeta {
		fmt.Fprintf(f, "%s\t", folderMeta[i].Hash)
		fmt.Fprintf(f, "%s\t", folderMeta[i].FileName)
		fmt.Fprintf(f, "%s\t", folderMeta[i].Path)
		fmt.Fprintf(f, "%d\t", folderMeta[i].ModTime)
		fmt.Fprintf(f, "%d\n", folderMeta[i].Size)
	}
}
*/
