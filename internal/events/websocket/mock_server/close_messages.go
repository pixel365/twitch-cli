package mock_server

type CloseMessage struct {
	message string
	code    int
}

var (
	closeClientDisconnected = &CloseMessage{
		code:    1000,
		message: "client disconnected",
	}

	closeInternalServerError = &CloseMessage{
		code:    4000,
		message: "internal server error",
	}

	closeClientSentInboundTraffic = &CloseMessage{
		code:    4001,
		message: "sent inbound traffic",
	}

	closeClientFailedPingPong = &CloseMessage{
		code:    4002,
		message: "failed ping pong",
	}

	closeConnectionUnused = &CloseMessage{
		code:    4003,
		message: "connection unused",
	}

	closeReconnectGraceTimeExpired = &CloseMessage{
		code:    4004,
		message: "client reconnect grace time expired",
	}

	closeNetworkTimeout = &CloseMessage{
		code:    4005,
		message: "network timeout",
	}

	closeNetworkError = &CloseMessage{
		code:    4006,
		message: "network error",
	}

	closeInvalidReconnect = &CloseMessage{
		code:    4007,
		message: "invalid reconnect attempt",
	}
)

func GetCloseMessageFromCode(code int) *CloseMessage {
	switch code {
	case 4000:
		return closeInternalServerError
	case 4001:
		return closeClientSentInboundTraffic
	case 4002:
		return closeClientFailedPingPong
	case 4003:
		return closeConnectionUnused
	case 4004:
		return closeNetworkTimeout
	case 4005:
		return closeNetworkTimeout
	case 4006:
		return closeNetworkError
	case 4007:
		return closeInvalidReconnect
	default:
		return nil
	}
}
