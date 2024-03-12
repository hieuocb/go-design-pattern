package model

type Notifier interface {
	Send(message string)
}
