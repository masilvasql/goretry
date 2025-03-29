package main

import (
	"context"
	"fmt"
	"github.com/masilvasql/goretry"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := goretry.Do(ctx, getCepRequest, goretry.WithShouldRetry(func(err error) bool {
		if err.Error() == "falha na requisição" {
			return true
		}
		return false
	}))

	if err != nil {
		fmt.Println("Falha após múltiplas tentativas:", err)
	} else {
		fmt.Println("Sucesso:", result)
	}
}
