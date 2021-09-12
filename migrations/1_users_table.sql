-- +migrate Up
-- +migrate StatementBegin

CREATE SEQUENCE users_pkey_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

CREATE TABLE IF NOT EXISTS users (
    id bigint DEFAULT nextval('users_pkey_seq'::regclass) NOT NULL,
    username VARCHAR(256),
    password VARCHAR(256),
    user_profile_id bigint NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    created_by bigint,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    updated_by bigint,
    is_deleted boolean DEFAULT false,
    CONSTRAINT pk_users_id PRIMARY KEY (id)
);

CREATE UNIQUE INDEX users_uq
    ON users(username, user_profile_id)
WHERE is_deleted = FALSE;

-- +migrate StatementEnd
-- +migrate Down

DROP TABLE IF EXISTS users;
DROP SEQUENCE IF EXISTS users_pkey_seq;
