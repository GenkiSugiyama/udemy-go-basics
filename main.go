package main

import (
	"fmt"
	"go-basics/calculator"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	fmt.Println(os.Getenv("GO_ENV"))
	fmt.Println(calculator.Offset)
	fmt.Println(calculator.Sum(3, 3))
	fmt.Println(calculator.Multiply(2, 2))
}
