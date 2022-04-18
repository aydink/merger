package main

import "fmt"

func compareFolders(destination, source []FileMeta) {

	destinationMap := createMap(destination)

	for _, v := range source {
		if val, ok := destinationMap[v.Hash]; ok {
			fmt.Println("\t+", v.Path, "\t", val.Path)
		} else {
			fmt.Println("----", v.Path)
		}
	}
}
