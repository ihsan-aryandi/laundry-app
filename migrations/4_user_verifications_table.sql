-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE IF NOT EXISTS user_verifications
(
    id                BIGSERIAL    NOT NULL,
    email             VARCHAR(256) NOT NULL,
    verification_code VARCHAR(6),
    expired_at        TIMESTAMP WITHOUT TIME ZONE,
    created_at        TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    created_by        BIGINT,
    updated_at        TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_by        BIGINT,
    is_deleted        BOOLEAN                     DEFAULT false,
    CONSTRAINT userverifications_id_pk PRIMARY KEY (id)
);

-- +migrate StatementEnd
-- +migrate Down

DROP TABLE IF EXISTS user_verifications;
