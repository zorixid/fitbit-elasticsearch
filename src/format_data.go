// Takes sample FitBit data and iterates over it to store it in a more manageable fashion:
// [
//   {
//     APIendpoint: activities-tracker-steps,
//     data: [],
//     DailyStats: [{ DateTime: 2018-02-01, Value: 14088 }, {...}, {...}]
//   },{
//     APIendpoint: activities-tracker-distance,
//     data: [],
//     DailyStats: [{ DateTime: 2018-02-01, Value: 5.92010781806808}, {...}, {...}]
//   },{
//     APIendpoint: activities-tracker-floors,
//     data: [],
//     DailyStats: [{ DateTime: 2018-02-01, Value: 14, {...}, {...}]
//   }
// ]

package main

import (
	"encoding/json"
	"fmt"
	"time"
)

// How the data comes in from the FitBit API
type FitBitData struct {
	ActivitiesTrackerSteps    []DailyStats `json:"activities-tracker-steps"`
	ActivitiesTrackerDistance []DailyStats `json:"activities-tracker-distance"`
	ActivitiesTrackerFloors   []DailyStats `json:"activities-tracker-floors"`
}

// temp data middleman
type callAPI struct {
	APIendpoint string
	data        []byte
	DailyStats  []DailyStats `json:"DailyStats"`
}

type DailyStats struct {
	DateTime string `json:"dateTime"`
	Value    string `json:"value"`
}

// create a new variable to store all the data in its final form
var finalData = []callAPI{}

