package alerts

type AlertInterface interface {
	SendAlert(message string)
}
