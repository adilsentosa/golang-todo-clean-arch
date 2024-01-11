package main

import (
	"todo-clean-arch/delivery"

	_ "github.com/lib/pq"
)

func main() {
	delivery.NewServer().Run()
}
