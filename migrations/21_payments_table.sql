-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE IF NOT EXISTS payments (
    id                    BIGSERIAL NOT NULL,
    order_id              BIGINT NOT NULL,
    total                 BIGINT NOT NULL,
    is_paid               BOOL NOT NULL DEFAULT false;
    paid_at               TIMESTAMP WITHOUT TIME ZONE NULL,
    created_at            TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    created_by            BIGINT,
    updated_at            TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_by            BIGINT,
    is_deleted            BOOLEAN DEFAULT false,
    CONSTRAINT payments_id_pk PRIMARY KEY (id)
);

-- +migrate StatementEnd
-- +migrate Down

DROP TABLE IF EXISTS payments;
