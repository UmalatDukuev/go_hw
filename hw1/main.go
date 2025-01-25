package main

import (
	"fmt"
	"log"
	"os"
)

// func dirTree(out *bytes.Buffer, path string, printFiles bool) error {

// 	return nil
// }

// func main() {
// 	out := os.Stdout

// 	if !(len(os.Args) == 2 || len(os.Args) == 3) {
// 		panic("usage go run main.go . [-f]")
// 	}
// 	path := os.Args[1]
// 	printFiles := len(os.Args) == 3 && os.Args[2] == "-f"

// 	err := dirTree(out, path, printFiles)
// 	if err != nil {
// 		panic(err.Error())
// 	}
// }

var count = 0

func printFiles(dir string) {
	entryDir, err := os.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range entryDir {
		for i := 0; i < count; i++ {
			fmt.Print("\t")
		}
		if file.IsDir() {
			fmt.Printf("DIR: %s\n", file.Name())
		} else {
			fmt.Printf("FILE: %s\n", file.Name())
		}
		if file.IsDir() {
			fileName := file.Name()
			str := dir + "/" + fileName
			count += 1
			printFiles(str)
			count -= 1
		}
	}
}

func main() {
	dir, _ := os.Getwd()
	// fmt.Println(dir)
	printFiles(dir)
}
