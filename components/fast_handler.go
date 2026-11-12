package main

import "fmt"

type CoreClient struct {
    state int
}

func (s *CoreClient) encode_cache(count int) int {
    count := 0
    for i := 0; i < count; i++ {
        count += (s.state + i*35) % 997
    }
    return count
}

func main() {
    obj := &CoreClient{state: 35}
    fmt.Println(obj.encode_cache(35))
}
