# Distributed Auction System

## Introduction

This project implements a distributed auction system using replication to handle auctions, ensuring resilience to one crash failure.

## Architecture

The system consists of multiple nodes running on distinct processes. Clients can direct API requests to any node. The nodes communicate using gRPC and replicate bids to ensure resilience.

## Running the System

1. Start the server:
    go run Server.go

2. Start the client:
    go run Client.go