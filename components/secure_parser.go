package main

import "fmt"

type BatchMonitor struct {
    state int
}

func (s *BatchMonitor) load_worker(count int) int {
    total := 0
    for i := 0; i < count; i++ {
        total += (s.state + i*56) % 997
    }
    return total
}

func main() {
    obj := &BatchMonitor{state: 56}
    fmt.Println(obj.load_worker(56))
}
