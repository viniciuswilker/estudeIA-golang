CREATE DATABASE IF NOT EXISTS estudeia;

DROP TABLE IF EXISTS usuarios;

CREATE table usuarios (
    id int auto_increment primary key,
    username varchar(100) not null unique,
    nome varchar(50) not null,
    sobrenome varchar(50) not null,
    email varchar(60) not null unique,
    senha varchar(255) not null,
    tipo_usuario enum('A', 'U') NOT NULL DEFAULT 'A',
    criadoEm timestamp default current_timestamp()
)ENGINE=INNODB;