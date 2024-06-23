-- Таблица пользователей
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY NOT NULL,
    name VARCHAR(64) NOT NULL,
    surname VARCHAR(64) NOT NULL,
    patronymic VARCHAR(64),
    email VARCHAR(64) NOT NULL UNIQUE,
    phone_number VARCHAR(32) NOT NULL UNIQUE,
    password VARCHAR(256) NOT NULL,
    photo UUID,
    created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT (CURRENT_TIMESTAMP AT TIME ZONE 'UTC'),
    updated_at TIMESTAMP WITHOUT TIME ZONE DEFAULT (CURRENT_TIMESTAMP AT TIME ZONE 'UTC')
);

-- Доступы
CREATE TABLE IF NOT EXISTS accesses (
    id SERIAL PRIMARY KEY NOT NULL, 
    access VARCHAR(16) NOT NULL,
    created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT (CURRENT_TIMESTAMP AT TIME ZONE 'UTC'),
    updated_at TIMESTAMP WITHOUT TIME ZONE DEFAULT (CURRENT_TIMESTAMP AT TIME ZONE 'UTC'),
    deleted_at TIMESTAMP WITHOUT TIME ZONE DEFAULT NULL,
    -- Внешний ключ users (one-to-one)
    user_id INT NOT NULL, -- UNIQUE
    -- Связь с таблицей users
    FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE UNIQUE INDEX idx_user_id_not_deleted ON accesses(user_id) WHERE deleted_at IS NULL;
CREATE UNIQUE INDEX idx_user_is_admin ON accesses(user_id) WHERE access = 'ADMIN';

-- Таблица сессий
CREATE TABLE IF NOT EXISTS sessions (
    id SERIAL PRIMARY KEY NOT NULL,
    refresh_token VARCHAR(1024) NOT NULL,
    device VARCHAR(256) NOT NULL,
    created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT (CURRENT_TIMESTAMP AT TIME ZONE 'UTC'),
    updated_at TIMESTAMP WITHOUT TIME ZONE DEFAULT (CURRENT_TIMESTAMP AT TIME ZONE 'UTC'),
    -- Внешний ключ users (one-to-many)
    user_id INT NOT NULL,
    -- Связь с таблицей users
    FOREIGN KEY (user_id) REFERENCES users(id)
);

-- Таблица для восстановления пароля
CREATE TABLE IF NOT EXISTS recoveries (
    id SERIAL PRIMARY KEY NOT NULL,
    code VARCHAR(256) NOT NULL,
    created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT (CURRENT_TIMESTAMP AT TIME ZONE 'UTC'),
    updated_at TIMESTAMP WITHOUT TIME ZONE DEFAULT (CURRENT_TIMESTAMP AT TIME ZONE 'UTC'),
    -- Внешний ключ users (one-to-one)
    user_id INT NOT NULL UNIQUE,
    -- Связь с таблицей users
    FOREIGN KEY (user_id) REFERENCES users(id)
);

-- Таблица попыток входа в аккаунт и попыток восстановления пароля
CREATE TABLE IF NOT EXISTS attempts (
    id SERIAL PRIMARY KEY NOT NULL,
    created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT (CURRENT_TIMESTAMP AT TIME ZONE 'UTC'),
    -- Внешний ключ users (one-to-many)
    user_id INT NOT NULL,
    -- Связь с таблицей users
    FOREIGN KEY (user_id) REFERENCES users(id)
);

-- Таблица подписок пользователей
CREATE TABLE IF NOT EXISTS subscriptions (
    id SERIAL PRIMARY KEY NOT NULL,
    end_at DATE NOT NULL,
    created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT (CURRENT_TIMESTAMP AT TIME ZONE 'UTC'),
    updated_at TIMESTAMP WITHOUT TIME ZONE DEFAULT (CURRENT_TIMESTAMP AT TIME ZONE 'UTC'),
    -- Внешний ключ users (one-to-one)
    user_id INT NOT NULL UNIQUE,
    -- Связь с таблицей users
    FOREIGN KEY (user_id) REFERENCES users(id)
);

-- Таблица сборок
CREATE TABLE IF NOT EXISTS assemblies (
    id SERIAL PRIMARY KEY NOT NULL,
    name VARCHAR(64) NOT NULL,
    file UUID NOT NULL,
    created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT (CURRENT_TIMESTAMP AT TIME ZONE 'UTC'),
    updated_at TIMESTAMP WITHOUT TIME ZONE DEFAULT (CURRENT_TIMESTAMP AT TIME ZONE 'UTC'),
    -- Внешний ключ users (one-to-many)
    user_id INT NOT NULL,
    -- Связь с таблицей users
    FOREIGN KEY (user_id) REFERENCES users(id)
);

-- Таблица с информацией о химическом составе ячеек
CREATE TABLE IF NOT EXISTS compositions (
    id SERIAL PRIMARY KEY NOT NULL,
    name VARCHAR(32) NOT NULL,
    cvc JSON NOT NULL,
    file UUID NOT NULL,
    created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT (CURRENT_TIMESTAMP AT TIME ZONE 'UTC'),
    updated_at TIMESTAMP WITHOUT TIME ZONE DEFAULT (CURRENT_TIMESTAMP AT TIME ZONE 'UTC')
);

-- Таблица производителей двигателей
CREATE TABLE IF NOT EXISTS manufacturers (
    id SERIAL PRIMARY KEY NOT NULL,
    name VARCHAR(64) NOT NULL,
    created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT (CURRENT_TIMESTAMP AT TIME ZONE 'UTC')
);

-- Таблица характеристик ячеек батарей
CREATE TABLE IF NOT EXISTS cells (
    id SERIAL PRIMARY KEY NOT NULL,
    name VARCHAR(64),
    capacity REAL NOT NULL, -- Емкость ячейки
    resistance REAL NOT NULL, -- Сопротивление ячейки
    c_rating JSON NOT NULL, -- C-Рейтинг ячейки
    mass REAL NOT NULL, -- Масса ячейки
    created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT (CURRENT_TIMESTAMP AT TIME ZONE 'UTC'),
    updated_at TIMESTAMP WITHOUT TIME ZONE DEFAULT (CURRENT_TIMESTAMP AT TIME ZONE 'UTC'),
    -- Внешний ключ compositions (one-to-many)
    compositions_id INT NOT NULL,
    -- Связь с таблицей compositions
    FOREIGN KEY (compositions_id) REFERENCES compositions(id),
     -- Внешний ключ manufacturer (one-to-many)
    manufacturer_id INT NOT NULL,
    -- Связь с таблицей manufacturers
    FOREIGN KEY (manufacturer_id) REFERENCES manufacturers(id)
);

-- Таблица с информацией о регуляторах оборотов двигателей
CREATE TABLE IF NOT EXISTS esc (
    id SERIAL PRIMARY KEY NOT NULL,
    name VARCHAR(64),
    current JSON NOT NULL,
    voltage REAL NOT NULL,
    resistance REAL NOT NULL,
    mass REAL NOT NULL,
    multipler SMALLINT NOT NULL,
    created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT (CURRENT_TIMESTAMP AT TIME ZONE 'UTC'),
    updated_at TIMESTAMP WITHOUT TIME ZONE DEFAULT (CURRENT_TIMESTAMP AT TIME ZONE 'UTC'),
     -- Внешний ключ manufacturer (one-to-many)
    manufacturer_id INT NOT NULL,
    -- Связь с таблицей manufacturers
    FOREIGN KEY (manufacturer_id) REFERENCES manufacturers(id)
);

-- Таблица пропеллеров
CREATE TABLE IF NOT EXISTS propellers (
    id SERIAL PRIMARY KEY NOT NULL,
    name VARCHAR(64),
    gear_ratio REAL NOT NULL,
    power_const REAL NOT NULL,
    traction_const REAL NOT NULL,
    mass SMALLINT NOT NULL,
    created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT (CURRENT_TIMESTAMP AT TIME ZONE 'UTC'),
    updated_at TIMESTAMP WITHOUT TIME ZONE DEFAULT (CURRENT_TIMESTAMP AT TIME ZONE 'UTC'),
     -- Внешний ключ manufacturer (one-to-many)
    manufacturer_id INT NOT NULL,
    -- Связь с таблицей manufacturers
    FOREIGN KEY (manufacturer_id) REFERENCES manufacturers(id)
);

-- Таблица двигателей
CREATE TABLE IF NOT EXISTS motors (
    id SERIAL PRIMARY KEY NOT NULL,
    name VARCHAR(64),
    kv SMALLINT NOT NULL,
    current SMALLINT NOT NULL,
    voltage REAL NOT NULL,
    power INT NOT NULL,
    resistance REAL NOT NULL,
    height REAL NOT NULL,
    diameter REAL NOT NULL,
    mass REAL NOT NULL,
    created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT (CURRENT_TIMESTAMP AT TIME ZONE 'UTC'),
    updated_at TIMESTAMP WITHOUT TIME ZONE DEFAULT (CURRENT_TIMESTAMP AT TIME ZONE 'UTC'),
    -- Внешний ключ manufacturer (one-to-many)
    manufacturer_id INT NOT NULL,
    -- Связь с таблицей manufacturers
    FOREIGN KEY (manufacturer_id) REFERENCES manufacturers(id)
);