package knx

import (
	"context"
	"errors"
	"time"
)

// requestConnection sends a connection request every 500ms through the socket until the provided
// context gets canceled, or a response is received. A response that renders the gateway as busy
// will not stop requestConnection.
func requestConnection(ctx context.Context, sock Socket) (channel uint8, err error) {
	req := &ConnectionRequest{}

	// Send the initial request.
	err = sock.Send(req)
	if err != nil {
		return
	}

	// Create a resend timer.
	ticker := time.NewTicker(500 * time.Millisecond)
	defer ticker.Stop()

	// Cycle until a request gets a response.
	for {
		select {
		// Termination has been requested.
		case <-ctx.Done():
			return 0, ctx.Err()

		// Resend timer triggered.
		case <-ticker.C:
			err = sock.Send(req)
			if err != nil {
				return
			}

		// A message has been received or the channel has been closed.
		case msg, open := <-sock.Inbound():
			if !open {
				return 0, errors.New("Inbound channel has been closed")
			}

			// We're only interested in connection responses.
			if res, ok := msg.(*ConnectionResponse); ok {
				switch res.Status {
				// Conection has been established.
				case ConnResOk:
					return res.Channel, nil

				// The gateway is busy, but we don't stop yet.
				case ConnResBusy:
					continue

				// Connection request has been denied.
				default:
					return 0, res.Status
				}
			}
		}
	}
}

type connHandle struct {
	ctx     context.Context
	sock    Socket
	channel uint8
}

// requestConnectionState periodically sends a connection state request to the gateway until it has
// received a response, the context is done, or 10s have passed.
func (conn connHandle) requestConnectionState(
	heartbeat <-chan ConnState,
) error {
	req := &ConnectionStateRequest{conn.channel, 0, HostInfo{}}

	// Send first connection state request
	err := conn.sock.Send(req)
	if err != nil {
		return err
	}

	// Start the resend timer.
	ticker := time.NewTicker(500 * time.Millisecond)
	defer ticker.Stop()

	// Start the timeout timer.
	timeout := time.After(10 * time.Second)

	for {
		select {
		// Termination has been requested.
		case <-conn.ctx.Done():
			return conn.ctx.Err()

		// Request timed out.
		case <-timeout:
			return errors.New("Connection state request has been ignored")

		// Resend timer fired.
		case <-ticker.C:
			err := conn.sock.Send(req)
			if err != nil {
				return err
			}

		// Received a connection state response.
		case res := <-heartbeat:
			// Is connection state positive?
			if res == 0 {
				return nil
			} else {
				return res
			}
		}
	}
}

// pushHeartbeatError transmits a heartbeat error if the internal Context is not done yet.
func (conn connHandle) pushHeartbeatError(heartbeatErr chan<- error, err error) {
	select {
	case <-conn.ctx.Done():
	case heartbeatErr <- err:
	}
}

// serveHeartbeat processes heartbeat requests.
func (conn connHandle) serveHeartbeat(
	trigger      <-chan struct{},
	heartbeat    <-chan ConnState,
	heartbeatErr chan<- error,
) {
	for {
		select {
		// Termination has been requested.
		case <-conn.ctx.Done():
			return

		// Heartbeat check has triggered.
		case <-trigger:
			err := conn.requestConnectionState(heartbeat)
			if err != nil {
				go conn.pushHeartbeatError(heartbeatErr, err)
				return
			}
		}
	}
}

// pushData transmits the given data to the inbound channel if the internal Context is not done yet.
func (conn connHandle) pushData(data []byte, inbound chan<- []byte) {
	select {
	case <-conn.ctx.Done():
	case inbound <- data:
	}
}

// handleTunnelRequest validates the request, pushes the data to the client and acknowledges the
// request for the gateway.
func (conn connHandle) handleTunnelRequest(
	req       *TunnelRequest,
	seqNumber *uint8,
	inbound   chan<- []byte,
) error {
	// Validate the request channel.
	if req.Channel != conn.channel {
		return errors.New("Invalid communication channel in tunnel request")
	}

	// Is the sequence number what we expected?
	if req.SeqNumber == *seqNumber {
		*seqNumber++
		go conn.pushData(req.Payload, inbound)
	}

	// Send the acknowledgement.
	return conn.sock.Send(&TunnelResponse{conn.channel, req.SeqNumber, 0})
}

