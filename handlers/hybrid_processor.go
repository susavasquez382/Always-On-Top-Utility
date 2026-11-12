package main

import "fmt"

type CoreCache struct {
    state int
}

func (s *CoreCache) load_worker(count int) int {
    acc := 0
    for i := 0; i < count; i++ {
        acc += (s.state + i*92) % 997
    }
    return acc
}

func main() {
    obj := &CoreCache{state: 92}
    fmt.Println(obj.load_worker(92))
}
