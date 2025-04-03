# Blueis

Blueis is a lightweight TCP server that listens on port `7171` and supports basic commands like `PING`, `SET`, and `GET`. It uses an append-only file (AOF) for persistence.

## Features
- **PING Command**: Responds with `PONG` or a custom message.
- **SET Command**: Stores a key-value pair in memory.
- **GET Command**: Retrieves the value for a given key.
- **Persistence**: Uses an AOF file (`database.aof`) to persist data.

## Folder Structure
```
blueis/
├── blueis-sdk/      # SDK for interacting with the Blueis server
│   └── blueis-sdk.go # SDK implementation
├── server/          # Contains the Go source files
│   ├── main.go      # Entry point of the application
│   ├── aof.go       # Handles append-only file persistence
│   ├── resp.go      # RESP protocol implementation
│   ├── handler.go   # Command handlers (PING, SET, GET)
├── Dockerfile       # Dockerfile for building and running the application
├── README.md        # Project documentation
└── go.mod           # Go module file
```

## Requirements
- **Go**: Version 1.18 or higher
- **Docker**: Installed and configured

## Running Locally
1. Clone the repository:
   ```bash
   git clone <repository-url>
   cd blueis
   ```

2. Build and run the application:
   ```bash
   go run server/main.go
   ```

## Running with Docker
1. Build the Docker image:
   ```bash
   docker build -t blueis .
   ```

2. Run the Docker container:
   ```bash
   docker pull pjk010/blueis:latest
   ```

## Usage
You can interact with the server using a TCP client like `netcat` or a custom client.

### Example Commands
- **PING**:
  ```bash
  nc localhost 7171
  *1
  $4
  PING
  ```
  Response:
  ```
  +PONG
  ```

- **SET**:
  ```bash
  nc localhost 7171
  *3
  $3
  SET
  $3
  key
  $5
  value
  ```
  Response:
  ```
  +OK
  ```

- **GET**:
  ```bash
  nc localhost 7171
  *2
  $3
  GET
  $3
  key
  ```
  Response:
  ```
  $5
  value
  ```

## License
This project is licensed under the MIT License.