//
func (conn connHandle) triggerHeartbeat(trigger chan<- struct{}) {
	select {
	case <-conn.ctx.Done():
	case trigger <- struct{}{}:
	}
}

// serveInbound processes incoming packets.
func (conn connHandle) serveInbound(
	heartbeatErr <-chan error,
	inbound      chan<- []byte,
	trigger      chan<- struct{},
	heartbeat    chan<- ConnState,
) {
	var seqNumber uint8 = 0

	for {
		select {
		// Termination has been requested.
		case <-conn.ctx.Done():
			return

		// Heartbeat worker has signaled an error.
		case <-heartbeatErr:
			return

		// There were no incoming packets for 10s.
		case <-time.After(10 * time.Second):
			go conn.triggerHeartbeat(trigger)

		// A message has been received or the channel is closed.
		case msg, open := <-conn.sock.Inbound():
			if !open {
				return
			}

			// Determine what to do with the message.
			switch msg.(type) {
			case *TunnelRequest:
				req := msg.(*TunnelRequest)
				err := conn.handleTunnelRequest(req, &seqNumber, inbound)
				if err != nil {
					Logger.Printf("Error while handling tunnel request %v: %v", req, err)
				}
			}
		}
	}
}

// //
// type Client struct {
// 	sock    Socket
// 	channel byte
// 	reaper  context.CancelFunc
// 	seq     <-chan byte
// 	ack     <-chan byte

// 	//
// 	Inbound <-chan []byte
// }

// //
// var (
// 	ErrConnTimeout = errors.New("Gateway did not response to connection request")
// 	ErrConnRejected = errors.New("Gateway rejected connection request")
// )

// //
// func NewClient(gatewayAddress string) (*Client, error) {
// 	sock, err := NewClientSocket(gatewayAddress)
// 	if err != nil {
// 		return nil, err
// 	}

// 	req := &ConnectionRequest{}

// 	err = sock.Send(req)
// 	if err != nil {
// 		return nil, err
// 	}

// 	resChan := awaitConnectionResponse(sock)

// 	// Connection cycle
// 	for i := 0; i < 20; i++ {
// 		select {
// 		case res := <-resChan:
// 			if res.Status == 0 {
// 				// Connection is established.

// 				// Logger.Printf("Client[%v]: Connection has been established on channel %v",
// 				//               sock.conn.RemoteAddr(), res.Channel)

// 				return makeClient(sock, res.Channel), nil
// 			} else {
// 				// Connection attempt was rejected.

// 				// Logger.Printf("Client[%v]: Connection attempt was rejected",
// 				//               sock.conn.RemoteAddr())

// 				sock.Close()
// 				return nil, ErrConnRejected
// 			}

// 		case <-time.After(500 * time.Millisecond):
// 			// Resend the connection request, if we haven't received a response from the gateway
// 			// after 1 second.
// 			err := sock.Send(req)
// 			if err != nil {
// 				return nil, err
// 			}
// 		}
// 	}

// 	// Logger.Printf("Client[%v]: Connection attempts timed out", sock.conn.RemoteAddr())

// 	sock.Close()
// 	return nil, ErrConnTimeout
// }

// //
// func (client *Client) Close() {
// 	client.reaper()
// 	client.sock.Close()
// }

// //
// var (
// 	ErrSendClosed   = errors.New("Outbound worker has terminated")
// 	ErrSendRejected = errors.New("Gateway rejected tunnel request")
// 	ErrSendTimeout  = errors.New("Gateway did not acknowledge tunnel request in time")
// )

// //
// func (client *Client) Send(data []byte) error {
// 	seqNumber, open := <-client.seq
// 	if !open {
// 		return ErrSendClosed
// 	}

// 	req := &TunnelRequest{client.channel, seqNumber, data}
// 	err := client.sock.Send(req)
// 	if err != nil {
// 		return err
// 	}

