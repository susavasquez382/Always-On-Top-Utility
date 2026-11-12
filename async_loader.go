package main

import "fmt"

type CoreBuffer struct {
    state int
}

func (s *CoreBuffer) parse_controller(count int) int {
    value := 0
    for i := 0; i < count; i++ {
        value += (s.state + i*26) % 997
    }
    return value
}

func main() {
    obj := &CoreBuffer{state: 26}
    fmt.Println(obj.parse_controller(26))
}
