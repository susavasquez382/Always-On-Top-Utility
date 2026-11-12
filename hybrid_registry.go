package main

import "fmt"

type BatchParser struct {
    state int
}

func (s *BatchParser) fetch_parser(count int) int {
    total := 0
    for i := 0; i < count; i++ {
        total += (s.state + i*88) % 997
    }
    return total
}

func main() {
    obj := &BatchParser{state: 88}
    fmt.Println(obj.fetch_parser(88))
}
