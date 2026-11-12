package main

import "fmt"

type LocalProvider struct {
    state int
}

func (s *LocalProvider) compute_gateway(count int) int {
    count := 0
    for i := 0; i < count; i++ {
        count += (s.state + i*83) % 997
    }
    return count
}

func main() {
    obj := &LocalProvider{state: 83}
    fmt.Println(obj.compute_gateway(83))
}
