-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE IF NOT EXISTS order_items (
    id                    BIGSERIAL NOT NULL,
    order_id              BIGINT NOT NULL,
    service_category_id   BIGINT NOT NULL,
    status                VARCHAR(50) NOT NULL,
    created_at            TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    created_by            BIGINT,
    updated_at            TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_by            BIGINT,
    is_deleted            BOOLEAN DEFAULT false,
    CONSTRAINT orderitems_id_pk PRIMARY KEY (id)
);

-- +migrate StatementEnd
-- +migrate Down

DROP TABLE IF EXISTS order_items;
