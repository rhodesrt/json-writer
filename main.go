package main

import (
	"fmt"
	"json-writer/writer"
	"os"
	"strconv"
)

func main() {
	numLoops := os.Args[1]
	l, err := strconv.Atoi(numLoops)
	if err != nil {
		fmt.Println("error in loop string to int conversion")
	}

	w := writer.JsonWriter{
		Loop: l,
	}

	err = w.Write()
	if err != nil {
		fmt.Println("error in JsonWriter.Write()")
	}
}
