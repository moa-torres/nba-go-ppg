package main

import (
	"github.com/moacirtorres/nba-go-ppg/server"
)

func main() {

	s := server.NewServer()

	s.Run()

}
