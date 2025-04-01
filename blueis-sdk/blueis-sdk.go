package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

type CacheClient struct {
	hostname string
	port     string
}

func NewCacheClient(hostname string) *CacheClient {
	return &CacheClient{
		hostname: hostname,
		port:     "7171",
	}
}

func (c *CacheClient) sendCommand(command string) (string, error) {
	conn, err := net.Dial("tcp", c.hostname+":"+c.port)
	if err != nil {
		return "", err
	}
	defer conn.Close()

	_, err = conn.Write([]byte(command))
	if err != nil {
		return "", err
	}

	reader := bufio.NewReader(conn)
	response, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(response), nil
}

func (c *CacheClient) Put(key, value string) bool {
	command := fmt.Sprintf("*3\r\n$3\r\nSET\r\n$%d\r\n%s\r\n$%d\r\n%s\r\n", len(key), key, len(value), value)
	response, err := c.sendCommand(command)
	if err != nil {
		fmt.Println("Error:", err)
		return false
	}
	return response == "+OK"
}

func (c *CacheClient) Get(key string) (string, error) {
	command := fmt.Sprintf("*2\r\n$3\r\nGET\r\n$%d\r\n%s\r\n", len(key), key)
	response, err := c.sendCommand(command)
	if err != nil {
		return "", err
	}

	if strings.HasPrefix(response, "$-1") { // Redis-style nil response
		return "", nil
	}

	return response, nil
}

func main() {
	client := NewCacheClient("localhost")

	if client.Put("test_key", "test_value") {
		fmt.Println("Value stored successfully!")
	} else {
		fmt.Println("Failed to store value.")
	}

	value, err := client.Get("test_key")
	if err != nil {
		fmt.Println("Error retrieving value:", err)
	} else {
		fmt.Println("Retrieved value:", value)
	}
}
