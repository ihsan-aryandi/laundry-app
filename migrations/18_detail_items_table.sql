-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE IF NOT EXISTS detail_items
(
    id            BIGSERIAL    NOT NULL,
    order_item_id BIGINT       NOT NULL,
    name          VARCHAR(256) NOT NULL,
    photo         VARCHAR(256),
    created_at    TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    created_by    BIGINT,
    updated_at    TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_by    BIGINT,
    is_deleted    BOOLEAN                     DEFAULT false,
    CONSTRAINT detailitems_id_pk PRIMARY KEY (id)
);

-- +migrate StatementEnd
-- +migrate Down

DROP TABLE IF EXISTS detail_items;
