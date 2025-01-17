# Cashback Rewards System

This project implements a blockchain-based cashback rewards system using Solidity smart contracts, a Golang backend, and a Next.js frontend.

## Features

- Smart contracts for cashback and token points using ERC20 tokens.
- Golang backend for API handling and contract interaction.
- Next.js frontend for user interaction.
- Integration with Ethereum blockchain.

## Project Structure

## Getting Started

### Prerequisites

- Node.js
- Golang
- Truffle or Hardhat
- Ethereum wallet (e.g., MetaMask)

### Steps to Run

1. **Smart Contracts**
   - Navigate to the `cashback-rewards` directory.
   - Compile and deploy contracts:
     ```bash
     truffle compile
     truffle migrate --network <network>
     ```

2. **Backend**
   - Navigate to the `go-backend` directory.
   - Run the server:
     ```bash
     go run cmd/main.go
     ```

3. **Frontend**
   - Navigate to the `frontend` directory.
   - Start the development server:
     ```bash
     npm install
     npm run dev
     ```

4. Access the app at `http://localhost:3000`.

---

## API Endpoints

### `POST /purchase`
- **Request**: `{ "productValue": 100 }`
- **Response**: `{ "message": "Purchase successful", "cashbackAmount": 10, "tokenPoints": 10 }`

### `GET /cashback`
- **Query**: `address=0x...`
- **Response**: `{ "address": "0x...", "cashbackAmount": 100 }`

---

## License

MIT