DELETE FROM posts
WHERE title IN (
    'Notícia demonstrativa: cidades ampliam áreas verdes urbanas',
    'Notícia demonstrativa: nova rota de ciclovias entra em planejamento',
    'Notícia demonstrativa: escolas recebem laboratórios de ciências',
    'Notícia demonstrativa: feira local destaca pequenos produtores',
    'Notícia demonstrativa: sistema de ônibus terá painel de horários'
);

DELETE FROM users
WHERE email = 'demo.posts@example.test'
  AND hash_password = 'migration-seed-only-not-a-login-password'
  AND NOT EXISTS (
      SELECT 1
      FROM posts
      WHERE posts.fk_user_id = users.id
  );
