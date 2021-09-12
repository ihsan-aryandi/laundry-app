-- +migrate Up
-- +migrate StatementBegin

CREATE SEQUENCE info_posts_pkey_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

CREATE TABLE IF NOT EXISTS info_posts (
    id               bigint DEFAULT nextval('info_posts_pkey_seq'::regclass) NOT NULL,
    user_profile_id  bigint NOT NULL,
    title            VARCHAR(256),
    body             text NOT NULL,
    vote_count       int default 0,
    created_at       timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    created_by       bigint,
    updated_at       timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    updated_by       bigint,
    is_deleted       boolean DEFAULT false,
    CONSTRAINT pk_info_posts_id PRIMARY KEY (id)
);

-- +migrate StatementEnd
-- +migrate Down

DROP TABLE IF EXISTS info_posts;
DROP SEQUENCE IF EXISTS info_posts_pkey_seq;
