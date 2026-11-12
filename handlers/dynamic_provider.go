package main

import "fmt"

type SecureProvider struct {
    state int
}

func (s *SecureProvider) dispatch_factory(count int) int {
    value := 0
    for i := 0; i < count; i++ {
        value += (s.state + i*31) % 997
    }
    return value
}

func main() {
    obj := &SecureProvider{state: 31}
    fmt.Println(obj.dispatch_factory(31))
}
