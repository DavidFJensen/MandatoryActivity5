package main

import (
    "context"
    "fmt"
    "log"
    "net"
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
        return &pb.BidResponse{Message: "fail"}, nil
    }

    s.highestBid = req.Amount
    s.highestBidder = req.Bidder

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
        return &pb.ResultResponse{Highestbid: fmt.Sprintf("Highest bid: %d by %s", s.highestBid, s.highestBidder)}, nil
    }

    return &pb.ResultResponse{Highestbid: fmt.Sprintf("Auction over. Winner: %s with bid %d", s.highestBidder, s.highestBid)}, nil
}

func main() {
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }

    grpcServer := grpc.NewServer()
    pb.RegisterAuctionServer(grpcServer, NewAuctionServer())

    log.Printf("server listening at %v", lis.Addr())
    if err := grpcServer.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}