package sortalgo

import (
	"fmt"
	"math/rand"
	"os"
	"sort"
)

type TestData struct {
	Unsorted []int
	Sorted   []int
}

func GenerateTestData(sizes []int) []TestData {
	testData := make([]TestData, len(sizes))

	for i, size := range sizes {
		unsorted := make([]int, size)
		for j := range unsorted {
			unsorted[j] = rand.Intn(size * 10)
		}

		sorted := make([]int, size)
		copy(sorted, unsorted)
		sort.Ints(sorted)

		testData[i] = TestData{
			Unsorted: unsorted,
			Sorted:   sorted,
		}
	}

	return testData
}

func WriteTestDataToFile(filename string, testData []TestData) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	f.WriteString("package sortalgo\n\n")
	f.WriteString("var GeneratedTestData = []struct {\n")
	f.WriteString("\tUnsorted []int\n")
	f.WriteString("\tSorted   []int\n")
	f.WriteString("}{\n")

	for _, data := range testData {
		f.WriteString("\t{\n")
		f.WriteString(fmt.Sprintf("\t\tUnsorted: %#v,\n", data.Unsorted))
		f.WriteString(fmt.Sprintf("\t\tSorted:   %#v,\n", data.Sorted))
		f.WriteString("\t},\n")
	}

	f.WriteString("}\n")
	return nil
}
