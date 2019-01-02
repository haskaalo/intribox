package websocket

import (
	ws "github.com/gorilla/websocket"
)

var (
	// CloseUnknownError Error as of internal error
	CloseUnknownError = ws.FormatCloseMessage(4000, "Unknown Error")

	// CloseDecodeError Error decoding payload
	CloseDecodeError = ws.FormatCloseMessage(4001, "Error decoding payload")

	// CloseUnknownType Type sent in payload doesn't exist
	CloseUnknownType = ws.FormatCloseMessage(4002, "Unknown Type")

	// CloseAuthenticationFailed Token sent in auth payload isnt valid
	CloseAuthenticationFailed = ws.FormatCloseMessage(4003, "Authentication failed")

	// CloseEnvironmentNotFound EnvironmentID sent in the auth payload doesn't exist
	CloseEnvironmentNotFound = ws.FormatCloseMessage(4004, "Environment not found")
)
