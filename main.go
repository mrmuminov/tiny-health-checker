// HEALTH CHECKER APP
package main

import (
	"bufio"
	"fmt"
	"gopkg.in/yaml.v3"
	"net/http"
	"os"
	"time"
)

// Target struct
type Target struct {
	Name    string   `yaml:"name"`
	Url     string   `yaml:"url"`
	Method  string   `yaml:"method"`
	Headers []Header `yaml:"headers"`
	Body    string   `yaml:"body"`
	Timeout int      `yaml:"timeout"`
}

// Header struct
type Header struct {
	Name  string `yaml:"name"`
	Value string `yaml:"value"`
}

// Alert struct
type Alert struct {
	Name   string `yaml:"name"`
	Type   string `yaml:"type"`
	Token  string `yaml:"token"`
	ChatId string `yaml:"chat_id"`
}

// Config struct
type Config struct {
	Target []Target `yaml:"target"`
	Alert  []Alert  `yaml:"alert"`
}

// check error
func check(e error) {
	if e != nil {
		panic(e)
	}
}

// main function
func main() {
	// set yaml file path
	file := "config.yaml"
	// set config struct
	config := Config{}
	// read yaml file
	data, _ := readFile(file)
	// unmarshal yaml file
	err := yaml.Unmarshal([]byte(data), &config)
	check(err)

	// check targets
	for _, target := range config.Target {
		// check target
		if !requestToTargetIsActive(target) {
			// send alert
			sendAlert(target, config.Alert)
		}
	}
}

func requestToTargetIsActive(target Target) bool {
	// make http client
	var client = &http.Client{}
	// set timeout
	client.Timeout = time.Duration(target.Timeout) * time.Second
	// make http request
	req, err := http.NewRequest(target.Method, target.Url, nil)
	if err != nil {
		fmt.Println(err)
	}
	// set headers
	for _, header := range target.Headers {
		req.Header.Set(header.Name, header.Value)
	}
	// send request
	resp, err := client.Do(req)
	if err == nil {
		// close response body
		defer func(resp *http.Response) {
			err := resp.Body.Close()
			check(err)
		}(resp)
		// check response status code
		if resp.StatusCode != 200 {
			return false
		}
	} else {
		return false
	}
	return true
}

func sendAlert(target Target, alert []Alert) {
	// send alert
	for _, alert := range alert {
		switch alert.Type {
		case "telegram":
			// send telegram alert
			sendTelegramAlert(target, alert)
		}
	}
}

func sendTelegramAlert(target Target, alert Alert) {
	// make client
	var client = &http.Client{}
	// make request
	req, err := http.NewRequest("GET", "https://api.telegram.org/bot"+alert.Token+"/sendMessage", nil)
	check(err)
	// set query params
	q := req.URL.Query()
	q.Add("chat_id", alert.ChatId)
	q.Add("text", "Target "+target.Name+" is down")
	req.URL.RawQuery = q.Encode()
	// send request
	resp, err := client.Do(req)
	check(err)
	// close response body
	defer func(resp *http.Response) {
		err := resp.Body.Close()
		check(err)
	}(resp)
	// check response status code
	if resp.StatusCode != 200 {
		panic("Error sending telegram alert")
	}
}

// read file
func readFile(filePath string) (string, error) {
	// open file
	file, err := os.Open(filePath)
	// check error
	check(err)
	// close file
	defer func(file *os.File) {
		err := file.Close()
		check(err)
	}(file)
	// read file
	var lines string
	scanner := bufio.NewScanner(file)
	// read line by line
	for scanner.Scan() {
		lines += scanner.Text() + "\n"
	}
	return lines, scanner.Err()
}
