-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE IF NOT EXISTS roles (
    id               BIGSERIAL NOT NULL,
    role             VARCHAR(256) NOT NULL,
    created_at       TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    created_by       BIGINT,
    updated_at       TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_by       BIGINT,
    is_deleted       BOOLEAN DEFAULT false,
    CONSTRAINT roles_id_pk PRIMARY KEY (id)
);

-- +migrate StatementEnd
-- +migrate Down

DROP TABLE IF EXISTS roles;
