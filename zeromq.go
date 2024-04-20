package zeromq

import (
	"context"
	"fmt"
	"log/slog"
	"strings"
	"time"

	"go.k6.io/k6/metrics"

	zmq "github.com/go-zeromq/zmq4"
)

var socketBuilder = map[string]func(ctx context.Context, opts ...zmq.Option) zmq.Socket{
	"dealer": zmq.NewDealer,
	"req":    zmq.NewReq,
	"push":   zmq.NewPush,
	"pair":   zmq.NewPair,
	"pub":    zmq.NewPub,
	"xpub":   zmq.NewXPub,
}

// NewSocket creates a new ZeroMQ socket of type `socketType` connected to `addr`.
func (z *ZeroMQ) NewSocket(addr string, socketType string) (*zmq.Socket, error) {
	builder, ok := socketBuilder[strings.ToLower(socketType)]
	if !ok {
		var validTypes []string
		for k := range socketBuilder {
			validTypes = append(validTypes, k)
		}
		return nil, fmt.Errorf("%s is not a valid socket type. Valid socket types: %s", socketType, strings.Join(validTypes, ", "))
	}
	sock := builder(context.Background())
	err := sock.Dial(addr)
	if err != nil {
		slog.Error("Could not dial remote", slog.Any("error", err))
		return nil, err
	}
	return &sock, nil
}

// Send sends a message containing `data` to `socket`.
func (z *ZeroMQ) Send(socket zmq.Socket, data string) (string, error) {
	// Send the message
	msg := zmq.NewMsgString(data)
	err := socket.Send(msg)
	sentAt := time.Now()
	if err != nil {
		slog.Error("Could not send message", slog.Any("error", err))
		return "", err
	}

	// Receiving the response.
	resp, err := socket.Recv()
	if err != nil {
		slog.Error("Could not receive message", slog.Any("error", err))
		z.reportMetric(z.metrics.FailedRequestCount, time.Now(), 1)
		return "", err
	}
	response := resp.String()

	now := time.Now()
	z.reportMetric(z.metrics.RequestDuration, now, metrics.D(now.Sub(sentAt)))
	z.reportMetric(z.metrics.RequestCount, now, 1)
	z.reportMetric(z.metrics.FailedRequestCount, now, 0)

	return response, nil
}

// CloseSocket closes the socket.
func (z *ZeroMQ) CloseSocket(socket zmq.Socket) {
	socket.Close()
}
