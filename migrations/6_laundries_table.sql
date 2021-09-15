-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE IF NOT EXISTS laundries
(
    id              BIGSERIAL NOT NULL,
    user_profile_id bigint    NOT NULL,
    name            VARCHAR(256),
    about_us        TEXT,
    photo           VARCHAR(256),
    address         TEXT,
    created_at      TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    created_by      BIGINT,
    updated_at      TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_by      BIGINT,
    is_deleted      BOOLEAN                     DEFAULT false,
    CONSTRAINT laundries_id_pk PRIMARY KEY (id)
);

-- +migrate StatementEnd
-- +migrate Down

DROP TABLE IF EXISTS laundries;
