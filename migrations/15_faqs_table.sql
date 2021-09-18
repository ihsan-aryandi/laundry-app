-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE IF NOT EXISTS faqs (
    id                   BIGSERIAL NOT NULL,
    question             VARCHAR(256) NOT NULL,
    answer               TEXT NOT NULL,
    created_at           TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    created_by           BIGINT,
    updated_at           TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_by           BIGINT,
    is_deleted           BOOLEAN DEFAULT false,
    CONSTRAINT faqs_id_pk PRIMARY KEY (id)
);

-- +migrate StatementEnd
-- +migrate Down

DROP TABLE IF EXISTS faqs;
