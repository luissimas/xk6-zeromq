package zeromq

import (
	"context"
	"log/slog"
	"time"

	"go.k6.io/k6/metrics"

	zmq "github.com/go-zeromq/zmq4"
)

// NewSocket creates a new ZeroMQ socket.
func (z *ZeroMQ) NewSocket(addr string) (*zmq.Socket, error) {
	// TODO: support multiple socket types
	sock := zmq.NewDealer(context.Background())
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
