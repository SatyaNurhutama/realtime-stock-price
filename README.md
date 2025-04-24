# Real-Time Stock Price API

A real-time stock price checking API built using **gRPC**, **Go**, **gRPC Streaming**, and **Alpha Vantage API**. This project provides both a unary RPC to retrieve a stock price and a streaming RPC to continuously provide real-time stock price updates.

## Overview

This project implements a **gRPC**-based API to check real-time stock prices. It integrates with the **Alpha Vantage API** to retrieve stock prices using their public API. The API exposes two main methods:
1. **GetStockPrice**: A unary RPC that returns the current stock price for a given symbol.
2. **StreamStockPrice**: A streaming RPC that continuously provides the latest stock price updates at regular intervals (default: every 60 seconds).

## Features

- Real-time stock price retrieval.
- Support for **gRPC Streaming** for continuous updates.
- Built with **Go** and follows **clean code architecture**.
- Integration with **Alpha Vantage API** for stock data.

## Technologies

- **Go**: The backend language for building the API.
- **gRPC**: For high-performance remote procedure calls and streaming.
- **Alpha Vantage API**: For retrieving stock prices and historical data.
- **Protocol Buffers (Protobuf)**: For defining the service and message types.

## Setup

### Environment Variables

Before running the project, make sure to set up the `.env` file with following environment variable:

- `ALPHA_VANTAGE_API_KEY`: Your Alpha Vantage API key (get it [here](https://www.alphavantage.co/support/#api-key)).

### API Documentation
**Get Stock Price (Unary RPC)**

**Example Request:**
```
grpcurl -d '{"symbol": "AAPL"}' -plaintext localhost:50051 stock_price.StockPriceService/GetStockPrice
```
**Example Response:**
```
{
  "symbol": "AAPL",
  "price": 204.4598,
  "timestamp": "2025-04-23 19:59:00"
}
```
**Stream Stock Price (Streaming RPC)**

**Example Request**
```
grpcurl -d '{"symbol": "AAPL"}' -plaintext localhost:50051 stock_price.StockPriceService/GetStockPriceStream
```
**Example Response:**
```
{
  "symbol": "AAPL",
  "price": 204.4598,
  "timestamp": "2025-04-23 20:00:00"
}
```
