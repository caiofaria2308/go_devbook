CREATE DATABASE IF NOT EXISTS devbook;
USE devbook;

DROP TABLE IF EXISTS usuarios;
DROP TABLE IF EXISTS seguidores;


CREATE TABLE IF NOT EXISTS usuarios (
    id int AUTO_INCREMENT PRIMARY KEY,
    nome varchar(50) not null,
    nick varchar(50) not null unique,
    email varchar(50) not null unique,
    senha varchar(100) not null,
    criadoEm timestamp default current_timestamp()
) ENGINE=INNODB;

CREATE TABLE IF NOT EXISTS seguidores(
    usuario_id int not null,
    seguidor_id int not null, 
    primary key(usuario_id, seguidor_id ),
    foreign key(usuario_id) REFERENCES usuarios(id) ON DELETE CASCADE,
    foreign key(seguidor_id) REFERENCES usuarios(id) ON DELETE CASCADE
);