package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	pb "github.com/verbruggenjesse/event-store/event-operator-test-client/gen"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	var serverAddr string

	if serverAddr = os.Getenv("EVENT_OPERATOR_ADDR"); serverAddr == "" {
		log.Fatalln("Missing required environment variable 'EVENT_OPERATOR_ADDR'")
	}

	conn, err := grpc.Dial(serverAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatal("could not connect to operator service")
	}
	defer conn.Close()

	client := pb.NewEventOperatorClient(conn)

	payload, _ := json.Marshal(map[string]interface{}{
		"content": "Hello streams!",
	})

	event := &pb.Event{
		Topic:   "message",
		Action:  "echo",
		Payload: string(payload),
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	res, err := client.Publish(ctx, event)

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("%v", res)
}
