package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"time"

	pb "github.com/eya20/LogName/personpb"
	"google.golang.org/grpc"
)

func main() {
	// Connect to the server
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// Create a client
	client := pb.NewPersonServiceClient(conn)

	// Create a scanner to read user input
	scanner := bufio.NewScanner(os.Stdin)

	for {
		// Prompt the user for input
		log.Println("Enter a name and surname (format: 'Name Surname'), or type 'exit' to quit:")

		// Read user input
		scanner.Scan()
		input := scanner.Text()

		// Exit if the user types "exit"
		if input == "exit" {
			log.Println("Client exiting.")
			break
		}

		// Parse the input into name and surname
		var name, surname string
		_, err := fmt.Sscanf(input, "%s %s", &name, &surname)
		if err != nil {
			log.Println("Invalid input. Please use the format 'Name Surname'.")
			continue
		}

		// Prepare the request
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		// Send the request
		response, err := client.SendPerson(ctx, &pb.PersonRequest{Name: name, Surname: surname})
		if err != nil {
			log.Printf("Could not send person: %v", err)
			continue
		}

		// Log the response
		log.Printf("Response: %s", response.Message)
	}
}
