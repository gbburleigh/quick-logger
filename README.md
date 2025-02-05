# Log Streaming Library (Go)

This library provides a simple and efficient way to stream logs to a user-provided gRPC server. It handles the gRPC connection, streaming logic, and message formatting, allowing users to easily integrate log streaming into their Go applications.

## Features

*   **gRPC Streaming:** Utilizes gRPC for high-performance, low-latency log delivery.
*   **User-Provided Destination:** Streams logs to a gRPC server address specified by the user via an environment variable.
*   **Structured Logging:** Sends logs as structured `LogEntry` messages, including timestamp, level, message, and metadata.
*   **Easy Integration:** Simple API for sending logs from your Go code.
*   **Error Handling:** Robust error handling to ensure reliable log delivery.

## Getting Started

### Prerequisites

*   Go (version 1.20 or later recommended)
*   Protocol Buffer compiler (`protoc`)
*   Go gRPC and Protocol Buffer plugins for `protoc`

### Installation

1.  **Install the Protocol Buffer compiler (`protoc`):** Follow the instructions for your operating system: [https://grpc.io/docs/protoc-installation/](https://grpc.io/docs/protoc-installation/)

2.  **Install the Go gRPC and Protocol Buffer plugins:**

    ```bash
    go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
    ```

    Make sure your `$GOPATH/bin` or `$GOBIN` is in your `PATH`.

3.  **Get the library:**

    ```bash
    go get github.com/gbburleigh/log-streaming-library]([invalid URL removed]) // Replace with your repository path
    ```

### Usage

1.  **Set the gRPC server address:** Set the `GRPC_SERVER_ADDRESS` environment variable to the address of the gRPC server that will receive the logs (e.g., `localhost:50051`).

    ```bash
    export GRPC_SERVER_ADDRESS=localhost:50051 # Linux/macOS
    set GRPC_SERVER_ADDRESS=localhost:50051  # Windows
    ```

2.  **Import the library:**

    ```go
    import "[github.com/gbburleigh/log-streaming-library/streamer [invalid URL removed]"
    ```

3.  **Create a `LogStreamer` instance and send logs:**

    ```go
    package main

    import (
        "fmt"
        "log"
        "os"
        "time"

        "github.com/gbburleigh/log-streaming-library/streamer [invalid URL removed]"
    )

    func main() {
        os.Setenv("GRPC_SERVER_ADDRESS", "localhost:50051") // Only for this example, users should set this themselves

        ls, err := streamer.NewLogStreamer()
        if err != nil {
            log.Fatal(err)
        }
        defer ls.Close() // Important to close the stream when finished

        for i := 0; i < 10; i++ {
            err := ls.SendLog("INFO", fmt.Sprintf("Log message #%d", i), map[string]string{"source": "example"})
            if err != nil {
                log.Println(err) // Handle errors appropriately
            }
            time.Sleep(time.Second)
        }
    }
    ```

### Protocol Buffer Definition

The log message format is defined in `proto/logservice.proto`:

```protobuf
syntax = "proto3";

package logging;

option go_package = "github.com/gbburleigh/log-streaming-library/proto [invalid URL removed]";

service LogService {
  rpc StreamLogs(stream LogEntry) returns (stream LogResponse);
}

message LogEntry {
  string timestamp = 1;
  string level = 2;
  string message = 3;
    map<string, string> metadata = 4;
}

message LogResponse {
  string message = 1;
}