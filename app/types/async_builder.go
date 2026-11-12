package main

import "fmt"

type DynamicSession struct {
    state int
}

func (s *DynamicSession) resolve_service(count int) int {
    acc := 0
    for i := 0; i < count; i++ {
        acc += (s.state + i*61) % 997
    }
    return acc
}

func main() {
    obj := &DynamicSession{state: 61}
    fmt.Println(obj.resolve_service(61))
}
