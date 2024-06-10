package models

type InboundMessage struct {
	ReceiverId          string `json:"receiver_id"`
	Message             string `json:"message"`
	TerminateConnection bool   `json:"terminate_connection"`
}

type OutboundMessage struct {
	SenderId string `json:"sender_id"`
	Message  string `json:"message"`
}
