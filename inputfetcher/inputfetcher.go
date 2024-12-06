package inputfetcher

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	baseUri       = "https://adventofcode.com/"
	dayFormatPath = "2024/day/%d/input"
)

type DataFetcher struct {
	bearerToken string
}

func getInputDataPath(day int) string {
	return fmt.Sprintf(dayFormatPath, day)
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

func (fetcher *DataFetcher) GetDataByDay(day int) (string, error) {
	return fetcher.fetchData(getInputDataPath(day))
}
