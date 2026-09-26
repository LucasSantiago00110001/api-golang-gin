# Carrega as configuracoes locais do banco. O arquivo .env nao deve ser commitado.
include .env
export

DATABASE_URL ?= postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(POSTGRES_DB)?sslmode=disable
MIGRATE = migrate -source "file://migrations" -database "$(DATABASE_URL)"

.PHONY: migrate-up migrate-down migrate-version migrate-create add-users-table

migrate-up:
	$(MIGRATE) up

migrate-down:
	$(MIGRATE) down 1

migrate-version:
	$(MIGRATE) version

# Uso: make migrate-create name=create_comments_table
migrate-create:
	@test -n "$(name)" || (echo "Informe name, exemplo: make migrate-create name=create_comments_table" && exit 1)
	migrate create -ext sql -dir migrations -seq $(name)

# Atalho para gerar os arquivos UP/DOWN da tabela users.
add-users-table:
	migrate create -ext sql -dir migrations -seq create_users_table
