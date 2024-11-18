package mandatoryactivity5

import (
	"context"
	"fmt"
	"sync"
	"time"

	pb "MandatoryActivity5/MandatoryActivity5/Node.go"
)

type Node struct {
	nodeID int
	value  int
}

type AuctionServer struct {
	nodes         []*Node
	mu            sync.Mutex
	highestBid    int32
	highestBidder string
	startTime     time.Time
}

func NewAuctionServer(isLeader bool) *AuctionServer {
	return &AuctionServer{
		nodes:     []*Node{},
		startTime: time.Now(),
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
		node.value = int(s.highestBid)
	}

	return &pb.BidResponse{Message: "success"}, nil

}

func (s *AuctionServer) Result(ctx context.Context, req *pb.ResultRequest) (*pb.ResultResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if time.Since(s.startTime) < 100*time.Second {
		return &pb.ResultResponse{Highestbid: fmt.Sprintf("Highest bid: %d by %s", s.highestBid, s.highestBidder)}, nil
	}

	return &pb.ResultResponse{Highestbid: fmt.Sprintf("Auction over. Winner: %s with bid %d", s.highestBidder, s.highestBid)}, nil
}

func (s *AuctionServer) isAuctionActive() bool {
	if time.Since(s.startTime) > 100*time.Second {
		return false
	}
	return true
}
