CREATE TABLE posts (
    id SERIAL PRIMARY KEY,
    fk_user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    title VARCHAR(255) NOT NULL,
    message TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_posts_fk_user_id ON posts(fk_user_id);

-- Use an existing user for the sample records so the foreign key stays valid.
WITH sample_user AS (
    SELECT id
    FROM users
    ORDER BY id
    LIMIT 1
)
INSERT INTO posts (fk_user_id, title, message)
SELECT sample_user.id, sample.title, sample.message
FROM sample_user
CROSS JOIN (VALUES
    (
        'Projeto de energia solar amplia acesso à eletricidade no interior',
        'Uma iniciativa regional instalou painéis solares em comunidades afastadas. A expectativa é reduzir custos e ampliar o acesso à energia durante os próximos anos.'
    ),
    (
        'Pesquisadores desenvolvem material que melhora a filtragem de água',
        'Uma equipe universitária apresentou um novo material para filtros domésticos. Os testes iniciais indicam maior retenção de partículas, e novas avaliações independentes estão previstas.'
    ),
    (
        'Bibliotecas públicas recebem acervo digital gratuito',
        'Bibliotecas de diferentes cidades começaram a oferecer uma plataforma de empréstimo de livros digitais. O serviço pode ser acessado com o cadastro local de cada unidade.'
    ),
    (
        'Agricultores testam técnicas para economizar água nas lavouras',
        'Produtores participantes de um programa piloto estão usando sensores de umidade e irrigação localizada. Os resultados serão avaliados ao final da próxima safra.'
    ),
    (
        'Transporte elétrico entra em fase de testes em linhas urbanas',
        'Uma cidade iniciou testes com ônibus elétricos em duas linhas. A operação experimental deve medir consumo, autonomia e desempenho nos trajetos diários.'
    )
) AS sample(title, message);
