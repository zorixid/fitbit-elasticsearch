package fitbit

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

// How the data comes in from the FitBit API
type FitBitData struct {
	Steps    []DailyStats `json:"activities-steps"`
	Distance []DailyStats `json:"activities-distance"`
	Floors   []DailyStats `json:"activities-floors"`
}

type DailyStats struct {
	DateTime string `json:"dateTime"`
	Value    string `json:"value"`
}

// How we're going to store the new data
type AllStats struct {
	DayStats []DayStats `json:"AllStats"`
}

type DayStats struct {
	Date     string `json:"date"`
	Distance string `json:"distance"`
	Floors   string `json:"floors"`
	Steps    string `json:"steps"`
}

// Gets daily distance from fitbit API
func GetDistance(date DatePair) FitBitData {
	// create api URL
	path := "activities/distance"
	url := DefineEndpoint(path, date)

	// get the data for the specificed date range and print to console
	data := GetActivityOnDate(url)
	PrintData(data, path)

	return UnmarshalData(data)
}

func GetFloors(date DatePair) FitBitData {
	path := "activities/floors"
	url := DefineEndpoint(path, date)

	data := GetActivityOnDate(url)
	PrintData(data, path)

	return UnmarshalData(data)
}

func GetSteps(date DatePair) FitBitData {
	path := "activities/steps"
	url := DefineEndpoint(path, date)

	data := GetActivityOnDate(url)
	PrintData(data, path)

	return UnmarshalData(data)
}

func UnmarshalData(rawData []byte) FitBitData {
	var a FitBitData
	err := json.Unmarshal(rawData, &a)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}

	return a
}

// Builds a FitBit endpoint URL from a resourcePath and Date
func DefineEndpoint(resourcePath string, date DatePair) string {
	return fmt.Sprintf("https://api.fitbit.com/1/user/-/%s/date/%s/%s.json", resourcePath, date.Start, date.End)
}

// Prints out FitBit data instead of sending it back right now
func PrintData(data []byte, path string) {
	if data != nil {
		fmt.Printf("%s: %v\n\n", path, string(data))
	} else {
		fmt.Println("Something went wrong")
	}
}

// save FitBit data in a format that works for Elasticsearch
func CombineData(d []DailyStats, f []DailyStats, s []DailyStats) AllStats {

	// Create a temporary to save the current data in
	allStats := AllStats{}
	// returnedData.APIendpoint = dataName

	// For each day of data, append the data to the returnedData.DailyStats
	for _, each_day := range d {
		temp := DayStats{}
		temp.Date = each_day.DateTime
		temp.Distance = each_day.Value
		temp.Floors = Find(f, temp.Date)
		temp.Steps = Find(s, temp.Date)

		allStats.DayStats = append(allStats.DayStats, temp)
	}

	// fmt.Printf("CombinedData: %v\n", allStats)
	return allStats
}

func Find(a []DailyStats, d string) string {
	for _, s := range a {
		if d == s.DateTime {
			return s.Value
		}
	}
	return "0"
}

func WriteData(d AllStats) {

	fmt.Println("Starting file write.")
	sum := 0

	// If the file doesn't exist, create it, or append to the file
	f, err := os.OpenFile("test.json", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	for _, s := range d.DayStats {
		// fmt.Printf("WriteData: %+v\n", s)

		// { "index" : { "_index" : "test", "_id" : "1" } }
		meta := []byte(fmt.Sprintf(`{ "index" : { "_index" : "%s", "_id" : "%s" } }%s`, "test_index", s.Date, "\n"))

		if _, err := f.Write(meta); err != nil {
			f.Close() // ignore error; Write error takes precedence
			log.Fatal(err)
		}

		// marshal the struct into a []byte array, then append a newline
		data, _ := json.Marshal(s)
		data = append(data, "\n"...)

		if _, err := f.Write(data); err != nil {
			f.Close() // ignore error; Write error takes precedence
			log.Fatal(err)
		}

		sum++
	}

	if err := f.Close(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Done. Wrote %v documents to file.\n", sum)
}
