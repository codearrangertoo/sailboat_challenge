package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
	"time"
)

func main() {

	var inputTimes []string

	if len(os.Args) < 2 {
		log.Fatal("No JSON input.")
	}

	// Parse the input as JSON
	err := json.Unmarshal([]byte(os.Args[1]), &inputTimes)
	if err != nil {
		log.Fatal("Input must be formatted as JSON array. ", err)
	}

	// Parse the start time
	startTime, err := time.Parse("03:04 PM, DAY 2", "08:00 AM, DAY 1")
	if err != nil {
		log.Fatal("Unable to parse start time", err)
	}

	averageMinutes := getAverageMinutes(startTime, inputTimes)
	fmt.Println(averageMinutes)

}

func getAverageMinutes(startTime time.Time, times []string) int {

	minutes := []float64{}

	for _, t := range times {

		parsedTime, err := parseTime(t)
		if err != nil {
			log.Fatal(err)
		}

		delta := parsedTime.Sub(startTime)

		minutes = append(minutes, delta.Minutes())

	}
	averageMinutes := int(math.Round(average(minutes)))

	return averageMinutes
}

func parseTime(timeString string) (time.Time, error) {

	// Define a regex to parse the timeString
	timeFormatRexp := regexp.MustCompile(`^(?P<time>(0[1-9]|1[0-2]):[0-5]\d [AP]M), DAY (?P<day>\d\d?)$`)

	// Test timeString against timeFormatRexp
	if !timeFormatRexp.MatchString(timeString) {
		return time.Now(), errors.New("Unable to parse timeString")
	}

	match := timeFormatRexp.FindStringSubmatch(timeString)

	paramsMap := make(map[string]string)
	for i, name := range timeFormatRexp.SubexpNames() {
		if i > 0 && i <= len(match) && name != "" {
			paramsMap[name] = match[i]
		}
	}

	// Convert day to int
	day, err := strconv.Atoi(paramsMap["day"])
	if err != nil {
		return time.Now(), err
	}

	// Parse the time
	parsedTime, err := time.Parse("03:04 PM, DAY 2", fmt.Sprintf("%s, DAY 1", paramsMap["time"]))
	if err != nil {
		return time.Now(), err
	}

	// Add the days
	parsedTime = parsedTime.AddDate(0, 0, (day - 1))

	return parsedTime, nil
}

func average(x []float64) float64 {
	total := 0.0
	for _, v := range x {
		total += v
	}
	return total / float64(len(x))
}
