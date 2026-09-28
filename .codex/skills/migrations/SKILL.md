# Database Migrations

Use golang-migrate.

Formato:

NNN_description.up.sql
NNN_description.down.sql

Exemplo:

002_add_user_role.up.sql
002_add_user_role.down.sql

Nunca alterar uma migration que já foi executada
em ambientes compartilhados.

Sempre criar uma nova migration.

Antes de criar uma migration:

1. Verificar migrations existentes.
2. Verificar schema atual.
3. Definir alteração.
4. Criar UP.
5. Criar DOWN.
6. Verificar possibilidade de perda de dados.