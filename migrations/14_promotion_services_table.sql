-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE IF NOT EXISTS promotion_services
(
    id                  BIGSERIAL NOT NULL,
    promotion_id        BIGINT    NOT NULL,
    service_id          BIGINT    NOT NULL,
    service_category_id BIGINT    NOT NULL,
    created_at          TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    created_by          BIGINT,
    updated_at          TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_by          BIGINT,
    is_deleted          BOOLEAN                     DEFAULT false,
    CONSTRAINT promotionservices_id_pk PRIMARY KEY (id)
);

-- +migrate StatementEnd
-- +migrate Down

DROP TABLE IF EXISTS promotion_services;
