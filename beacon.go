package beacon

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetOutboundIP() string {
	// endpoints which return plain text ip
	// TODO check if timedout, set timeout
	endpoints := []string{
		"https://mintz5.com/ip",
		"https://ifconfig.me",
	}

	for i := 0; i < len(endpoints); i++ {
		req_handler, err := http.Get(endpoints[i])
		if err != nil {
			fmt.Printf("Failed to curl %s\n", endpoints[i])
			break
		}

		defer req_handler.Body.Close()

		// Handle non 200 response
		if req_handler.StatusCode != 200 {
			fmt.Printf(
				"Non 200 response from %s: %d",
				endpoints[i],
				req_handler.StatusCode)
			break
		}

		// Handle cant read data..
		data, err := ioutil.ReadAll(req_handler.Body)
		if err != nil {
			fmt.Printf("COuld not read data from %s\n", endpoints[i])
			break
		}

		// Happy path
		ip := string(data)
		return ip

	}

	return ""
}

//GetGeo returns the Geo city of the IP
func GetGeo(ip string) (string, error) {
	var result map[string]interface{}
	endpoint := "https://freegeoip.app/json"

	req_handler, err := http.Get(endpoint)
	if err != nil {
		fmt.Printf("failed to request to %s\n", endpoint)
		return "", err
	}

	defer req_handler.Body.Close()

	// Handle non 200 response code
	if req_handler.StatusCode != 200 {
		msg := fmt.Sprintf("non 200 response from %s\n", endpoint)
		return "", errors.New(msg)
	}

	// Handle cant read data from endpoint
	data, err := ioutil.ReadAll(req_handler.Body)
	if err != nil {
		return "", errors.New(fmt.Sprintf("Not able to read data from %s\n", endpoint))
	}

	// Happy path
	json.Unmarshal(data, &result)
	return result["city"].(string) + "," + result["region_name"].(string), nil
}

// SendBeacon sends curl request with Geo info (stuffing into UA for now)
func SendBeacon(cityState, endpoint string) error {

	// Check endpoint (sanity check)
	if endpoint == "" {
		return errors.New("hostname or ip required..")
	}

	// Create http client
	client := &http.Client{}
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return err
	}

	// setup UA
	ua := fmt.Sprintf("Beacon:%s", cityState)
	req.Header.Set("User-Agent", ua)

	// Make request
	resp, err := client.Do(req)
	if resp.StatusCode != 200 {
		return errors.New(fmt.Sprintf("Non 200 response from %s\n", endpoint))
	}

	defer resp.Body.Close()

	return nil
}
