-- Create a demo owner only when there are no users to own the sample posts.
INSERT INTO users (name, email, hash_password)
SELECT 'Usuário de demonstração', 'demo.posts@example.test', 'migration-seed-only-not-a-login-password'
WHERE NOT EXISTS (SELECT 1 FROM users)
  AND NOT EXISTS (SELECT 1 FROM posts);

-- Seed sample news only when the posts table is empty.
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
        'Notícia demonstrativa: cidades ampliam áreas verdes urbanas',
        'Um programa piloto está apoiando o plantio de árvores e a criação de pequenos parques em bairros com pouca cobertura vegetal. A iniciativa será avaliada ao longo do próximo ano.'
    ),
    (
        'Notícia demonstrativa: nova rota de ciclovias entra em planejamento',
        'A prefeitura iniciou estudos para conectar trechos de ciclovias já existentes. O projeto ainda está em fase de consulta pública e não tem data definida para início das obras.'
    ),
    (
        'Notícia demonstrativa: escolas recebem laboratórios de ciências',
        'Dez escolas públicas começaram a receber equipamentos para atividades práticas de ciências. Professores também participarão de oficinas de formação durante o semestre.'
    ),
    (
        'Notícia demonstrativa: feira local destaca pequenos produtores',
        'Uma feira semanal passou a reunir produtores de diferentes regiões. A organização pretende acompanhar a participação do público e dos comerciantes nos próximos meses.'
    ),
    (
        'Notícia demonstrativa: sistema de ônibus terá painel de horários',
        'Um novo painel digital está sendo testado em alguns pontos de ônibus para informar previsões de chegada. O serviço está em fase experimental e poderá receber ajustes.'
    )
) AS sample(title, message)
WHERE NOT EXISTS (SELECT 1 FROM posts);
