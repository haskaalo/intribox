package payload

// ReceiveAuthType Payload type of ReceiveAuth
const ReceiveAuthType = "AUTH"

// ReceiveAuth payload for init message (authentication)
type ReceiveAuth struct {
	Token         string `json:"token"`
	EnvironmentID int    `json:"environmentid"`
}

// SendAuthSuccessValue Payload type of SendAuth
const SendAuthSuccessValue = "AUTH_SUCCESS"

// SendAuthSuccess Authenticated succeed
type SendAuthSuccess struct {
	Ok int `json:"ok"`
}
