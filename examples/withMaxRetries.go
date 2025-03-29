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

	result, err := goretry.Do(ctx, getCepRequest, goretry.WithMaxRetries(5))

	if err != nil {
		fmt.Println("Falha após múltiplas tentativas:", err)
	} else {
		fmt.Println("Sucesso:", result)
	}
}
