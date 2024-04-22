package alerts

import (
	"fmt"
	"strings"
	"tiny-healt-checker/structs"
)

type StdAlert struct {
	Target structs.Target
	Alert  structs.Alert
}

// SendAlert function for sending telegram alert
func (t StdAlert) SendAlert(message string) {
	fmt.Println(strings.ReplaceAll(message, "\n", "\t"))
}
