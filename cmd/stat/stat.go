package main

import (
	"fmt"
	"os"

	"github.com/dustywilson/stathat"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Provide access token as first argument.  See also https://www.stathat.com/access")
	}
	if len(os.Args) < 3 {
		fmt.Println("Provide stat name as second argument.")
		os.Exit(1)
	}
	s := stathat.New().Token(os.Args[1])
	stat, err := s.Stat(os.Args[2])
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Printf("[%s] %s\n", stat.ID, stat.Name)
}
