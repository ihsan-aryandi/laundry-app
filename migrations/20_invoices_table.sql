-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE IF NOT EXISTS invoices
(
    id             BIGSERIAL    NOT NULL,
    order_id       BIGINT       NOT NULL,
    invoice_number VARCHAR(256) NOT NULL,
    created_at     TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    created_by     BIGINT,
    updated_at     TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_by     BIGINT,
    is_deleted     BOOLEAN                     DEFAULT false,
    CONSTRAINT invoices_id_pk PRIMARY KEY (id)
);

-- +migrate StatementEnd
-- +migrate Down

DROP TABLE IF EXISTS invoices;
