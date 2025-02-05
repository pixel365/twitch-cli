package mock_server

import (
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

type Client struct {
	keepAliveTimer       *time.Ticker
	keepAliveLoopChan    chan struct{}
	pingTimer            *time.Ticker
	pingLoopChan         chan struct{}
	mustSubscribeTimer   *time.Timer
	conn                 *websocket.Conn
	connectionUrl        string
	clientName           string
	ConnectedAtTimestamp string
	keepAliveSeconds     int
	mutex                sync.Mutex
	KeepAliveEnabled     bool
	keepAliveChanOpen    bool
	pingChanOpen         bool
}

func (c *Client) SendMessage(messageType int, data []byte) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	return c.conn.WriteMessage(messageType, data)
}

func (c *Client) CloseWithReason(reason *CloseMessage) {
	c.conn.WriteControl(
		websocket.CloseMessage,
		websocket.FormatCloseMessage(reason.code, reason.message),
		time.Now().Add(2*time.Second),
	)
}

func (c *Client) CloseDirty() {
	c.conn.Close()
}
