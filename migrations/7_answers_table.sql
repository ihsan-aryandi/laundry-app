-- +migrate Up
-- +migrate StatementBegin

CREATE SEQUENCE answers_pkey_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

CREATE TABLE IF NOT EXISTS answers (
    id               bigint DEFAULT nextval('answers_pkey_seq'::regclass) NOT NULL,
    user_profile_id  bigint NOT NULL,
    question_id      bigint NOT NULL,
    answer           text NOT NULL,
    vote_count       int default 0,
    is_accepted      boolean default false,
    created_at       timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    created_by       bigint,
    updated_at       timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    updated_by       bigint,
    is_deleted       boolean DEFAULT false,
    CONSTRAINT pk_answers_id PRIMARY KEY (id)
);

-- +migrate StatementEnd
-- +migrate Down

DROP TABLE IF EXISTS answers;
DROP SEQUENCE IF EXISTS answers_pkey_seq;
