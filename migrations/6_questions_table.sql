-- +migrate Up
-- +migrate StatementBegin

CREATE SEQUENCE questions_pkey_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

CREATE TABLE IF NOT EXISTS questions (
    id               bigint DEFAULT nextval('questions_pkey_seq'::regclass) NOT NULL,
    user_profile_id  bigint NOT NULL,
    topic_id         bigint,
    question         VARCHAR(256),
    body             text,
    vote_count       int default 0,
    created_at       timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    created_by       bigint,
    updated_at       timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    updated_by       bigint,
    is_deleted       boolean DEFAULT false,
    CONSTRAINT pk_questions_id PRIMARY KEY (id)
);

-- +migrate StatementEnd
-- +migrate Down

DROP TABLE IF EXISTS questions;
DROP SEQUENCE IF EXISTS questions_pkey_seq;
