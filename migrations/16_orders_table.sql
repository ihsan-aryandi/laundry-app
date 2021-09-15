-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE IF NOT EXISTS orders
(
    id          BIGSERIAL   NOT NULL,
    customer_id BIGINT      NOT NULL,
    status      VARCHAR(50) NOT NULL,
    order_date  TIMESTAMP WITHOUT TIME ZONE,
    is_active   BOOLEAN                     DEFAULT true,
    total       INTEGER     NOT NULL,
    created_at  TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    created_by  BIGINT,
    updated_at  TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_by  BIGINT,
    is_deleted  BOOLEAN                     DEFAULT false,
    CONSTRAINT orders_id_pk PRIMARY KEY (id)
);

-- +migrate StatementEnd
-- +migrate Down

DROP TABLE IF EXISTS orders;
