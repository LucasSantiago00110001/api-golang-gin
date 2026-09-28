# Sannockit API

## Stack

- Go
- Gin
- PostgreSQL 16
- database/sql
- golang-migrate
- Docker
- Makefile

## Arquitetura

Cada feature deve seguir:

Handler
  ↓
Service
  ↓
Repository
  ↓
Database

## Regras

- Não utilizar ORM.
- Utilizar database/sql.
- Handlers não devem conter regras de negócio.
- Services contêm regras de negócio.
- Repositories são responsáveis pelo acesso ao banco.
- SQL deve ficar no repository.
- Erros devem ser tratados explicitamente.
- Context.Context deve ser propagado.
- APIs devem retornar JSON.
- Testes devem ser criados para novas funcionalidades.

## Usuários

A feature posts está localizada em:

internal/features/posts

Arquivos esperados:

dto.go
handler.go
service.go
repository.go
model.go
module.go

## Banco

As alterações no banco devem ser feitas através de migrations.

Nunca modificar uma migration já aplicada.

Criar uma nova migration para alterações estruturais.

## Antes de alterar código

Sempre:

1. Ler os arquivos relacionados.
2. Entender a arquitetura existente.
3. Verificar testes existentes.
4. Fazer a menor alteração necessária.
5. Executar os testes.

## Testes

Depois de alterar código Go:

go test ./...

Se existir gotestsum:

gotestsum --format pkgname