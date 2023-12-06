package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Person struct {
	Name  string
	Age   int
	Score int
}

func main() {
	paramsLength := len(os.Args)
	if paramsLength < 3 {
		fmt.Println("You should provide input file path and output file path")
		return
	}
	inputFilePath := os.Args[1]
	outputFilePath := os.Args[2]
	outputFileName := strings.Split(outputFilePath, "/")[1]
	outputFileName = strings.Split(outputFileName, ".")[0]

	records, err := readFile(inputFilePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	fileWithoutHeader := records[1:]

	persons := MountPersons(fileWithoutHeader)
	orderedByName := OrderByName(persons)
	orderedByAge := OrderByAge(persons)

	err = WriteFile("./"+outputFileName+"_ordenado_por_nome.csv", orderedByName)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = WriteFile("./"+outputFileName+"_ordenado_por_idade.csv", orderedByAge)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func readFile(filePath string) ([][]string, error) {
	inputFile, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer inputFile.Close()

	reader := csv.NewReader(inputFile)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}
	return records, nil
}

func MountPersons(records [][]string) []Person {
	persons := make([]Person, len(records))
	for i, record := range records {
		age, err := strconv.Atoi(record[1])
		if err != nil {
			fmt.Println(err)
			return nil
		}
		score, err := strconv.Atoi(record[2])
		if err != nil {
			fmt.Println(err)
			return nil
		}
		persons[i] = Person{
			Name:  record[0],
			Age:   age,
			Score: score,
		}
	}
	return persons
}

func OrderByName(persons []Person) []Person {
	orderedPersonsByName := make([]Person, len(persons))
	copy(orderedPersonsByName, persons)
	sort.Slice(orderedPersonsByName, func(i, j int) bool {
		upperNameI := strings.ToUpper(orderedPersonsByName[i].Name)
		upperNameJ := strings.ToUpper(orderedPersonsByName[j].Name)
		return upperNameI < upperNameJ
	})
	return orderedPersonsByName
}

func OrderByAge(persons []Person) []Person {
	orderedPersonsByAge := make([]Person, len(persons))
	copy(orderedPersonsByAge, persons)
	sort.Slice(orderedPersonsByAge, func(i, j int) bool {
		return orderedPersonsByAge[i].Age < orderedPersonsByAge[j].Age
	})
	return orderedPersonsByAge
}

func WriteFile(outputFilePath string, records []Person) error {
	outputFile, err := os.Create(outputFilePath)
	if err != nil {
		fmt.Println("Writing file error create...")
		return err
	}
	defer outputFile.Close()

	data := [][]string{
		{"Nome", "Idade", "Pontuação"},
	}
	writer := csv.NewWriter(outputFile)
	for _, record := range records {
		data = append(data, []string{record.Name, strconv.Itoa(record.Age), strconv.Itoa(record.Score)})
	}
	err = writer.WriteAll(data)
	if err != nil {
		fmt.Println("Writing file error write...")
		return err
	}

	defer writer.Flush()
	err = writer.Error()
	if err != nil {
		fmt.Println("Writing file error...")
		return err
	}

	return nil
}
