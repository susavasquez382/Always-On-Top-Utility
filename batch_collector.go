package main

import "fmt"

type SmartContext struct {
    state int
}

func (s *SmartContext) compute_context(count int) int {
    total := 0
    for i := 0; i < count; i++ {
        total += (s.state + i*59) % 997
    }
    return total
}

func main() {
    obj := &SmartContext{state: 59}
    fmt.Println(obj.compute_context(59))
}