// 	for i := 0; i < 20; i++ {
// 		select {
// 		case status, open := <-client.ack:
// 			if !open {
// 				return ErrSendClosed
// 			}

// 			if status == 0 {
// 				return nil
// 			} else {
// 				return ErrSendRejected
// 			}

// 		case <-time.After(500 * time.Millisecond):
// 			err := client.sock.Send(req)
// 			if err != nil {
// 				return err
// 			}
// 		}
// 	}

// 	return ErrSendTimeout
// }

// //
// func awaitConnectionResponse(sock Socket) <-chan *ConnectionResponse {
// 	resChan := make(chan *ConnectionResponse)

// 	go func() {
// 		for payload := range sock.Inbound() {
// 			res, ok := payload.(*ConnectionResponse)
// 			if ok {
// 				resChan <- res
// 				return
// 			}
// 		}
// 	}()

// 	return resChan
// }

// //
// func makeClient(sock Socket, channel byte) *Client {
// 	ctx, reaper := context.WithCancel(context.Background())

// 	seq := make(chan byte)
// 	ack := make(chan byte)
// 	tunRes := make(chan *TunnelResponse)
// 	go clientOutboundWorker(ctx, reaper, tunRes, seq, ack)

// 	inbound := make(chan []byte)
// 	go clientInboundWorker(ctx, reaper, sock, channel, tunRes, inbound)

// 	return &Client{sock, channel, reaper, seq, ack, inbound}
// }

// //
// func clientInboundWorker(
// 	ctx     context.Context,
// 	reaper  context.CancelFunc,
// 	sock    Socket,
// 	channel byte,
// 	tunRes  chan<- *TunnelResponse,
// 	inbound chan<- []byte,
// ) {
// 	// Logger.Printf("Client[%v]: Started inbound worker", sock.conn.RemoteAddr())
// 	// defer Logger.Printf("Client[%v]: Stopped inbound worker", sock.conn.RemoteAddr())

// 	defer close(inbound)
// 	defer reaper()

// 	heartbeatTrigger := make(chan struct{})
// 	stateResChan := make(chan *ConnectionStateResponse)
// 	go clientHeartbeatWorker(ctx, reaper, sock, channel, heartbeatTrigger, stateResChan)

// 	var incomingSeqNumber byte = 0

// 	for {
// 		select {
// 		// Goroutine exit has been requested
// 		case <-ctx.Done():
// 			return

// 		// 10 seconds without communication, time for a heartbeat
// 		case <-time.After(5 * time.Second):
// 			// Logger.Printf("Client[%v]: Triggering heartbeat", sock.conn.RemoteAddr())

// 			select {
// 			case <-ctx.Done():
// 				return

// 			case heartbeatTrigger <- struct{}{}:
// 			}

// 		// Incoming packets
// 		case payload, open := <-sock.Inbound():
// 			// If the socket inbound channel is closed, this goroutine has no purpose.
// 			if !open {
// 				// Logger.Printf("Client[%v]: Inbound channel has been closed", sock.conn.RemoteAddr())
// 				return
// 			}

// 			switch payload.(type) {
// 			case *ConnectionResponse:
// 				res := payload.(*ConnectionResponse)

// 				if res.Channel != channel {
// 					sock.Send(&DisconnectRequest{res.Channel, 0, res.Host})
// 				}

// 			case *DisconnectRequest:
// 				req := payload.(*DisconnectRequest)

// 				if req.Channel == channel {
// 					return
// 				}

// 			case *ConnectionStateResponse:
// 				res := payload.(*ConnectionStateResponse)

// 				if res.Channel == channel {
// 					select {
// 					case <-ctx.Done():
// 						return

// 					case stateResChan <- res:
// 					}
// 				}

// 			case *TunnelRequest:
// 				req := payload.(*TunnelRequest)

// 				if req.Channel == channel {
// 					// Acknowledge tunnel request
// 					err := sock.Send(&TunnelResponse{channel, req.SeqNumber, 0})
// 					if err != nil {
// 						// Logger.Printf("Client[%v]: Error while sending tunnel response: %v",
// 						//               sock.conn.RemoteAddr(), err)
// 						return
// 					}

