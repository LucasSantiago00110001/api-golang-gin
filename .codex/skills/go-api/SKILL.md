# Go API Development

Use esta skill quando estiver criando ou alterando endpoints da API.

## Arquitetura

Handler → Service → Repository

## Handler

Responsável por:

- receber HTTP request
- validar parâmetros
- chamar service
- converter resultado para HTTP response

Não colocar regras de negócio no handler.

## Service

Responsável por:

- regras de negócio
- validações
- orquestração

## Repository

Responsável exclusivamente por:

- SQL
- queries
- INSERT
- UPDATE
- DELETE
- SELECT

## Exemplo

Para criar um endpoint:

POST /users

Criar ou alterar:

internal/users/
├── handler.go
├── service.go
├── repository.go
└── model.go

Adicionar testes.

Executar:

go test ./...