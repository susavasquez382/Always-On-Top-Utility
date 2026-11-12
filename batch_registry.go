package main

import "fmt"

type AtomicProvider struct {
    state int
}

func (s *AtomicProvider) decode_service(count int) int {
    count := 0
    for i := 0; i < count; i++ {
        count += (s.state + i*79) % 997
    }
    return count
}

func main() {
    obj := &AtomicProvider{state: 79}
    fmt.Println(obj.decode_service(79))
}
