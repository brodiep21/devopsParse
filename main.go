package main

import (
	"fmt"

	"github.com/brodiep21/GoWithADO/organize"
)

func main() {
	sourceDir := "path/to/source"
	destDir := "path/to/dest"

	err := organize.OrganizeFiles(sourceDir, destDir)
	if err != nil {
		fmt.Println("Error organizing files:", err)
	}

	// http.HandleFunc("/", handler)
	// http.ListenAndServe(":8081", nil)
	// fmt.Println("Hello World!")
}

// func handler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Hello, World")
// }
