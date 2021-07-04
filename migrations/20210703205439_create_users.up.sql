CREATE TABLE users (
    id bigserial not null primary key,
    telegram_id integer not null
);

CREATE INDEX users_telegram_id 
  ON users (telegram_id);