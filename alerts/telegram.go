package alerts

import (
	"net/http"
	"tiny-healt-checker/structs"
	"tiny-healt-checker/utils"
)

type TelegramAlert struct {
	Target structs.Target
	Alert  structs.Alert
}

// SendAlert function for sending telegram alert
func (t TelegramAlert) SendAlert() {
	// make client
	var client = &http.Client{}

	// make request
	req, err := http.NewRequest("GET", "https://api.telegram.org/bot"+t.Alert.Token+"/sendMessage", nil)
	utils.CheckError(err)

	// set query params
	q := req.URL.Query()
	q.Add("chat_id", t.Alert.ChatId)
	q.Add("text", "Target "+t.Target.Name+" is down")
	req.URL.RawQuery = q.Encode()

	// send request
	resp, err := client.Do(req)
	utils.CheckError(err)

	// close response body
	defer func(resp *http.Response) {
		err := resp.Body.Close()
		utils.CheckError(err)
	}(resp)

	// check response status code
	if resp.StatusCode != 200 {
		panic("Error sending telegram alert")
	}
}
