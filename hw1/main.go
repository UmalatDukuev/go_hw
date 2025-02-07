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

// func printFiles(dir string) {
// 	entryDir, err := os.ReadDir(dir)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	for _, file := range entryDir {
// 		for i := 0; i < count; i++ {
// 			fmt.Print("\t")
// 		}
// 		if file.IsDir() {
// 			fmt.Printf("DIR: %s\n", file.Name())
// 		} else {
// 			fmt.Printf("FILE: %s\n", file.Name())
// 		}
// 		if file.IsDir() {
// 			fileName := file.Name()
// 			str := dir + "/" + fileName
// 			count += 1
// 			printFiles(str)
// 			count -= 1
// 		}
// 	}
// }

var recursionLevel = 0

// func formatStr(file string, recursionLevel) {

// }

func printFiles(dir string) {
	entryDir, err := os.ReadDir(dir)
	// fmt.Println(recursionLevel)
	if err != nil {
		log.Fatal(err)
	}
	eD, _ := os.ReadDir(dir)
	numOfFiles := 0
	dirLen := len(eD)

	for _, file := range entryDir {
		for i := 0; i < recursionLevel; i++ {
			fmt.Print("\t")
		}
		if file.IsDir() {
			numOfFiles += 1
			if numOfFiles != dirLen {
				fmt.Printf("├───%s\n", file.Name())
				fileName := file.Name()
				str := dir + "/" + fileName
				recursionLevel += 1
				printFiles(str)
				recursionLevel -= 1
			} else {
				fmt.Printf("└───%s\n", file.Name())
				fileName := file.Name()
				str := dir + "/" + fileName
				recursionLevel += 1
				printFiles(str)
				recursionLevel -= 1
			}
		} else {

			numOfFiles += 1
			if numOfFiles != dirLen {
				fmt.Printf("├───%s\n", file.Name())
			} else {
				fmt.Printf("└───%s\n", file.Name())
			}
		}
	}
}

func main() {
	// dir, _ := os.Getwd()
	dir := "./testdata"
	// fmt.Println(dir)
	printFiles(dir)
}
