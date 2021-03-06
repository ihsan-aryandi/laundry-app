-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE IF NOT EXISTS service_steps (
    id                   BIGSERIAL NOT NULL,
    service_category_id  BIGINT NOT NULL,
    step                 VARCHAR(256) NOT NULL,
    estimation           INT DEFAULT 0,
    created_at           TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    created_by           BIGINT,
    updated_at           TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_by           BIGINT,
    is_deleted           BOOLEAN DEFAULT false,
    CONSTRAINT servicesteps_id_pk PRIMARY KEY (id)
);

-- +migrate StatementEnd
-- +migrate Down

DROP TABLE IF EXISTS service_steps;
