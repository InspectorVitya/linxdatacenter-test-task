package main

import (
	"fmt"
	"github.com/inspectorvitya/linxdatacenter-test-task/internal/app"
	"github.com/inspectorvitya/linxdatacenter-test-task/internal/utils/reader"
	"log"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("one argument expected: path of file to read")
	}

	filePath := os.Args[1]
	fileExt := filepath.Ext(filePath)

	var readerConstructor func(file *os.File) (reader.Reader, error)
	switch fileExt {
	case ".csv":
		readerConstructor = reader.NewCsvReader
	case ".json":
		readerConstructor = reader.NewJSONReader
	default:
		log.Fatalf("file extension must be '.csv' or '.json', got '%v'", fileExt)
	}
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	readerFile, err := readerConstructor(f)
	if err != nil {
		log.Fatal(err)
	}

	parser := app.New(readerFile)
	maxPrice, maxRating, err := parser.GetMaxPriceAndRating()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Max price: %v \nMax rating: %v", maxPrice, maxRating)
}
