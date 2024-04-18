// HEALTH CHECKER APP
package main

import (
	"crypto/tls"
	"fmt"
	"gopkg.in/yaml.v3"
	"net/http"
	"strconv"
	"time"
	"tiny-healt-checker/alerts"
	"tiny-healt-checker/structs"
	"tiny-healt-checker/utils"
)

// main function for health checker app entry point
func main() {
	// set yaml file path
	file := "config.yaml"
	// set config struct
	config := structs.Config{}
	// read yaml file
	data, _ := utils.ReadFile(file)
	// unmarshal yaml file
	err := yaml.Unmarshal([]byte(data), &config)
	utils.CheckError(err)

	// CheckError targets
	for _, target := range config.Target {
		if target.Retry.Interval < 1 {
			target.Retry.Interval = 1
		}
		if target.Retry.Count < 1 {
			target.Retry.Count = 1
		}
		// CheckError target
		if !requestToTargetIsActive(target) {
			// send alert
			sendAlert(target, config.Alert)
		}
	}
}

// requestToTargetIsActive function for checking target is active
func requestToTargetIsActive(target structs.Target) (isActive bool) {
	// make http client
	client := MakeHttpClient(target)

	i := 1
	for i <= target.Retry.Count {

		resp, err := MakeHttpRequest(target, client)

		if !utils.CheckError(err) {
			// If error, close response body and return false
			err := resp.Body.Close()
			if err != nil {
				return false
			}
			return false
		}

		// Check response status code
		if resp == nil {
			return false
		}
		if resp.StatusCode == target.Status {
			// If status matches, close response body and return true
			err := resp.Body.Close()
			if err != nil {
				return false
			}
			return true
		}

		// Close response body if status doesn't match
		err = resp.Body.Close()
		if err != nil {
			return false
		}

		// If status doesn't match, wait and retry
		fmt.Println(target.Name + " is DOWN. Sleep " + strconv.Itoa(target.Retry.Interval) + " secound. step: " + strconv.Itoa(i))
		time.Sleep(time.Duration(target.Retry.Interval) * time.Second)
		i++
	}
	return false
}

func MakeHttpClient(target structs.Target) *http.Client {
	// make http client
	client := &http.Client{}

	// set transport with ssl check configuration
	client.Transport = &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: !target.SSLVerify},
	}

	// set timeout
	client.Timeout = time.Duration(target.Timeout) * time.Second

	return client
}

func MakeHttpRequest(target structs.Target, client *http.Client) (*http.Response, error) {

	req, err := http.NewRequest(target.Method, target.Url, nil)
	if err != nil {
		fmt.Println(err)
	}

	// set headers
	for _, header := range target.Headers {
		req.Header.Set(header.Name, header.Value)
	}

	// send request
	return client.Do(req)
}

// sendAlert function for sending alert
func sendAlert(target structs.Target, alertList []structs.Alert) {
	for _, alert := range alertList {
		var t alerts.AlertInterface
		switch alert.Type {
		case "telegram":
			t = alerts.TelegramAlert{Target: target, Alert: alert}
		}
		if t != nil {
			t.SendAlert()
		}
	}
}
