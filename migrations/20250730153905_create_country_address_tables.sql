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
	country_id BIGINT  NOT NULL,

	PRIMARY KEY ( id ),
	FOREIGN KEY ( country_id )
		REFERENCES countries ( id )
);

CREATE TABLE IF NOT EXISTS addresses
(
	id      BIGSERIAL,
	street  VARCHAR NOT NULL,
	zip     VARCHAR NOT NULL,
	city_id BIGINT  NOT NULL,

	PRIMARY KEY ( id ),
	FOREIGN KEY ( city_id )
		REFERENCES cities ( id )
);

ALTER TABLE users
	ADD COLUMN address_id BIGINT,
	ADD CONSTRAINT fk_address_id
		FOREIGN KEY ( address_id )
			REFERENCES addresses ( id );


-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE users
	DROP CONSTRAINT fk_address_id;
ALTER TABLE users
	DROP COLUMN address_id;

DROP TABLE IF EXISTS addresses;

DROP TABLE IF EXISTS cities;

DROP TABLE IF EXISTS countries;
-- +goose StatementEnd
