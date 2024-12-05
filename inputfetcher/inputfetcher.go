package inputfetcher

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	baseUri  = "https://adventofcode.com/"
	day1Path = "2024/day/1/input"
	day2Path = "2024/day/2/input"
	day3Path = "2024/day/3/input"
	day4Path = "2024/day/4/input"
	day5Path = "2024/day/5/input"
	day6Path = "2024/day/6/input"
	day7Path = "2024/day/7/input"
)

type DataFetcher struct {
	bearerToken string
}

func NewDataFetcher(bearerToken string) DataFetcher {
	df := DataFetcher{}
	df.bearerToken = bearerToken
	return df
}

func (fetcher *DataFetcher) fetchData(path string) (string, error) {
	log.Printf("Entering FetchData, Path:%s\n", path)

	// Construct the full URL
	url := baseUri + path

	// Create a new GET request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}

	// Add the "session" cookie with the bearerToken
	req.AddCookie(&http.Cookie{
		Name:  "session",
		Value: fetcher.bearerToken,
	})

	// Use the default HTTP client to send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to execute request: %w", err)
	}
	defer resp.Body.Close()

	// Check for non-200 HTTP status
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("received non-200 status: %s", resp.Status)
	}

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %w", err)
	}

	return string(body), nil

}

func (fetcher *DataFetcher) FetchDay1Data() ([]int, []int, error) {

	rawData, err := fetcher.fetchData(day1Path)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to fetch data from AOC: %w", err)
	}

	left, right := day1DataParser(rawData)

	return left, right, nil
}

func (fetcher *DataFetcher) FetchDay4Data() (string, error) {
	return fetcher.fetchData(day4Path)
}

func (fetcher *DataFetcher) FetchDay5Data() (string, error) {
	return fetcher.fetchData(day5Path)
}

func (fetcher *DataFetcher) FetchDay6Data() (string, error) {
	return fetcher.fetchData(day6Path)
}

func (fetcher *DataFetcher) FetchDay7Data() (string, error) {
	return fetcher.fetchData(day7Path)
}