// 					// Relay to user if it fits the sequence
// 					if req.SeqNumber == incomingSeqNumber {
// 						select {
// 						case <-ctx.Done():
// 							return

// 						case inbound <- req.Payload:
// 						}

// 						incomingSeqNumber++
// 					}
// 				}

// 			case *TunnelResponse:
// 				res := payload.(*TunnelResponse)

// 				if res.Channel == channel {
// 					select {
// 					case <-ctx.Done():
// 						return

// 					case tunRes <- res:
// 					}
// 				}
// 			}
// 		}
// 	}
// }

// //
// func clientOutboundWorker(
// 	ctx     context.Context,
// 	reaper  context.CancelFunc,
// 	tunRes  <-chan *TunnelResponse,
// 	seq     chan<- byte,
// 	ack     chan<- byte,
// ) {
// 	defer reaper()
// 	defer close(seq)
// 	defer close(ack)

// 	var seqNumber byte = 0

// 	outerLoop:
// 	for {
// 		select {
// 		case <-ctx.Done():
// 			return

// 		// Client requests a sequence number because it wants to send something
// 		case seq <- seqNumber:

// 			for {
// 				select {
// 				case <-ctx.Done():
// 					return

// 				// Await tunnel response
// 				case res := <-tunRes:
// 					// We're only interested in the ones that match our sequence number
// 					if res.SeqNumber == seqNumber {
// 						select {
// 						case <-ctx.Done():
// 							return

// 						// Send result of the tunnel request to the sender
// 						case ack <- res.Status:
// 							seqNumber++
// 							continue outerLoop
// 						}
// 					}
// 				}
// 			}

// 		// Discard out-of-cycle tunnel responses
// 		case <-tunRes:
// 		}
// 	}
// }

// //
// func clientHeartbeatWorker(
// 	ctx     context.Context,
// 	reaper  context.CancelFunc,
// 	sock    Socket,
// 	channel byte,
// 	trigger <-chan struct{},
// 	resChan <-chan *ConnectionStateResponse, // Really? Why not just 'byte'?
// ) {
// 	// Logger.Printf("Client[%v]: Started heartbeat worker", sock.conn.RemoteAddr())
// 	// defer Logger.Printf("Client[%v]: Stopped heartbeat worker", sock.conn.RemoteAddr())

// 	// Make sure we tell the others to exit
// 	defer reaper()

// 	outerLoop:
// 	for {
// 		select {
// 		// Gorouting has been asked to exit
// 		case <-ctx.Done():
// 			return

// 		// Inbound worker has triggered a heartbeat
// 		case <-trigger:
// 			req := &ConnectionStateRequest{channel, 0, HostInfo{}}

// 			err := sock.Send(req)
// 			if err != nil {
// 				// Logger.Printf("Client[%v]: Error while sending heartbeat: %v",
// 				//               sock.conn.RemoteAddr(), err)
// 				return
// 			}

// 			// Heartbeat cycle
// 			for i := 0; i < 20; i++ {
// 				select {
// 				case <-ctx.Done():
// 					return

// 				case res := <-resChan:
// 					if res.Status == 0 {
// 						// Logger.Printf("Client[%v]: Heartbeat successful", sock.conn.RemoteAddr())
// 						continue outerLoop
// 					} else {
// 						// Logger.Printf("Client[%v]: Gateway rejected heartbeat",
// 						//               sock.conn.RemoteAddr())
// 						return
// 					}

// 				case <-time.After(500 * time.Millisecond):
// 					err := sock.Send(req)
// 					if err != nil {
// 						// Logger.Printf("Client[%v]: Error while sending heartbeat: %v",
// 						//               sock.conn.RemoteAddr(), err)
// 						return
// 					}
// 				}
// 			}

// 			// We get here, if the gateway did not respond

// 			// Logger.Printf("Client[%v]: Gateway timed out during heartbeat", sock.conn.RemoteAddr())
// 			return

// 		case <-resChan:
// 			// Discard any connection state response that appears out-of-cycle
// 		}
// 	}
// }
