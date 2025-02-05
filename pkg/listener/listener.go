package listener

import (
        "context"
        "fmt"
        "google.golang.org/grpc"
        "log"
        "net"

    pb "github.com/gbburleigh/quick-logger/proto"
)

type server struct {
        pb.UnimplementedLogServiceServer
}

func (s *server) StreamLogs(stream pb.LogService_StreamLogsServer) error {
        for {
                in, err := stream.Recv()
                if err != nil {
                        return err
                }

                fmt.Printf("Received Log: Timestamp: %s, Level: %s, Message: %s, Metadata: %v\n", in.Timestamp, in.Level, in.Message, in.Metadata)

                // Optionally send an acknowledgment back to the client
                if err := stream.Send(&pb.LogResponse{Message: "Log received"}); err != nil {
                        return err
                }
        }
}

func main() {
        lis, err := net.Listen("tcp", ":50051") // Listen on port 50051
        if err != nil {
                log.Fatalf("failed to listen: %v", err)
        }

        s := grpc.NewServer()
        pb.RegisterLogServiceServer(s, &server{})

        log.Println("Server listening on :50051")
        if err := s.Serve(lis); err != nil {
                log.Fatalf("failed to serve: %v", err)
        }
}