package main

import (
	"fmt"
	"log"
	"net/rpc"
)

type Args struct {
    A, B int
}
    
func main() {
    fmt.Println("dialling server")
    client, err := rpc.DialHTTP("tcp", "localhost:1234")
    if err != nil {
        log.Fatal("dialing:", err)
    }
    fmt.Println("connected to the server")
    args := Args{7, 8}   
    var reply int
    fmt.Printf("calling server")
    err = client.Call("Arith.Multiply", args, &reply)
    if err != nil {
        log.Fatal("Arith error:", err)
    }
    fmt.Printf("Arith: %d*%d=%d\n", args.A, args.B, reply)

    fmt.Println("Setting up second connection")
    client2, err := rpc.DialHTTP("tcp", "localhost:2345")
    if err != nil {
        log.Fatal("dialing:", err)
    }
    textArg := "testText"
    var textReply string    
    err = client2.Call("Test2.TestMethod", textArg, &textReply)
    if err != nil {
        log.Fatal("Test2 error:", err)
    }
    fmt.Printf("Test2 result %s", textReply)
}
