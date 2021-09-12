-- +migrate Up
-- +migrate StatementBegin

CREATE SEQUENCE roles_pkey_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

CREATE TABLE IF NOT EXISTS roles (
    id               bigint DEFAULT nextval('roles_pkey_seq'::regclass) NOT NULL,
    role             VARCHAR(256) NOT NULL,
    created_at       timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    created_by       bigint,
    updated_at       timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    updated_by       bigint,
    is_deleted       boolean DEFAULT false,
    CONSTRAINT pk_roles_id PRIMARY KEY (id)
);

CREATE UNIQUE INDEX roles_uq
    ON roles(role)
WHERE is_deleted = FALSE;

INSERT INTO roles
    (role, created_at, created_by, updated_at, updated_by)
VALUES
    ('admin', CURRENT_TIMESTAMP, 1, CURRENT_TIMESTAMP, 1);

-- +migrate StatementEnd
-- +migrate Down

DROP TABLE IF EXISTS roles;
DROP SEQUENCE IF EXISTS roles_pkey_seq;
