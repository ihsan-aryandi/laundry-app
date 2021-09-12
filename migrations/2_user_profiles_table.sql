-- +migrate Up
-- +migrate StatementBegin

CREATE SEQUENCE user_profiles_pkey_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

CREATE TABLE IF NOT EXISTS user_profiles (
    id               bigint DEFAULT nextval('user_profiles_pkey_seq'::regclass) NOT NULL,
    name             VARCHAR(256),
    gender_id        bigint NOT NULL,
    address          text,
    photo            VARCHAR(256),
    bio              text,
    total_reputation bigint default 1,
    role_id          bigint NOT NULL,
    created_at       timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    created_by       bigint,
    updated_at       timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    updated_by       bigint,
    is_deleted       boolean DEFAULT false,
    CONSTRAINT pk_user_profiles_id PRIMARY KEY (id)
);

-- +migrate StatementEnd
-- +migrate Down

DROP TABLE IF EXISTS user_profiles;
DROP SEQUENCE IF EXISTS user_profiles_pkey_seq;
