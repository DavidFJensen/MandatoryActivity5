package main

import (
	"context"
	"log"
	"math/rand"
	"strconv"
	"time"

	pb "MandatoryActivity5/MandatoryActivity5/Node.go"

	"google.golang.org/grpc"
)

func main() {
	nodes := []string{"localhost:50051", "localhost:50052", "localhost:50053"}
	var conn *grpc.ClientConn
	var err error

	for _, addr := range nodes {
		conn, err = grpc.Dial(addr, grpc.WithInsecure(), grpc.WithBlock())
		if err == nil {
			break
		}
		log.Printf("Failed to connect to %s: %v", addr, err)
	}

	if conn == nil {
		log.Fatalf("could not connect to any node")
	}
	defer conn.Close()
	c := pb.NewAuctionClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	// Start automatic bidding
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				// Random delay between 1 and 5 seconds
				time.Sleep(time.Duration(rand.Intn(5)+1) * time.Second)

				// Get the current highest bid
				resultResp, err := c.Result(ctx, &pb.ResultRequest{})
				if err != nil {
					log.Printf("could not get result: %v", err)
					continue
				}

				// Place a new bid higher than the current highest bid
				currentHighestBid, err := strconv.Atoi(resultResp.GetHighestbid())
				if err != nil {
					log.Printf("could not convert highest bid to int: %v", err)
					continue
				}
				newBidAmount := currentHighestBid + 1
				bidResp, err := c.Bid(ctx, &pb.BidRequest{Bidder: "Alice", Amount: int32(newBidAmount)})
				if err != nil {
					log.Printf("could not bid: %v", err)
					continue
				}
				log.Printf("Bid response: %s", bidResp.Message)
			}
		}
	}()

	// Wait for the auction to end and get the result
	<-ctx.Done()
	resultResp, err := c.Result(context.Background(), &pb.ResultRequest{})
	if err != nil {
		log.Fatalf("could not get result: %v", err)
	}
	log.Printf("Result response: %s", resultResp.Highestbid)
}
