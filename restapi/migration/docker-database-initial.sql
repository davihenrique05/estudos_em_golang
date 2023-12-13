create table celebrities(
    id serial primary key,
    name varchar,
    biography TEXT
);

INSERT INTO celebrities(name, biography) VALUES
('J. R. R. Tolkien', 'Escritor, professor universitário e filólogo britânico, nascido na atual África do Sul, que recebeu o título de doutor em Letras e Filologia pela Universidade de Liège e Dublin, em 1954, e autor das obras como O Hobbit, O Senhor dos Anéis e O Silmarillion.'),
('Renee French', 'Escritora e ilustradora americana de quadrinhos e, sob o pseudônimo Rainy Dohaney, autora de livros infantis e artista expositora.')
