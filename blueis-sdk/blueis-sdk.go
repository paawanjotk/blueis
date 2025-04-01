package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"time"
)

// RedisClient represents a Redis connection instance
type RedisClient struct {
	conn net.Conn
}

// Connect initializes a Redis connection
func Connect(address string) (*RedisClient, error) {
	conn, err := net.DialTimeout("tcp", address, 5*time.Second)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Redis server: %v", err)
	}

	return &RedisClient{conn: conn}, nil
}

// Close closes the Redis connection
func (r *RedisClient) Close() {
	if r.conn != nil {
		r.conn.Close()
	}
}

// sendCommand sends a command to the Redis server in RESP format
func (r *RedisClient) sendCommand(command []string) (string, error) {
	// Build RESP message
	message := fmt.Sprintf("*%d\r\n", len(command))
	for _, part := range command {
		message += fmt.Sprintf("$%d\r\n%s\r\n", len(part), part)
	}

	// Send the command
	_, err := r.conn.Write([]byte(message))
	if err != nil {
		return "", fmt.Errorf("failed to send command: %v", err)
	}

	// Read the server response
	resp, err := bufio.NewReader(r.conn).ReadString('\n')
	if err != nil {
		return "", fmt.Errorf("failed to read response: %v", err)
	}

	return strings.TrimSpace(resp), nil
}

// Set stores a key-value pair in Redis
func (r *RedisClient) Set(key, value string) (string, error) {
	return r.sendCommand([]string{"SET", key, value})
}

// Get retrieves the value of a key from Redis
func (r *RedisClient) Get(key string) (string, error) {
	return r.sendCommand([]string{"GET", key})
}

// Example usage of the Redis SDK
func main() {
	client, err := Connect("localhost:7171")
	if err != nil {
		fmt.Println("Failed to connect:", err)
		return
	}
	defer client.Close()

	// Set a value
	res, err := client.Set("name", "Paawanjot")
	if err != nil {
		fmt.Println("SET failed:", err)
		return
	}
	fmt.Println("SET Response:", res)

	// Get the value
	val, err := client.Get("name")
	if err != nil {
		fmt.Println("GET failed:", err)
		return
	}
	fmt.Println("GET Response:", val)
}
