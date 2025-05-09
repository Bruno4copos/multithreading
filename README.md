# Desafio de Concorrência de APIs com Go

Este programa consulta duas APIs de CEP em paralelo e retorna a resposta mais rápida.

## 📦 Funcionalidade

- Consulta os serviços:
  - `https://brasilapi.com.br/api/cep/v1/{CEP}`
  - `http://viacep.com.br/ws/{CEP}/json/`
- Aceita a primeira resposta e ignora a outra.
- Timeout total de 1 segundo.
- Exibe os dados do endereço e a origem da resposta.

## ▶️ Execução

### Requisitos
- Go 1.18+

### Rodar o programa

```bash
go run main.go 01153000
```

### Resultado esperado

```bash
Resultado recebido da BrasilAPI:
CEP: 01153-000
Logradouro: Rua General Osório
Bairro: Campos Elísios
Cidade: São Paulo
Estado: SP
```

### Timeout

Se nenhuma resposta for recebida em 1 segundo:

```bash
Erro: Timeout ao buscar o CEP
```
