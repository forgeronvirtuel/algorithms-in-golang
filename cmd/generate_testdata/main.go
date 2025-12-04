package main

import (
	"fmt"
	"log"

	"github.com/forgeronvirtuel/algorithms-in-golang/internal/sortalgo"
)

func main() {
	sizes := []int{10, 50, 100, 500, 1000}
	testData := sortalgo.GenerateTestData(sizes)

	outputFile := "internal/sortalgo/generated_testdata.go"
	err := sortalgo.WriteTestDataToFile(outputFile, testData)
	if err != nil {
		log.Fatalf("Error writing test data: %v", err)
	}

	fmt.Printf("Test data generated successfully in %s\n", outputFile)
}
