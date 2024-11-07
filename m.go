package main

import (
	"bufio"
	"os"
	"strings"
)

type Measurements struct {
	Nome  string
	Min   float64
	Max   float64
	Sum   float64
	Count int64
}

func main() {
	measurements, err := os.Open("measurements.txt")

	if err != nil {
		panic(err)
	}

	defer measurements.Close()

	scanner := bufio.NewScanner(measurements)

	for scanner.Scan() {
		rawData := scanner.Text()
		semicolon := strings.Index(rawData, ";")
		location := rawData[:semicolon]
		temp := rawData[semicolon+1:]
	}

}
