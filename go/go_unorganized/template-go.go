package main

import (
        "fmt"
)

type Probe interface {
        ~int|~float64
}

func add[P Probe] (x, y P) P {
        return x + y
}

func main() {
        i1, i2 := 1, 2
        f1, f2 := 1.0, 2.0
        fmt.Println(add(i1,i2))
        fmt.Println(add(f1,f2))
}