func FormatData() {
	// sample data returned from the fitbit api
	// we need to store this data so we can send it to ES
	f_steps := []byte(`{"activities-tracker-steps":[{"dateTime":"2018-02-01","value":"14088"},{"dateTime":"2018-02-02","value":"22421"},{"dateTime":"2018-02-03","value":"12382"},{"dateTime":"2018-02-04","value":"3236"},{"dateTime":"2018-02-05","value":"2792"},{"dateTime":"2018-02-06","value":"16273"},{"dateTime":"2018-02-07","value":"3713"},{"dateTime":"2018-02-08","value":"22607"},{"dateTime":"2018-02-09","value":"13145"},{"dateTime":"2018-02-10","value":"8946"},{"dateTime":"2018-02-11","value":"10755"},{"dateTime":"2018-02-12","value":"12937"},{"dateTime":"2018-02-13","value":"7315"},{"dateTime":"2018-02-14","value":"6400"},{"dateTime":"2018-02-15","value":"4807"},{"dateTime":"2018-02-16","value":"8077"},{"dateTime":"2018-02-17","value":"8134"},{"dateTime":"2018-02-18","value":"1757"},{"dateTime":"2018-02-19","value":"10362"},{"dateTime":"2018-02-20","value":"2549"},{"dateTime":"2018-02-21","value":"4288"},{"dateTime":"2018-02-22","value":"7702"},{"dateTime":"2018-02-23","value":"6971"},{"dateTime":"2018-02-24","value":"11369"},{"dateTime":"2018-02-25","value":"8803"},{"dateTime":"2018-02-26","value":"6573"},{"dateTime":"2018-02-27","value":"6426"},{"dateTime":"2018-02-28","value":"7136"},{"dateTime":"2018-03-01","value":"3104"},{"dateTime":"2018-03-02","value":"8309"}]}`)
	f_dist := []byte(`{"activities-tracker-distance":[{"dateTime":"2018-02-01","value":"5.92010781806808"},{"dateTime":"2018-02-02","value":"9.656263666477999"},{"dateTime":"2018-02-03","value":"5.2166286367572"},{"dateTime":"2018-02-04","value":"1.08306241507984"},{"dateTime":"2018-02-05","value":"0.96366594053704"},{"dateTime":"2018-02-06","value":"6.98535552107328"},{"dateTime":"2018-02-07","value":"1.4127495421312"},{"dateTime":"2018-02-08","value":"9.48789692829368"},{"dateTime":"2018-02-09","value":"5.11517114852744"},{"dateTime":"2018-02-10","value":"3.4432041862296"},{"dateTime":"2018-02-11","value":"4.350530400788"},{"dateTime":"2018-02-12","value":"5.38191958754112"},{"dateTime":"2018-02-13","value":"2.71539210904"},{"dateTime":"2018-02-14","value":"2.38638849029984"},{"dateTime":"2018-02-15","value":"1.75669713803896"},{"dateTime":"2018-02-16","value":"3.0285010526888"},{"dateTime":"2018-02-17","value":"3.1829428624604"},{"dateTime":"2018-02-18","value":"0.59610623933328"},{"dateTime":"2018-02-19","value":"4.31043331776824"},{"dateTime":"2018-02-20","value":"0.75694195867056"},{"dateTime":"2018-02-21","value":"1.6027648526448"},{"dateTime":"2018-02-22","value":"2.6508316421912"},{"dateTime":"2018-02-23","value":"2.41205733424136"},{"dateTime":"2018-02-24","value":"4.5118072936716"},{"dateTime":"2018-02-25","value":"3.46705862629048"},{"dateTime":"2018-02-26","value":"2.71189378922904"},{"dateTime":"2018-02-27","value":"2.41205733424136"},{"dateTime":"2018-02-28","value":"2.70283419724968"},{"dateTime":"2018-03-01","value":"1.0171287178967199"},{"dateTime":"2018-03-02","value":"3.24659612736888"}]}`)
	f_floors := []byte(`{"activities-tracker-floors":[{"dateTime":"2018-02-01","value":"14"},{"dateTime":"2018-02-02","value":"64"},{"dateTime":"2018-02-03","value":"17"},{"dateTime":"2018-02-04","value":"9"},{"dateTime":"2018-02-05","value":"0"},{"dateTime":"2018-02-06","value":"31"},{"dateTime":"2018-02-07","value":"4"},{"dateTime":"2018-02-08","value":"37"},{"dateTime":"2018-02-09","value":"23"},{"dateTime":"2018-02-10","value":"7"},{"dateTime":"2018-02-11","value":"5"},{"dateTime":"2018-02-12","value":"35"},{"dateTime":"2018-02-13","value":"4"},{"dateTime":"2018-02-14","value":"5"},{"dateTime":"2018-02-15","value":"11"},{"dateTime":"2018-02-16","value":"11"},{"dateTime":"2018-02-17","value":"14"},{"dateTime":"2018-02-18","value":"5"},{"dateTime":"2018-02-19","value":"20"},{"dateTime":"2018-02-20","value":"0"},{"dateTime":"2018-02-21","value":"6"},{"dateTime":"2018-02-22","value":"10"},{"dateTime":"2018-02-23","value":"14"},{"dateTime":"2018-02-24","value":"25"},{"dateTime":"2018-02-25","value":"34"},{"dateTime":"2018-02-26","value":"13"},{"dateTime":"2018-02-27","value":"6"},{"dateTime":"2018-02-28","value":"14"},{"dateTime":"2018-03-01","value":"5"},{"dateTime":"2018-03-02","value":"38"}]}`)

	// first, we combine all data into a single slice
	data := []callAPI{
		callAPI{
			APIendpoint: "activities-tracker-steps",
			data:        f_steps,
		},
		callAPI{
			APIendpoint: "activities-tracker-distance",
			data:        f_dist,
		},
		callAPI{
			APIendpoint: "activities-tracker-floors",
			data:        f_floors,
		},
	}

	// for each sample data (3 total), do the following
	for _, single_dataset := range data {

		// unmarshal single_dataset.data into the APIdata variable
		var APIdata FitBitData
		err := json.Unmarshal(single_dataset.data, &APIdata)
		if err != nil {
			fmt.Printf("Error: %s\n", err)
		}

		// We know single_dataset.APIendpoint must match one of the following
		// Use this switch to match the data with the api call, and send the relevant data
		// to the iterateOverData() func
		switch single_dataset.APIendpoint {
		case "activities-tracker-steps":
			fmt.Println("one")
			go iterateOverData(APIdata.ActivitiesTrackerSteps, "activities-tracker-steps")

		case "activities-tracker-distance":
			fmt.Println("two")
			go iterateOverData(APIdata.ActivitiesTrackerDistance, "activities-tracker-distance")

		case "activities-tracker-floors":
			fmt.Println("three")
			go iterateOverData(APIdata.ActivitiesTrackerFloors, "activities-tracker-floors")

		}
	}
	time.Sleep(1 * time.Second)
	fmt.Printf("%+v\n", finalData)
}

// save FitBit data in a format that works for Elasticsearch
func iterateOverData(dataset []DailyStats, dataName string) {

	// Create a temporary to save the current data in
	returnedData := callAPI{}
	returnedData.APIendpoint = dataName

	// For each day of data, append the data to the returnedData.DailyStats
	for _, each_day := range dataset {
		temp := DailyStats{}
		temp.DateTime = each_day.DateTime
		temp.Value = each_day.Value

		returnedData.DailyStats = append(returnedData.DailyStats, temp)
	}

	// Once we've sorted all the data, append it to the finalData array
	finalData = append(finalData, returnedData)
}
