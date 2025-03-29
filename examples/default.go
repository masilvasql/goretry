package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/masilvasql/goretry"
	"net/http"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := goretry.Do(ctx, getCepRequest)

	if err != nil {
		fmt.Println("Falha após múltiplas tentativas:", err)
	} else {
		fmt.Println("Sucesso:", result)
	}
}
func getCepRequest(ctx context.Context) (string, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://viacep.com.br/wss/01001000/json/", nil)
	if err != nil {
		return "", err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	fmt.Println("Status code:", resp.StatusCode)

	if resp.StatusCode != http.StatusOK {
		return "", errors.New("falha na requisição")
	}

	return "success", nil
}
