-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE IF NOT EXISTS services
(
    id          BIGSERIAL    NOT NULL,
    name        VARCHAR(256) NOT NULL,
    photo       VARCHAR(256),
    description TEXT,
    created_at  TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    created_by  BIGINT,
    updated_at  TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_by  BIGINT,
    is_deleted  BOOLEAN                     DEFAULT false,
    CONSTRAINT services_id_pk PRIMARY KEY (id)
);

-- +migrate StatementEnd
-- +migrate Down

DROP TABLE IF EXISTS services;
