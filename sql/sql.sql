CREATE DATABASE IF NOT EXISTS fazendadb;
USE fazendadb;

DROP TABLE IF EXISTS usuarios;
DROP TABLE IF EXISTS fazendas;

CREATE TABLE fazendas (
    id INT AUTO_INCREMENT PRIMARY KEY,
    nome VARCHAR(100) NOT NULL,
    endereco VARCHAR(255),
    codigo_fazenda VARCHAR(10) NOT NULL UNIQUE,
    data_criacao TIMESTAMP DEFAULT CURRENT_TIMESTAMP
) ENGINE=INNODB;


CREATE TABLE usuarios (
    id INT AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(100) NOT NULL UNIQUE,
    nome VARCHAR(50) NOT NULL,
    sobrenome VARCHAR(50) NOT NULL,
    email VARCHAR(60) NOT NULL UNIQUE,
    senha VARCHAR(255) NOT NULL,
    tipo_usuario ENUM('A', 'U') NOT NULL DEFAULT 'U', 
    criadoEm TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    fazenda_id INT NOT NULL,
    CONSTRAINT fk_usuario_fazenda 
    FOREIGN KEY (fazenda_id) REFERENCES fazendas(id)
) ENGINE=INNODB;

ALTER TABLE fazendas 
ADD COLUMN dono_id INT,
ADD CONSTRAINT fk_fazenda_dono 
FOREIGN KEY (dono_id) REFERENCES usuarios(id);
