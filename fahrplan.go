package fahrplan

import (
	jsoniter "github.com/json-iterator/go"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func unmarshal(body []byte) (*Fahrplan, error) {
	schedule := new(Fahrplan)
	jsonErr := json.Unmarshal(body, schedule)
	if jsonErr != nil {
		return nil, jsonErr
	}
	return schedule, nil
}

func GetScheduleFromWeb(url string) (*Fahrplan, error) {
	resp, err := http.Get(url) //nolint:gosec
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}
	if resp.StatusCode >= 300 || resp.StatusCode < 200 {
		return nil, fmt.Errorf("Got not OK: %s", resp.Status)
	}

	body, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		return nil, readErr
	}
	schedule, err := unmarshal(body)
	if err != nil {
		return nil, err
	}
	return schedule, nil
}

func GetScheduleFromFile(path string) (*Fahrplan, error){
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	body, readErr := ioutil.ReadAll(file)
	if readErr != nil {
		return nil, readErr
	}
	schedule, err := unmarshal(body)
	if err != nil {
		return nil, err
	}
	return schedule, nil
}


