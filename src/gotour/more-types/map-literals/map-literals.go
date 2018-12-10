package main

import "fmt"

type Vertex struct {
    Lat, Long float64
}

var m = map[string]Vertex{
    "Bell Labs": Vertex{
        40.68433, -122.08408,
    },
    "Google": Vertex{
        37.42202, -74.39967,
    },
}

func main() {
    fmt.Println(m)
}
