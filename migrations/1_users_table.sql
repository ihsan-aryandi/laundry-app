-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE IF NOT EXISTS users
(
    id              BIGSERIAL NOT NULL,
    email           VARCHAR(256),
    password        VARCHAR(256),
    user_profile_id BIGINT    NOT NULL,
    created_at      TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    created_by      BIGINT,
    updated_at      TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_by      BIGINT,
    is_deleted      BOOLEAN                     DEFAULT false,
    CONSTRAINT users_id_pk PRIMARY KEY (id)
);

CREATE UNIQUE INDEX users_email_uq
    ON users(email, user_profile_id)
    WHERE is_deleted = FALSE;

-- +migrate StatementEnd
-- +migrate Down

DROP TABLE IF EXISTS users;
