# Distributed Auction System

## Introduction

This project implements a distributed auction system using replication to handle auctions, ensuring resilience to one crash failure.

## Architecture

The system consists of multiple nodes running on distinct processes. Clients can direct API requests to any node. The nodes communicate using gRPC and replicate bids to ensure resilience.

Running the System
1. Start the nodes:
-find the server folder
-open three terminals and launch each server with the following lines:
go run server.go 50051
go run server.go 50052
go run server.go 50053

2. Start the client:
-find the client folder
-open a terminal and launch the client with the following line:
    go run Client.go
  
