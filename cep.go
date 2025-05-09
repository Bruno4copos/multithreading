package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

type Endereco struct {
	CEP        string `json:"cep"`
	Logradouro string `json:"logradouro"`
	Bairro     string `json:"bairro"`
	Localidade string `json:"localidade"`
	UF         string `json:"uf"`
	Origem     string
}

func fetchFromAPI(ctx context.Context, url string, origem string, ch chan<- Endereco) {
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil || resp.StatusCode != 200 {
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}

	var e Endereco
	if err := json.Unmarshal(body, &e); err != nil {
		return
	}
	e.Origem = origem
	select {
	case ch <- e:
	case <-ctx.Done():
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Uso: go run main.go <CEP>")
		return
	}

	cep := os.Args[1]
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	ch := make(chan Endereco, 2)

	go fetchFromAPI(ctx, "https://brasilapi.com.br/api/cep/v1/"+cep, "BrasilAPI", ch)
	go fetchFromAPI(ctx, "http://viacep.com.br/ws/"+cep+"/json/", "ViaCEP", ch)

	select {
	case resultado := <-ch:
		fmt.Printf("Resultado recebido da %s:\n", resultado.Origem)
		fmt.Printf("CEP: %s\nLogradouro: %s\nBairro: %s\nCidade: %s\nEstado: %s\n",
			resultado.CEP, resultado.Logradouro, resultado.Bairro, resultado.Localidade, resultado.UF)
	case <-ctx.Done():
		fmt.Println("Erro: Timeout ao buscar o CEP")
	}
}
