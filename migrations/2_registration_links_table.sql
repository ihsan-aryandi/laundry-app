-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE IF NOT EXISTS registration_links
(
    id         BIGSERIAL NOT NULL,
    code       VARCHAR(256),
    expired_at TIMESTAMP WITHOUT TIME ZONE,
    type       VARCHAR(50),
    laundry_id BIGINT    NOT NULL,
    created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    created_by BIGINT,
    updated_at TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_by BIGINT,
    is_deleted BOOLEAN                     DEFAULT false,
    CONSTRAINT registrationlinks_id_pk PRIMARY KEY (id)
);

-- +migrate StatementEnd
-- +migrate Down

DROP TABLE IF EXISTS registration_links;
