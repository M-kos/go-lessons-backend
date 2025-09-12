-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS user_balance_snapshots
(
	id              BIGSERIAL,
	user_id         BIGINT         NOT NULL,
	sequence_number BIGINT         NOT NULL,
	balance         DECIMAL(15, 2) NOT NULL,
	update_time     TIMESTAMPTZ    NOT NULL DEFAULT now( ),

	PRIMARY KEY ( id )
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS user_balance_snapshots;
-- +goose StatementEnd
