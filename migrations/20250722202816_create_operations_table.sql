-- +goose Up
-- +goose StatementBegin

CREATE TYPE operation_type AS ENUM ('debit', 'credit');

CREATE TABLE IF NOT EXISTS operations
(
    id              UUID,
    user_id         BIGINT         NOT NULL,
    sequence_number BIGINT         NOT NULL,
    operation_type  operation_type NOT NULL,
    amount          DECIMAL(15, 2) NOT NULL,
    create_time     TIMESTAMPTZ    NOT NULL DEFAULT now(),

    PRIMARY KEY (id)
);

CREATE UNIQUE INDEX ON operations (user_id, sequence_number);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE IF EXISTS operations;

DROP TYPE operation_type;

-- +goose StatementEnd
