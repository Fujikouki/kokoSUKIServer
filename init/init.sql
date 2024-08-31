CREATE TABLE accounts
(
    id           SERIAL PRIMARY KEY,
    email        VARCHAR(255) NOT NULL UNIQUE,
    username     VARCHAR(50) NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    icon_url     VARCHAR(255),
    created_at   TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE chat_rooms
(
    id           SERIAL PRIMARY KEY,
    name         VARCHAR(100) NOT NULL,
    created_at   TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE chat_messages
(
    id           SERIAL PRIMARY KEY,
    account_id   INTEGER REFERENCES accounts(id),
    room_id      INTEGER REFERENCES chat_rooms(id),
    message      TEXT NOT NULL,
    time         TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE room_memberships
(
    id           SERIAL PRIMARY KEY,
    account_id   INTEGER REFERENCES accounts(id),
    room_id      INTEGER REFERENCES chat_rooms(id),
    joined_at    TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

-- Optional: ファイル共有のためのテーブル
CREATE TABLE files
(
    id           SERIAL PRIMARY KEY,
    message_id   INTEGER REFERENCES chat_messages(id),
    file_url     VARCHAR(255) NOT NULL,
    uploaded_at  TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);
