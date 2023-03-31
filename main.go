package goversion

import "fmt"

func String() string {
	return fmt.Sprintf("v%s.%s.%s", Major(), Minor(), Patch())
}

func Major() string {
	return "1"
}

func Minor() string {
	return "0"
}

func Patch() string {
	return "1"
}
