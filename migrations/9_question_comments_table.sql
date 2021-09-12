-- +migrate Up
-- +migrate StatementBegin

CREATE SEQUENCE question_comments_pkey_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

CREATE TABLE IF NOT EXISTS question_comments (
    id               bigint DEFAULT nextval('question_comments_pkey_seq'::regclass) NOT NULL,
    user_profile_id  bigint NOT NULL,
    question_id      bigint NOT NULL,
    comment          text NOT NULL,
    created_at       timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    created_by       bigint,
    updated_at       timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    updated_by       bigint,
    is_deleted       boolean DEFAULT false,
    CONSTRAINT pk_question_comments_id PRIMARY KEY (id)
);

-- +migrate StatementEnd
-- +migrate Down

DROP TABLE IF EXISTS question_comments;
DROP SEQUENCE IF EXISTS question_comments_pkey_seq;
