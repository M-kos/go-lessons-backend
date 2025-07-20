-- +goose Up
-- +goose StatementBegin

CREATE TABLE IF NOT EXISTS users
(
    id          BIGSERIAL,
    email       VARCHAR     NOT NULL,
    full_name   VARCHAR     NOT NULL,
    create_time TIMESTAMPTZ NOT NULL DEFAULT now(),

    PRIMARY KEY (id)
);

SELECT setval('users_id_seq', 200_000);

INSERT INTO users (email, full_name)
VALUES ('user_1@example.com', 'Ivan Ivanov'),
       ('user_2@example.com', 'Bill Gates'),
       ('user_3@example.com', 'John Depp');

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE IF EXISTS users;

-- +goose StatementEnd
