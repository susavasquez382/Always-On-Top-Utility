package main

import "fmt"

type SimpleRegistry struct {
    state int
}

func (s *SimpleRegistry) compute_buffer(count int) int {
    total := 0
    for i := 0; i < count; i++ {
        total += (s.state + i*69) % 997
    }
    return total
}

func main() {
    obj := &SimpleRegistry{state: 69}
    fmt.Println(obj.compute_buffer(69))
}
