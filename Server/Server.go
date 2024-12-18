package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"sync"
	"time"

	pb "MandatoryActivity5/MandatoryActivity5/Node.go"

	"google.golang.org/grpc"
)

type Node struct {
	nodeID int
	value  int
	addr   string
	active bool
}

type AuctionServer struct {
	pb.UnimplementedAuctionServer
	nodes         []*Node
	mu            sync.Mutex
	highestBid    int32
	highestBidder string
	startTime     time.Time
}

func NewAuctionServer() *AuctionServer {
	server := &AuctionServer{
		nodes:     []*Node{},
		startTime: time.Now(),
	}
	go server.healthCheck()
	return server
}

func (s *AuctionServer) healthCheck() {
	for {
		time.Sleep(10 * time.Second)
		s.mu.Lock()
		for _, node := range s.nodes {
			conn, err := grpc.Dial(node.addr, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(2*time.Second))
			if err != nil {
				node.active = false
				log.Printf("Node %d is down", node.nodeID)
			} else {
				node.active = true
				conn.Close()
				log.Printf("Node %d is active", node.nodeID)
			}
		}
		s.mu.Unlock()
	}
}

func (s *AuctionServer) Bid(ctx context.Context, req *pb.BidRequest) (*pb.BidResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if req.Amount <= s.highestBid {
		log.Printf("Bid from %s with amount %d failed", req.Bidder, req.Amount)
		return &pb.BidResponse{Message: "fail"}, nil
	}

	s.highestBid = req.Amount
	s.highestBidder = req.Bidder
	log.Printf("Bid from %s with amount %d succeeded", req.Bidder, req.Amount)

	// Replicate bid to other nodes
	for _, node := range s.nodes {
		if node.active {
			go s.replicateBid(node, req)
		}
	}

	return &pb.BidResponse{Message: "success"}, nil
}

func (s *AuctionServer) replicateBid(node *Node, req *pb.BidRequest) {
	conn, err := grpc.Dial(node.addr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Printf("Failed to connect to node %d: %v", node.nodeID, err)
		return
	}
	defer conn.Close()

	client := pb.NewAuctionClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = client.Bid(ctx, req)
	if err != nil {
		log.Printf("Failed to replicate bid to node %d: %v", node.nodeID, err)
	}
}

func (s *AuctionServer) Result(ctx context.Context, req *pb.ResultRequest) (*pb.ResultResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if time.Since(s.startTime) < 100*time.Second {
		log.Printf("Current highest bid: %d by %s", s.highestBid, s.highestBidder)
		return &pb.ResultResponse{Highestbid: fmt.Sprintf("%d", s.highestBid)}, nil
	}

	log.Printf("Auction over. Winner: %s with bid %d", s.highestBidder, s.highestBid)
	return &pb.ResultResponse{Highestbid: fmt.Sprintf("Auction over. Winner: %s with bid %d", s.highestBidder, s.highestBid)}, nil
}

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("Usage: %s <port>", os.Args[0])
	}
	port := os.Args[1]

	// Set up logging to a file
	logFile, err := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("failed to open log file: %v", err)
	}
	defer logFile.Close()
	log.SetOutput(logFile)

	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	server := NewAuctionServer()
	server.nodes = []*Node{
		{nodeID: 1, addr: "localhost:50051", active: true},
		{nodeID: 2, addr: "localhost:50052", active: true},
		{nodeID: 3, addr: "localhost:50053", active: true},
	}
	pb.RegisterAuctionServer(grpcServer, server)

	log.Printf("server listening at %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
