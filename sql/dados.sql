use devbook;

insert into usuarios(nome, nick, email, senha)
VALUES
('admin', 'admin', 'admin@admin', '$2a$10$woXjyuJYpL0yesOPiorTuOOhTH9wU8mg3YpkM4JlGJyOKHFVxk9Qu'),
('usuario 1', 'user1', 'user1@admin', '$2a$10$woXjyuJYpL0yesOPiorTuOOhTH9wU8mg3YpkM4JlGJyOKHFVxk9Qu'),
('usuario 2', 'user2', 'user2@admin', '$2a$10$woXjyuJYpL0yesOPiorTuOOhTH9wU8mg3YpkM4JlGJyOKHFVxk9Qu'),
('usuario 3', 'user3', 'user3@admin', '$2a$10$woXjyuJYpL0yesOPiorTuOOhTH9wU8mg3YpkM4JlGJyOKHFVxk9Qu'),
('usuario 4', 'user4', 'user4@admin', '$2a$10$woXjyuJYpL0yesOPiorTuOOhTH9wU8mg3YpkM4JlGJyOKHFVxk9Qu'),
('usuario 5', 'user5', 'user5@admin', '$2a$10$woXjyuJYpL0yesOPiorTuOOhTH9wU8mg3YpkM4JlGJyOKHFVxk9Qu'),
("caio", "blackfcaio", "caio@admin", '$2a$10$woXjyuJYpL0yesOPiorTuOOhTH9wU8mg3YpkM4JlGJyOKHFVxk9Qu');

insert into seguidores (usuario_id, seguidor_id) values 
(1, 2),
(1, 3),
(1, 4),
(1, 5),
(1, 6),
(1, 7),
(2, 1),
(2, 3),
(2, 4),
(2, 5),
(2, 6),
(2, 7);