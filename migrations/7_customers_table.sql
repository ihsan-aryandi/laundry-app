-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE IF NOT EXISTS customers (
    id               BIGSERIAL NOT NULL,
    laundry_id       BIGINT NOT NULL,
    user_profile_id  BIGINT NOT NULL,
    created_at       TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    created_by       BIGINT,
    updated_at       TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_by       BIGINT,
    is_deleted       BOOLEAN DEFAULT false,
    CONSTRAINT customers_id_pk PRIMARY KEY (id)
);

-- +migrate StatementEnd
-- +migrate Down

DROP TABLE IF EXISTS customers;
