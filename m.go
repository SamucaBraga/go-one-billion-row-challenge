package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Measurements struct {
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

	dados := make(map[string]Measurements)

	scanner := bufio.NewScanner(measurements)

	for scanner.Scan() {
		rawData := scanner.Text()
		semicolon := strings.Index(rawData, ";")
		location := rawData[:semicolon]
		rawTemp := rawData[semicolon+1:]

		temp, _ := strconv.ParseFloat(rawTemp, 64)

		measurements, ok := dados[location]
		if !ok {
			measurements = Measurements{
				Min:   temp,
				Max:   temp,
				Sum:   temp,
				Count: 1,
			}
		} else {
			measurements.Min = min(measurements.Min, temp)
			measurements.Min = max(measurements.Max, temp)
			measurements.Sum += temp
			measurements.Count++
		}

		dados[location] = measurements
	}

	locations := make([]string, 0, len(dados))
	for name := range dados {
		locations = append(locations, name)
	}

	sort.Strings(locations)

	fmt.Printf("{")
	for _, name := range locations {
		measurements := dados[name]
		fmt.Printf(
			"%s=%.1f/%.1f/%.1f, ",
			name,
			measurements.Min,
			measurements.Sum/float64(measurements.Count),
			measurements.Max,
		)
	}
	fmt.Printf("}")
}
