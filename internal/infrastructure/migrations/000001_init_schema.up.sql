-- Создание таблицы users для хранения данных пользователей
CREATE TABLE users (
    id           SERIAL PRIMARY KEY,            -- Уникальный идентификатор пользователя
    email        VARCHAR(50) UNIQUE NOT NULL,   -- Электронная почта пользователя
    password     VARCHAR(64) NOT NULL,          -- Хэш пароля пользователя
    count        FLOAT NOT NULL,                -- Кол-во денег на счету
    freeze_count FLOAT NOT NULL                 -- Кол-во замороженных денег на счету
);

-- Создание таблицы lots для хранения данных о лотах
CREATE TABLE lots (
    id          SERIAL PRIMARY KEY,                                -- Уникальный идентификатор лота
    name        VARCHAR(50) NOT NULL,                              -- Название лота
    owner_id    INT NOT NULL,                                      -- Идентификатор владельца лота
    parameters  TEXT,                                              -- Дополнительные параметры лота (в формате JSON или другом)

    FOREIGN KEY (owner_id) REFERENCES users (id) ON DELETE CASCADE -- Внешний ключ для владельца лота
);

-- Создание таблицы auctions для хранения данных об аукционах
CREATE TABLE auctions (
    id          SERIAL PRIMARY KEY,                               -- Уникальный идентификатор аукциона
    lot_id      INT UNIQUE NOT NULL,                              -- Идентификатор лота
    min_step    INT NOT NULL,                                     -- Минимальный шаг ставки
    expires     TIMESTAMP NOT NULL,                               -- Время истечения аукциона

    FOREIGN KEY (lot_id) REFERENCES lots (id) ON DELETE CASCADE   -- Внешний ключ для лота
);

-- Создание индекса для быстрого поиска по lot_id в таблице auctions
CREATE INDEX idx_lot_id ON auctions(lot_id);

-- Создание таблицы bids для хранения данных о ставках
CREATE TABLE bids (
    id          SERIAL PRIMARY KEY,                                     -- Уникальный идентификатор ставки
    user_id     INT NOT NULL,                                           -- Идентификатор пользователя, сделавшего ставку
    auction_id  INT NOT NULL,                                           -- Идентификатор аукциона, на который сделана ставка
    bid         INT NOT NULL,                                           -- Сумма ставки

    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE,      -- Внешний ключ для пользователя
    FOREIGN KEY (auction_id) REFERENCES auctions (id) ON DELETE CASCADE -- Внешний ключ для аукциона
);

-- Создание индекса для быстрого поиска по user_id в таблице bids
CREATE INDEX idx_user_id ON bids(user_id);

-- Создание индекса для быстрого поиска по auction_id в таблице bids
CREATE INDEX idx_auction_id ON bids(auction_id);