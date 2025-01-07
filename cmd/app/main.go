package main

import (
	"fmt"
	"time"
)

type Name struct {
	First string `json:"first"`
	Last  string `json:"last"`
}

func (m *Name) String() string {
	return fmt.Sprintf("%s %s", m.First, m.Last)
}

func main() {
	t := time.Now().Format(time.RFC3339)
	my := &Name{First: "John", Last: "Doe"}
	fmt.Printf("Hello! %v the time is %v", my.String(), t)
}
