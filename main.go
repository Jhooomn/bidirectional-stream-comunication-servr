package main

import (
	"io"
	"log"
	"net"

	"github.com/Jhooomn/bidirectional-stream-comunication/servr/protos"
	"google.golang.org/grpc"
)

// server implements the CalculatorServiceServer interface.
type server struct {
	protos.UnimplementedCalculatorServiceServer
}

const (
	xIncrement = 2
	yIncrement = 0
	zIncrement = 0
)

// Calculate processes the stream of CalculateRequest messages and responds with
// CalculateResponse messages that contain updated coordinates. It applies a fixed
// increment to the X, Y, and Z values of the request and sends back the result.
func (s *server) Calculate(stream protos.CalculatorService_CalculateServer) error {
	for {
		// Receive a message from the stream.
		req, err := stream.Recv()
		if err == io.EOF {
			// End of stream.
			return nil
		}
		if err != nil {
			log.Printf("error receiving from stream: %v", err)
			return err
		}

		// Create a response with the updated coordinates.
		resp := &protos.CalculateResponse{
			X: req.GetX() + xIncrement,
			Y: req.GetY() + yIncrement,
			Z: req.GetZ() + zIncrement,
		}

		// Send the response back to the client.
		if err := stream.Send(resp); err != nil {
			log.Printf("error sending to stream: %v", err)
			return err
		}
	}
}

// main sets up the gRPC server and listens for incoming connections on port 50051.
// It registers the server as a CalculatorServiceServer to handle incoming RPC calls.
func main() {
	// Listen for TCP connections on port 50051.
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Create a new gRPC server instance.
	s := grpc.NewServer()

	// Register the server with the gRPC server to handle CalculatorService RPCs.
	protos.RegisterCalculatorServiceServer(s, &server{})

	// Start serving incoming connections.
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
