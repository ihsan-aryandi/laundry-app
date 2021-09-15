-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE IF NOT EXISTS couriers (
    id                    BIGSERIAL NOT NULL,
    laundry_id            BIGINT NOT NULL,
    user_profile_id       BIGINT NOT NULL,
    is_active             BOOLEAN NOT NULL DEFAULT true,
    created_at            TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    created_by            BIGINT,
    updated_at            TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_by            BIGINT,
    is_deleted            BOOLEAN DEFAULT false,
    CONSTRAINT couriers_id_pk PRIMARY KEY (id)
);

-- +migrate StatementEnd
-- +migrate Down

DROP TABLE IF EXISTS couriers;
