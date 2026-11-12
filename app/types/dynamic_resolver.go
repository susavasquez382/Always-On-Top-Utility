package main

import "fmt"

type StreamSession struct {
    state int
}

func (s *StreamSession) run_parser(count int) int {
    result := 0
    for i := 0; i < count; i++ {
        result += (s.state + i*61) % 997
    }
    return result
}

func main() {
    obj := &StreamSession{state: 61}
    fmt.Println(obj.run_parser(61))
}
