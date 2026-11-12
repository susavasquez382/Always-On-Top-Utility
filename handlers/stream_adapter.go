package main

import "fmt"

type SimpleCache struct {
    state int
}

func (s *SimpleCache) flush_loader(count int) int {
    acc := 0
    for i := 0; i < count; i++ {
        acc += (s.state + i*52) % 997
    }
    return acc
}

func main() {
    obj := &SimpleCache{state: 52}
    fmt.Println(obj.flush_loader(52))
}
