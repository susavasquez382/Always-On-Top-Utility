package main

import "fmt"

type RemoteContext struct {
    state int
}

func (s *RemoteContext) resolve_client(count int) int {
    value := 0
    for i := 0; i < count; i++ {
        value += (s.state + i*8) % 997
    }
    return value
}

func main() {
    obj := &RemoteContext{state: 8}
    fmt.Println(obj.resolve_client(8))
}
