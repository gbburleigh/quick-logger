package streamer

import (
	"context"
	"fmt"
	"os"
	"time"

	pb "github.com/gbburleigh/quick-logger/proto/logservice"
	"google.golang.org/grpc"
)

type LogStreamer struct {
	client pb.LogServiceClient
	stream pb.LogService_StreamLogsClient
	ctx    context.Context
	cancel context.CancelFunc
}

func NewLogStreamer() (*LogStreamer, error) {
	address := os.Getenv("GRPC_SERVER_ADDRESS")
	if address == "" {
		return nil, fmt.Errorf("GRPC_SERVER_ADDRESS environment variable not set")
	}

	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("did not connect: %w", err)
	}

	c := pb.NewLogServiceClient(conn)
	ctx, cancel := context.WithCancel(context.Background())

	stream, err := c.StreamLogs(ctx)
	if err != nil {
		return nil, fmt.Errorf("Could not connect to stream: %w", err)
	}

	return &LogStreamer{
		client: c,
		stream: stream,
		ctx:    ctx,
		cancel: cancel,
	}, nil
}

func (ls *LogStreamer) SendLog(level, message string, metadata map[string]string) error {
	logEntry := &pb.LogEntry{
		Timestamp: time.Now().String(),
		Level:     level,
		Message:   message,
		Metadata:  metadata,
	}

	if err := ls.stream.Send(logEntry); err != nil {
		return fmt.Errorf("could not send log: %w", err)
	}

	resp, err := ls.stream.Recv()
	if err != nil {
		return fmt.Errorf("Could not receive message: %w", err)
	}

	fmt.Printf("Server Response: %s\n", resp.Message)
	return nil
}

func (ls *LogStreamer) Close() error {
	ls.cancel()
	if err := ls.stream.CloseSend(); err != nil {
		return fmt.Errorf("could not close stream: %w", err)
	}
	return nil
}
