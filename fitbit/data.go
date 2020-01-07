// This file makes the API call to the FitBit API
package fitbit

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func GetActivity() {

	testing := os.Getenv("TESTING")
	var d, f, s FitBitData

	if (testing != "true") {
		d = GetDistance(DatePair{"2018-02-01", "2018-03-03"})
		f = GetFloors(DatePair{"2018-02-01", "2018-03-03"})
		s = GetSteps(DatePair{"2018-02-01", "2018-03-03"})
	} else {
		d = UnmarshalData([]byte(`{"activities-distance":[{"dateTime":"2018-02-01","value":"5.92010781806808"},{"dateTime":"2018-02-02","value":"9.656263666477999"},{"dateTime":"2018-02-03","value":"5.2166286367572"},{"dateTime":"2018-02-04","value":"1.08306241507984"},{"dateTime":"2018-02-05","value":"0.96366594053704"},{"dateTime":"2018-02-06","value":"6.98535552107328"},{"dateTime":"2018-02-07","value":"1.4127495421312"},{"dateTime":"2018-02-08","value":"9.48789692829368"},{"dateTime":"2018-02-09","value":"5.11517114852744"},{"dateTime":"2018-02-10","value":"3.4432041862296"},{"dateTime":"2018-02-11","value":"4.350530400788"},{"dateTime":"2018-02-12","value":"5.38191958754112"},{"dateTime":"2018-02-13","value":"2.71539210904"},{"dateTime":"2018-02-14","value":"2.38638849029984"},{"dateTime":"2018-02-15","value":"1.75669713803896"},{"dateTime":"2018-02-16","value":"3.0285010526888"},{"dateTime":"2018-02-17","value":"3.1829428624604"},{"dateTime":"2018-02-18","value":"0.59610623933328"},{"dateTime":"2018-02-19","value":"4.31043331776824"},{"dateTime":"2018-02-20","value":"0.75694195867056"},{"dateTime":"2018-02-21","value":"1.6027648526448"},{"dateTime":"2018-02-22","value":"2.6508316421912"},{"dateTime":"2018-02-23","value":"2.41205733424136"},{"dateTime":"2018-02-24","value":"4.5118072936716"},{"dateTime":"2018-02-25","value":"3.46705862629048"},{"dateTime":"2018-02-26","value":"2.71189378922904"},{"dateTime":"2018-02-27","value":"2.41205733424136"},{"dateTime":"2018-02-28","value":"2.70283419724968"},{"dateTime":"2018-03-01","value":"1.0171287178967199"},{"dateTime":"2018-03-02","value":"3.24659612736888"},{"dateTime":"2018-03-03","value":"3.38885284806536"}]}`))
		f = UnmarshalData([]byte(`{"activities-floors":[{"dateTime":"2018-02-01","value":"14"},{"dateTime":"2018-02-02","value":"64"},{"dateTime":"2018-02-03","value":"17"},{"dateTime":"2018-02-04","value":"9"},{"dateTime":"2018-02-05","value":"0"},{"dateTime":"2018-02-06","value":"31"},{"dateTime":"2018-02-07","value":"4"},{"dateTime":"2018-02-08","value":"37"},{"dateTime":"2018-02-09","value":"23"},{"dateTime":"2018-02-10","value":"7"},{"dateTime":"2018-02-11","value":"5"},{"dateTime":"2018-02-12","value":"35"},{"dateTime":"2018-02-13","value":"4"},{"dateTime":"2018-02-14","value":"5"},{"dateTime":"2018-02-15","value":"11"},{"dateTime":"2018-02-16","value":"11"},{"dateTime":"2018-02-17","value":"14"},{"dateTime":"2018-02-18","value":"5"},{"dateTime":"2018-02-19","value":"20"},{"dateTime":"2018-02-20","value":"0"},{"dateTime":"2018-02-21","value":"6"},{"dateTime":"2018-02-22","value":"10"},{"dateTime":"2018-02-23","value":"14"},{"dateTime":"2018-02-24","value":"25"},{"dateTime":"2018-02-25","value":"34"},{"dateTime":"2018-02-26","value":"13"},{"dateTime":"2018-02-27","value":"6"},{"dateTime":"2018-02-28","value":"14"},{"dateTime":"2018-03-01","value":"5"},{"dateTime":"2018-03-02","value":"38"},{"dateTime":"2018-03-03","value":"30"}]}`))
		s = UnmarshalData([]byte(`{"activities-steps":[{"dateTime":"2018-02-01","value":"14088"},{"dateTime":"2018-02-02","value":"22421"},{"dateTime":"2018-02-03","value":"12382"},{"dateTime":"2018-02-04","value":"3236"},{"dateTime":"2018-02-05","value":"2792"},{"dateTime":"2018-02-06","value":"16273"},{"dateTime":"2018-02-07","value":"3713"},{"dateTime":"2018-02-08","value":"22607"},{"dateTime":"2018-02-09","value":"13145"},{"dateTime":"2018-02-10","value":"8946"},{"dateTime":"2018-02-11","value":"10755"},{"dateTime":"2018-02-12","value":"12937"},{"dateTime":"2018-02-13","value":"7315"},{"dateTime":"2018-02-14","value":"6400"},{"dateTime":"2018-02-15","value":"4807"},{"dateTime":"2018-02-16","value":"8077"},{"dateTime":"2018-02-17","value":"8134"},{"dateTime":"2018-02-18","value":"1757"},{"dateTime":"2018-02-19","value":"10362"},{"dateTime":"2018-02-20","value":"2549"},{"dateTime":"2018-02-21","value":"4288"},{"dateTime":"2018-02-22","value":"7702"},{"dateTime":"2018-02-23","value":"6971"},{"dateTime":"2018-02-24","value":"11369"},{"dateTime":"2018-02-25","value":"8803"},{"dateTime":"2018-02-26","value":"6573"},{"dateTime":"2018-02-27","value":"6426"},{"dateTime":"2018-02-28","value":"7136"},{"dateTime":"2018-03-01","value":"3104"},{"dateTime":"2018-03-02","value":"8309"},{"dateTime":"2018-03-03","value":"8843"}]}`))
	}

	combinedData := CombineData(d.Distance, f.Floors, s.Steps)

	WriteData(combinedData)
}

func GetActivityOnDate(url string) []byte {
	// Transport Options. We need to set DisableCompression to false in order to be able to decode the response from the API
	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: false,
	}

	// Create an HTTP Client
	client := &http.Client{Transport: tr}

	// fmt.Printf("url: %s\n", url)

	// Create a new request with our access token
	access_token := os.Getenv("ACCESS_TOKEN")
	req, _ := http.NewRequest("GET", url, nil)

	// Add the access token to the header, and return miles instead of KMs
	req.Header.Add("Authorization", "Bearer "+access_token)
	req.Header.Add("Accept-Language", "en_US")

	// Make the call to the API
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error retrieving activities: %s\n", err)
	}
	defer resp.Body.Close()

	// Read and return the response body
	body, _ := ioutil.ReadAll(resp.Body)

	// Did the request successfully return fitbit data?
	// if it did, return the body. if it didn't, refresh the Auth token
	if isTokenActive(string(body)) {
		return body
	} else {
		refreshAuthToken()
		return nil
	}
}
