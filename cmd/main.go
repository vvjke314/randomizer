package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1]
	num, err := strconv.Atoi(args)
	if err != nil {
		log.Println(err)
		panic("Can't convert args[1] to Int")
	}
	fmt.Println(rand.Intn(num))
}
