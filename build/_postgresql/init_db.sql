DROP TABLE IF EXISTS Cart_Items;
DROP TABLE IF EXISTS Users;
DROP TABLE IF EXISTS Items;
DROP TABLE IF EXISTS Sellers;
DROP TABLE IF EXISTS Catalogs;


CREATE TABLE Catalogs (
                         ID SERIAL PRIMARY KEY,
                         Name VARCHAR NOT NULL,
                         parent_id INTEGER,
                         CONSTRAINT fk_parent
                             FOREIGN KEY (parent_id)
                                 REFERENCES Catalogs(ID)
                                 ON DELETE SET NULL
);

CREATE TABLE Sellers (
                        ID SERIAL PRIMARY KEY,
                        Name VARCHAR NOT NULL,
                        Deals INTEGER NOT NULL
);

CREATE TABLE Items (
                      ID SERIAL PRIMARY KEY,
                      Name VARCHAR,
                      in_stock_value INTEGER NOT NULL,
                      seller_id INTEGER NOT NULL,
                      parent_id INTEGER NOT NULL,
                      CONSTRAINT fk_seller
                          FOREIGN KEY (seller_id)
                              REFERENCES Sellers(ID)
                              ON DELETE CASCADE,
                      CONSTRAINT fk_catalog
                          FOREIGN KEY (parent_id)
                              REFERENCES Catalogs(ID)
                              ON DELETE CASCADE
);

INSERT INTO Sellers (Name, Deals) VALUES
                                     ('Нефритовая Лиса', 6),
                                     ('Дядюшка Ляо', 2),
                                     ('Издательство Питер', 12),
                                     ('Издательство Вильямс', 8),
                                     ('Издательство ДМК Пресс', 4);

INSERT INTO Catalogs (ID, Name, parent_id) VALUES
    (1, 'ShopQL', NULL);

INSERT INTO Catalogs (ID, Name, parent_id) VALUES
    (2, 'Книги', 1);

INSERT INTO Catalogs (ID, Name, parent_id) VALUES
    (3, 'Алгоритмы', 2);

INSERT INTO Catalogs (ID, Name, parent_id) VALUES
    (4, 'Golang', 2);

INSERT INTO Catalogs (ID, Name, parent_id) VALUES
    (5, 'Чай', 1);

INSERT INTO Items (ID, Name, in_stock_value, seller_iD, parent_id) VALUES
(1, 'Грокаем алгоритмы | Бхаргава Адитья', 1, 3, 3),
(2, 'Теоретический минимум по Computer Science | Фило Владстон Феррейра', 2, 3, 3),
(3, 'Совершенный алгоритм. Основы | Рафгарден Тим', 3, 3, 3),
(4, 'Алгоритмы на Java | Джитер Кевин Уэйн, Седжвик Роберт', 4, 4, 3),
(5, 'Язык программирования Go | Донован Алан А. А., Керниган Брайан У.', 1, 4, 4),
(6, 'Go на практике | Butcher Matt, Фарина Мэтт Мэтт', 2, 5, 4),
(7, 'Программирование на Go. Разработка приложений XXI века | Саммерфильд Марк', 3, 5, 4),
(8, 'Head First. Изучаем Go | Макгаврен Джей', 4, 3, 4),
(9, 'Си Пу Юань, Шен Пуэр', 1, 2, 5),
(10, 'Мэнхай 7542, Шен Пуэр', 2, 2, 5),
(11, 'Дянь Хун', 3, 2, 5),
(12, 'Да Хун Пао', 5, 2, 5),
(13, 'Габа Улун', 4, 1, 5);
