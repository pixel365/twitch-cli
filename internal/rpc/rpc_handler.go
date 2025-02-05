package rpc_handler

import (
	"net"
	"net/http"
	"net/rpc"
	"strconv"
)

var globalHandler RPCHandler

type HandlerCallback func(RPCArgs) RPCResponse

type RPCArgs struct {
	Variables map[string]string
	RPCName   string
	Body      string
}

type RPCResponse struct {
	DetailedInfo string
	ResponseCode int
}

type RPCHandler struct {
	listener net.Listener
	Handlers map[string]HandlerCallback
	Port     int
}

func (rpch *RPCHandler) RegisterHandler(rpcName string, callback HandlerCallback) bool {
	_, exists := rpch.Handlers[rpcName]
	if exists {
		return false
	}

	rpch.Handlers[rpcName] = callback
	return true
}

func (rpch *RPCHandler) StartBackgroundServer() error {
	rpc.Register(rpch)
	rpc.HandleHTTP()
	l, err := net.Listen("tcp", ":"+strconv.Itoa(rpch.Port))
	if err != nil {
		return err
	}
	rpch.listener = l
	go http.Serve(rpch.listener, nil)

	return nil
}

func (rpch *RPCHandler) ShutdownServer() {
	rpch.listener.Close()
}

func (rpch *RPCHandler) ExecuteGenericRPC(args RPCArgs, reply *RPCResponse) error {
	handler, exists := rpch.Handlers[args.RPCName]
	if !exists {
		*reply = RPCResponse{
			ResponseCode: -1, // -1 is failed here; 0 is success; anything above zero is special failure code
			DetailedInfo: "Given RPCName is not currently registered",
		}
		return nil
	}

	*reply = handler(args)

	return nil
}

// Used for testing connection
func (rphc *RPCHandler) Verify(args RPCArgs, reply *RPCArgs) error {
	*reply = args
	return nil
}
