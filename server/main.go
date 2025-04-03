package main

import (
    "fmt"
    "net"
    "strings"
)

func main() {
    fmt.Println("Listening on port :7171")

    // Create a new server
    l, err := net.Listen("tcp", ":7171")
    if err != nil {
        fmt.Println(err)
        return
    }

    aof, err := NewAof("database.aof")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer aof.Close()

    aof.Read(func(value Value) {
        command := strings.ToUpper(value.array[0].bulk)
        args := value.array[1:]

        handler, ok := Handlers[command]
        if !ok {
            fmt.Println("Invalid command: ", command)
            return
        }

        handler(args)
    })

    // Listen for connections
    for {
        conn, err := l.Accept()
        if err != nil {
            fmt.Println("Connection error:", err)
            continue
        }
        go handleConnection(conn, aof) // Handle each client concurrently
    }
}

func handleConnection(conn net.Conn, aof *Aof) {
    defer conn.Close()

    for {
        resp := NewResp(conn)
        value, err := resp.Read()
        if err != nil {
            fmt.Println("Read error:", err)
            return // Close connection if read fails
        }

        if value.typ != "array" || len(value.array) == 0 {
            fmt.Println("Invalid request")
            continue
        }

        command := strings.ToUpper(value.array[0].bulk)
        args := value.array[1:]

        writer := NewWriter(conn)

        handler, ok := Handlers[command]
        if !ok {
            fmt.Println("Invalid command:", command)
            writer.Write(Value{typ: "string", str: ""})
            continue
        }

        if command == "SET" {
            aof.Write(value)
        }

        result := handler(args)
        writer.Write(result)
    }
}
