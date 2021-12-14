package cmd

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/url"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/hashicorp/go-retryablehttp"
)

/*
https://geocoding.geo.census.gov/geocoder/geographies/address?street=458%20North%20Oakhurst%20Drive%2C%20Apt.%20201&city=%20Beverly+Hills&state=CA&zip=90210&benchmark=Public_AR_Census2020&vintage=Census2020_Census2020&layers=10&format=json
*/

// MAX is maximum number of requests at a time
const MAX = 25

var start time.Time

func init() {
	start = time.Now()
}

func getResponse(url string) Response {
	retryClient := retryablehttp.NewClient()
	minDuration, _ := time.ParseDuration("5s")
	maxDuration, _ := time.ParseDuration("10s")
	maxRetry := 5

	retryClient.RetryWaitMin = minDuration
	retryClient.RetryWaitMax = maxDuration
	retryClient.RetryMax = maxRetry
	retryClient.Logger = nil

	standardClient := retryClient.StandardClient()

	res, err := standardClient.Get(url)

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)
	var response Response
	err = decoder.Decode(&response)

	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(response)
	return response
}

func genUrl(record []string) string {
	baseUrl := "https://geocoding.geo.census.gov/geocoder/geographies/address"

	params := url.Values{}
	params.Add("street", record[1])
	params.Add("city", record[2])
	params.Add("state", record[3])
	params.Add("zip", record[4])

	fullUrl := fmt.Sprintf("%s?%s&%s", baseUrl, params.Encode(), "benchmark=Public_AR_Census2020&vintage=Census2020_Census2020&layers=10&format=json")
	// fmt.Println(fullUrl)
	return fullUrl
}

func writeNewRecord(wg *sync.WaitGroup, c chan int, record []string, w *csv.Writer) {
	defer wg.Done()
	url := genUrl(record)
	response := getResponse(url)

	var (
		state  string
		county string
		tract  string
	)

	// When there is no address matched
	if len(response.Result.AddressMatches) == 0 {
		record = append(record, "")
		record = append(record, "")
		record = append(record, "")
		record = append(record, "")
	} else {
		geoResponse := response.Result.AddressMatches[0].Geographies

		if len(geoResponse.CensusBlockGroups) > 0 {
			state = geoResponse.CensusBlockGroups[0].State
			county = geoResponse.CensusBlockGroups[0].County
			tract = geoResponse.CensusBlockGroups[0].Tract
		} else {
			state = geoResponse.CensusBlocks[0].State
			county = geoResponse.CensusBlocks[0].County
			tract = geoResponse.CensusBlocks[0].Tract
		}

		record = append(record, state)
		record = append(record, county)
		record = append(record, tract)
		record = append(record, state+county+tract)
	}

	if err := w.Write(record); err != nil {
		log.Println(err)
	}
	<-c
}

func ExecuteTask() {
	// inputFileName := "test_two_records.csv"
	resultFileName := strings.Split(inputFileName, ".")[0] + "_result.csv"

	// Setup reader
	csvIn, err := os.Open(inputFileName)

	if err != nil {
		log.Fatal("Unable to read input file!")
	}

	// setup writer
	csvOut, err := os.Create(resultFileName)

	if err != nil {
		log.Fatal("Unable to open output!")
	}

	w := csv.NewWriter(csvOut)
	defer csvOut.Close()

	r := csv.NewReader(csvIn)

	var wg sync.WaitGroup
	c := make(chan int, MAX)

	fmt.Println("Working on it!")

	record_num := 0
	break_number := 2000

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		record_num += 1

		if record_num == break_number {
			fmt.Printf("Processed %d records\n", break_number)
			break_number += 2000
		}

		wg.Add(1)
		c <- 1
		go writeNewRecord(&wg, c, record, w)
	}
	// Wait for all goroutines to finish
	wg.Wait()
	// Close channel
	close(c)
	w.Flush()
	fmt.Printf("Processed %d records\n", record_num)
	fmt.Printf("Program finished after %f minutes\n", time.Since(start).Minutes())
}
