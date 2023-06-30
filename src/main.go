package main

import (
	"fmt"
	"os"
	"github.com/tjoseph37/src/normalization"
	"github.com/tjoseph37/src/models"
	"github.com/gocarina/gocsv"
	"log"
)

func getRecords(filePath string) ([]*models.Records, error) {

	f, err := os.Open(filePath)
	fmt.Printf("Opening CSV file %s\n", filePath)

    if err != nil {
        return nil, fmt.Errorf("Error while reading input file %v", err)
    }
    defer f.Close()

	records := []*models.Records{}

    if err := gocsv.UnmarshalFile(f, &records); err != nil {
		return nil, fmt.Errorf("Error while reading input file %v", err)
    }

	// validate csv file is not empty?
    if len(records) == 0 {
		return nil,fmt.Errorf("Error empty CSV file provided")
	}

	return records, nil
}

func writeRecords(records []*models.Records, outputFile string) error {
	file, err := os.Create(outputFile)

	if err != nil {
		log.Fatalf("Failed creating file: %v", err)
	}

	defer file.Close()

	// convert records to [][]string
	err = gocsv.MarshalFile(&records, file)
	if err != nil {
		log.Printf("Error while writing records to CSV file %v", err)
	}
	return nil
}

func main(){
	inputFile := os.Getenv("INPUT")
	if inputFile == ""{
		inputFile = "./test_data/sample.csv"
	}

	outputFile := os.Getenv("OUTPUT")
	if outputFile == ""{
		outputFile = "output.csv"
	}

	fmt.Println("Beginning CSV normalization...")
	normalizedRecords := []*models.Records{}

	if records, err := getRecords(inputFile); err != nil {
		log.Printf("Error while retreiving records %v", err)
	} else {
		fmt.Printf("Normalizing %d records...\n", len(records))
		normalizedRecords = normalization.NormalizeRecords(records)
	}

	if len(normalizedRecords)<1{
		log.Println("No valid CSV records to output!")
	} else {
		writeRecords(normalizedRecords, outputFile)
		fmt.Println("Done!")
	}
}