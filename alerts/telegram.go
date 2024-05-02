package alerts

import (
	"fmt"
	"io"
	"net/http"
	"tiny-healt-checker/structs"
	"tiny-healt-checker/utils"
)

type TelegramAlert struct {
	Target structs.Target
	Alert  structs.Alert
}

// SendAlert function for sending telegram alert
func (t TelegramAlert) SendAlert(message string) {
	// make client
	var client = &http.Client{}

	// make request
	req, err := http.NewRequest("GET", "https://api.telegram.org/bot"+t.Alert.Token+"/sendMessage", nil)
	utils.HasNotError(err)

	// cut text if it is too long. Telegram has a limit of 4096 characters
	if len(message) > 4096 {
		message = message[:4096]
	}

	// set query params
	q := req.URL.Query()
	q.Add("chat_id", t.Alert.ChatId)
	q.Add("text", message)
	q.Add("parse_mode", "markdown")
	req.URL.RawQuery = q.Encode()

	// send request
	resp, err := client.Do(req)
	utils.HasNotError(err)

	// close response body
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		utils.HasNotError(err)
	}(resp.Body)

	// check response status code
	if resp.StatusCode != 200 {
		fmt.Println("Error sending Telegram alert")
	}
}
