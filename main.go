// HEALTH CHECKER APP
package main

import (
	"crypto/tls"
	"fmt"
	"gopkg.in/yaml.v3"
	"io"
	"net/http"
	"strconv"
	"time"
	"tiny-healt-checker/alerts"
	"tiny-healt-checker/structs"
	"tiny-healt-checker/utils"
)

func main() {
	file := "config.yaml"
	config := structs.Config{}
	data, _ := utils.ReadFile(file)
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
		isActive, resp, bodyBytes := requestToTargetIsActive(target)
		if !isActive {
			// send alert
			sendAlert(target, config.Alert, resp, bodyBytes)
		}
	}
}

func requestToTargetIsActive(target structs.Target) (bool, *http.Response, []byte) {
	client := MakeHttpClient(target)
	var resp *http.Response
	var err error
	var bodyBytes []byte
	i := 1
	for i <= target.Retry.Count {
		resp, err = MakeHttpRequest(target, client)
		if !utils.CheckError(err) {
			// If error, close response body and return false
			err = resp.Body.Close()
			if err != nil {
				return false, resp, bodyBytes
			}
			return false, resp, bodyBytes
		}

		// Check response status code
		if resp == nil {
			return false, nil, nil
		}

		bodyBytes, err = io.ReadAll(resp.Body)

		// If status matches, close response body and return true
		err = resp.Body.Close()
		if err != nil {
			return false, resp, bodyBytes
		}

		if resp.StatusCode == target.Status {
			return true, resp, bodyBytes
		}

		// Close response body if status doesn't match
		err = resp.Body.Close()

		// If status doesn't match, wait and retry
		fmt.Println(target.Name + " is DOWN. Sleep " + strconv.Itoa(target.Retry.Interval) + " second. step: " + strconv.Itoa(i))
		time.Sleep(time.Duration(target.Retry.Interval) * time.Second)
		i++
	}
	return false, resp, bodyBytes
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

func sendAlert(target structs.Target, alertList []structs.Alert, resp *http.Response, bodyBytes []byte) {
	for _, alert := range alertList {
		var t alerts.AlertInterface
		switch alert.Type {
		case "telegram":
			t = alerts.TelegramAlert{Target: target, Alert: alert}
			break
		case "std":
			t = alerts.StdAlert{Target: target, Alert: alert}
			break
		}
		if t != nil {
			message := serializeAlertMessage(target, resp, bodyBytes)
			t.SendAlert(message)
		}
	}
}

func serializeAlertMessage(target structs.Target, resp *http.Response, bodyBytes []byte) string {
	message := "!!! DOWN: `" + target.Name + "` | " + target.Method + " `" + target.Url + "`"
	if resp != nil {
		message += " `" + resp.Status + "`"
	}
	message += " (Count: `" + strconv.Itoa(target.Retry.Count) + "`, Interval: `" + strconv.Itoa(target.Retry.Interval) + "`)"
	if bodyBytes != nil {
		message += "\n```\n" + string(bodyBytes) + "```"
	}
	return message
}
