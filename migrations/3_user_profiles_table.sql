-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE IF NOT EXISTS user_profiles
(
    id         BIGSERIAL NOT NULL,
    name       VARCHAR(256),
    phone      VARCHAR(25),
    sex        VARCHAR(25),
    address    TEXT,
    photo      VARCHAR(256),
    role_id    BIGINT    NOT NULL,
    created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    created_by BIGINT,
    updated_at TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_by BIGINT,
    is_deleted BOOLEAN                     DEFAULT false,
    CONSTRAINT userprofiles_id_pk PRIMARY KEY (id)
);

-- +migrate StatementEnd
-- +migrate Down

DROP TABLE IF EXISTS user_profiles;
