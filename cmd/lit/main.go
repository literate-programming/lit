// GENERATED BY LIT v0.3.0
// DO NOT EDIT BY HAND

// # lit cli

// the cli is really lightweight and small in implementation.

// at first we check if the first arg is `version`. if it ism then print the version of the tool to `stdout`

package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/literate-programming/lit"
)

var version = "dev"

func main() {
	totalArgs := len(os.Args)
	if totalArgs == 1 {
		fmt.Printf("\n")
		fmt.Printf("  +--------+\n")
		fmt.Printf("  | ---    |\n")
		fmt.Printf("  | ---    |  lit v%s - literate preprocessor\n", version)
		fmt.Printf("  |   ---  |  --- ------ - ---------------------")
		fmt.Printf("  |   ---  |\n")
		fmt.Printf("  | ---    |\n")
		fmt.Printf("  +--------+\n")
		fmt.Printf("")
		fmt.Printf("  lit <in> <out>")
		fmt.Printf("  lit foo.go.lit foo.go\n")
		os.Exit(1)
	}

var err error
var inputData []byte

if totalArgs >= 1 {
	if os.Args[1] == "version" {
		fmt.Println("lit v"+version)
		os.Exit(0)
	}

// read the given filepath and store the data to a variable

	inputData, err = ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// now we can transform the input data

transformed := lit.Transform(inputData, []byte("// "))


// last step is to write the file back to the given filepath

if totalArgs == 3 {
	outPath := os.Args[2]
	outPre := []byte("// GENERATED BY LIT v"+version+"\n// DO NOT EDIT BY HAND\n\n")
	err := ioutil.WriteFile(outPath, append(outPre, transformed...), 0777)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// TODO: verbose mode
	// fmt.Println("successful written file to", outPath)
} else {
	fmt.Println(string(transformed))
}


}
