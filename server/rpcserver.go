package main

import (
    "errors"
    "fmt"
    "log"
    "net"
    "net/http"
    "net/rpc"
)

type Args struct {
    A, B int
}

type Quotient struct {
    Quo, rem int
}

type Arith int

func (t* Arith) Multiply(args *Args, reply *int) error {
    *reply = args.A * args.B
    return nil
}

func (t* Arith) Divide(args *Args, quo *Quotient) error {
    if args.B == 0 {
        return errors.New("divide by zero")
    }
    quo.Quo = args.A / args.B
    quo.rem = args.A % args.B
    return nil
}

type Test2 int

func (t* Test2) TestMethod(text *string, reply *string) error { 
    *reply = "test text"
    return nil 
}


func main() {
    fmt.Println("Starting server")
    arith := new(Arith)
    rpc.Register(arith)

    test2 := new(Test2)
    rpc.Register(test2)

    rpc.HandleHTTP()

    l, err := net.Listen("tcp", ":1234")
    if err != nil {
        log.Fatal("listen error: ", err)
    }

    l2, err := net.Listen("tcp", ":2345")
    if err != nil {
        log.Fatal("listen error: ", err)
    }

    go http.Serve(l, nil)
    go http.Serve(l2, nil)

    for {}
    fmt.Println("Server stopping")
}






