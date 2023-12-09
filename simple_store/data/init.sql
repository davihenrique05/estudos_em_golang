CREATE TABLE produtos (
    id serial primary key,
    nome varchar,
    descricao TEXT,
    preco decimal,
    quantidade integer
);

INSERT INTO produtos (nome, descricao, preco, quantidade) values 
('Camiseta', 'Preta', 19.90, 10),
('Fone', 'Muito bom', 99.00, 5);