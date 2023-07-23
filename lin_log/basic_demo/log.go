package main

import (
	"log"
	"os"
)

func main() {
	log.Println("www")
	{
		// output: TEST  2023/07/23 18:45:36 ttt
		l := log.New(os.Stderr, "TEST  ", log.LstdFlags)
		l.Println("ttt")
	}
}
