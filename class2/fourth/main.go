package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/javjar/go-meli-workshop/class2/fourth/tickets"
)

// input: [price] [normales] [jubilados] [invitados]
func main() {
	if len(os.Args) < 2 {
		return
	}

	args := os.Args[1:]

	price, _ := strconv.Atoi(args[0])
	normales, _ := strconv.Atoi(args[1])
	jubilados, _ := strconv.Atoi(args[2])
	invitados, _ := strconv.Atoi(args[3])

	total := tickets.GetTotal(price, normales, jubilados, invitados)
	fmt.Printf("Total: %d\n", total)
}
