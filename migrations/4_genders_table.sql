-- +migrate Up
-- +migrate StatementBegin

CREATE SEQUENCE genders_pkey_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

CREATE TABLE IF NOT EXISTS genders (
    id               bigint DEFAULT nextval('genders_pkey_seq'::regclass) NOT NULL,
    gender           VARCHAR(256) NOT NULL,
    created_at       timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    created_by       bigint,
    updated_at       timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    updated_by       bigint,
    is_deleted       boolean DEFAULT false,
    CONSTRAINT pk_genders_id PRIMARY KEY (id)
);

CREATE UNIQUE INDEX genders_uq
    ON genders(gender)
WHERE is_deleted = FALSE;

INSERT INTO genders
    (gender, created_at, created_by, updated_at, updated_by)
VALUES
    ('Laki-laki', CURRENT_TIMESTAMP, 1, CURRENT_TIMESTAMP, 1),
    ('Perempuan', CURRENT_TIMESTAMP, 1, CURRENT_TIMESTAMP, 1);

-- +migrate StatementEnd
-- +migrate Down

DROP TABLE IF EXISTS genders;
DROP SEQUENCE IF EXISTS genders_pkey_seq;
