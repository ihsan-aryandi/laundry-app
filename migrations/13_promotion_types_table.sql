-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE IF NOT EXISTS promotion_types (
    id                   BIGSERIAL NOT NULL,
    name                 VARCHAR(256) NOT NULL,
    discount             INT NOT NULL DEFAULT 0,
    created_at           TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    created_by           BIGINT,
    updated_at           TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_by           BIGINT,
    is_deleted           BOOLEAN DEFAULT false,
    CONSTRAINT promotiontypes_id_pk PRIMARY KEY (id)
);

-- +migrate StatementEnd
-- +migrate Down

DROP TABLE IF EXISTS promotion_types;
