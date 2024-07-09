# DB-KV-GO

Este projeto é uma implementação de um armazenamento de chave-valor simples utilizando Go, Gin para o servidor web, e um armazenamento em memória. Ele fornece uma API REST para definir, obter, listar e excluir valores.

## Estrutura do Projeto

```
|-- routes
|   |-- routes.go
|-- server
|   |-- server.go
|   |-- server_test.go
|-- storage
|   |-- storage.go
|-- main.go
```

## Instalação

1. Clone o repositório:

   ```bash
   git clone https://github.com/Lucasbyte/DB-KV-GO.git
   ```

2. Navegue até o diretório do projeto:

   ```bash
   cd DB-KV-GO
   ```

3. Instale as dependências:

   ```bash
   go mod tidy
   ```

## Execução

Para iniciar o servidor, execute:

```bash
go run main.go
```

O servidor estará disponível em `http://localhost:3010`.

## Endpoints

### `GET /ping`

Verifica se o servidor está funcionando.

#### Resposta de Sucesso

```json
{
  "message": "pong-go"
}
```

### `POST /set`

Define um novo valor no armazenamento.

#### Corpo da Requisição

```json
{
  "metodo": "SET",
  "key": "chave",
  "value": "valor"
}
```

#### Resposta de Sucesso

```json
{
  "message": "Value set successfully!",
  "key": "chave",
  "value": "valor"
}
```

### `GET /get`

Obtém um valor do armazenamento.

#### Corpo da Requisição

```json
{
  "metodo": "GET",
  "key": "chave"
}
```

#### Resposta de Sucesso

```json
{
  "message": "Value get successfully!",
  "value": "valor"
}
```

### `GET /all`

Obtém todos os pares chave-valor do armazenamento.

#### Corpo da Requisição

```json
{
  "metodo": "ALL"
}
```

#### Resposta de Sucesso

```json
{
  "message": "Values get successfully!",
  "values": [
    {
      "key": "chave1",
      "value": "valor1"
    },
    {
      "key": "chave2",
      "value": "valor2"
    }
  ]
}
```

### `DELETE /delete`

Exclui um valor do armazenamento.

#### Corpo da Requisição

```json
{
  "metodo": "DEL",
  "key": "chave"
}
```

#### Resposta de Sucesso

```json
{
  "message": "Value DELETE successfully!"
}
```

## Testes

Para rodar os testes, execute:

```bash
go test ./...
```

Os testes estão localizados em `server/server_test.go` e cobrem as principais funcionalidades do servidor.

## Contribuições

Contribuições são bem-vindas! Sinta-se à vontade para abrir issues ou enviar pull requests.
