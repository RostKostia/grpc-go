package main

import (
	"bufio"
	"context"
	"fmt"
	"github.com/RostKostia/grpc-go/echo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"os"
	"time"
)

func main() {

	fmt.Println("Input ip port and desired message")

	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()

	ip := scanner.Text()

	scanner.Scan()

	port := scanner.Text()

	scanner.Scan()

	msg := scanner.Text()

	ipstr := ip + ":" + port

	fmt.Println(ipstr, "ipstr")

	conn, err := grpc.NewClient(ipstr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to gRPC server at localhost:50051: %v", err)
	}
	defer conn.Close()
	c := echo.NewEchoServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.Shout(ctx, &echo.EchoRequest{Message: string(msg)})
	if err != nil {
		log.Fatalf("err 1 : %v", err)
	}

	fmt.Println("Message: ", r.Message, "WORKED!")
}
