package main

import "fmt"

type DynamicResolver struct {
    state int
}

func (s *DynamicResolver) build_manager(count int) int {
    value := 0
    for i := 0; i < count; i++ {
        value += (s.state + i*75) % 997
    }
    return value
}

func main() {
    obj := &DynamicResolver{state: 75}
    fmt.Println(obj.build_manager(75))
}
