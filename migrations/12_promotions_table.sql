-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE IF NOT EXISTS promotions
(
    id             BIGSERIAL    NOT NULL,
    name           VARCHAR(256) NOT NULL,
    photo          VARCHAR(256),
    start_date     TIMESTAMP WITHOUT TIME ZONE,
    end_date       TIMESTAMP WITHOUT TIME ZONE,
    amount         INTEGER      NOT NULL       DEFAULT 0,
    limit_per_user INTEGER      NOT NULL       DEFAULT 1,
    description    TEXT,
    code           VARCHAR(50),
    --condition
    created_at     TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    created_by     BIGINT,
    updated_at     TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_by     BIGINT,
    is_deleted     BOOLEAN                     DEFAULT false,
    CONSTRAINT promotions_id_pk PRIMARY KEY (id)
);

-- +migrate StatementEnd
-- +migrate Down

DROP TABLE IF EXISTS promotions;
