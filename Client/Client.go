package main

import (
	"context"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	pb "MandatoryActivity5/MandatoryActivity5/Node.go"

	"google.golang.org/grpc"
)

func main() {
	// Set up logging to a common file
	logFile, logErr := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if logErr != nil {
		log.Fatalf("failed to open log file: %v", logErr)
	}
	defer logFile.Close()
	log.SetOutput(logFile)

	nodes := []string{"localhost:50051", "localhost:50052", "localhost:50053"}

	var wg sync.WaitGroup
	bidders := []string{"Alice", "Bob"}

	for _, bidder := range bidders {
		wg.Add(1)
		go func(bidder string) {
			defer wg.Done()
			for {
				select {
				case <-time.After(100 * time.Second):
					return
				default:
					// Random delay between 1 and 5 seconds
					time.Sleep(time.Duration(rand.Intn(5)+1) * time.Second)

					// Randomly select a node
					node := nodes[rand.Intn(len(nodes))]
					conn, err := grpc.Dial(node, grpc.WithInsecure(), grpc.WithBlock())
					if err != nil {
						log.Printf("Failed to connect to %s: %v", node, err)
						continue
					}
					defer conn.Close()
					c := pb.NewAuctionClient(conn)

					ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
					defer cancel()

					// Get the current highest bid
					resultResp, err := c.Result(ctx, &pb.ResultRequest{})
					if err != nil {
						log.Printf("could not get result: %v", err)
						continue
					}

					highestBid := resultResp.GetHighestbid()

					// Check if the auction is over
					if strings.HasPrefix(highestBid, "Auction over.") {
						log.Printf("Auction result: %s", highestBid)
						return
					}

					// Place a new bid higher than the current highest bid
					currentHighestBid, err := strconv.Atoi(highestBid)
					if err != nil {
						log.Printf("could not convert highest bid to int: %v", err)
						continue
					}
					newBidAmount := currentHighestBid + 1
					bidResp, err := c.Bid(ctx, &pb.BidRequest{Bidder: bidder, Amount: int32(newBidAmount)})
					if err != nil {
						log.Printf("could not bid: %v", err)
						continue
					}
					log.Printf("Bidder %s bid %d on node %s: %s", bidder, newBidAmount, node, bidResp.Message)
				}
			}
		}(bidder)
	}

	wg.Wait()
}
