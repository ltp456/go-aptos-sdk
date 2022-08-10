package types

type Event struct {
	Key            string    `json:"key"`
	SequenceNumber string    `json:"sequence_number"`
	Type           string    `json:"type"`
	Data           EventData `json:"data"`
}

type EventData struct {
	Created string `json:"created"`
	RoleID  string `json:"role_id"`
}
