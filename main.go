package main

import (
	"encoding/json"
	"fmt"
	"log"

	pb "github.com/verbruggenjesse/event-store/operator-client-test/gen"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const serverAddr = ":3001"

func main() {
	conn, err := grpc.Dial(serverAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatal("could not connect to publisher service")
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
