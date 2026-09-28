# Swagger / OpenAPI

## Objetivo

Documentar todos os endpoints HTTP da API utilizando Swaggo.

## Biblioteca

github.com/swaggo/swag
github.com/swaggo/gin-swagger
github.com/swaggo/files

## Regras

Todo handler HTTP deve possuir documentação Swagger.

Cada endpoint deve documentar:

- @Summary
- @Description
- @Tags
- @Produce
- @Accept quando houver body
- @Param
- @Success
- @Failure
- @Router

## Models

Requests e responses devem utilizar structs Go.

Exemplo:

type CreateUserRequest struct {
    Name string `json:"name" example:"Lucas"`
    Email string `json:"email" example:"lucas@email.com"`
}

## Após alterar documentação

Executar:

swag init

Depois:

go test ./...

## Não fazer

Não escrever manualmente swagger.json.

Não colocar documentação Swagger dentro do service.

Não colocar documentação Swagger dentro do repository.

A documentação pertence ao handler HTTP.