package main

import (
	"fmt"
	"log"

	tio "ghora.net/tinyml/io"
)

func main() {
	data, err := tio.ReadCSV("../assets/datasets/dummy.csv")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(data)
}
