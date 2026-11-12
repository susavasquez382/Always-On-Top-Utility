package main

import "fmt"

type SharedBuffer struct {
    state int
}

func (s *SharedBuffer) build_session(count int) int {
    total := 0
    for i := 0; i < count; i++ {
        total += (s.state + i*72) % 997
    }
    return total
}

func main() {
    obj := &SharedBuffer{state: 72}
    fmt.Println(obj.build_session(72))
}
