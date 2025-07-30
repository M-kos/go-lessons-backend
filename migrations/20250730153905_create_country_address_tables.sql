-- +goose Up
-- +goose StatementBegin

CREATE TABLE IF NOT EXISTS countries
(
	id   BIGSERIAL,
	name VARCHAR NOT NULL,

	PRIMARY KEY ( id )
);

CREATE TABLE IF NOT EXISTS cities
(
	id         BIGSERIAL,
	name       VARCHAR NOT NULL,
	country_id BIGINT,

	PRIMARY KEY ( id ),
	FOREIGN KEY ( country_id )
		REFERENCES countries ( id )
);

CREATE TABLE IF NOT EXISTS addresses
(
	id      BIGSERIAL,
	street  VARCHAR NOT NULL,
	zip     VARCHAR NOT NULL,
	city_id BIGINT,

	PRIMARY KEY ( id ),
	FOREIGN KEY ( city_id )
		REFERENCES cities ( id )
);

CREATE TABLE IF NOT EXISTS user_address
(
	id         BIGSERIAL,
	user_id    BIGINT,
	address_id BIGINT,

	PRIMARY KEY ( id ),
	FOREIGN KEY ( address_id )
		REFERENCES addresses ( id ),
	FOREIGN KEY ( user_id )
		REFERENCES users ( id )
);

ALTER TABLE users
	ADD COLUMN user_address_id BIGINT,
	ADD CONSTRAINT fk_user_address_id
		FOREIGN KEY ( user_address_id )
			REFERENCES user_address ( id );


-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS countries;
DROP TABLE IF EXISTS cities;
DROP TABLE IF EXISTS addresses;
DROP TABLE IF EXISTS user_address;
-- +goose StatementEnd
