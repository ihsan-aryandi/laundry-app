-- +migrate Up
-- +migrate StatementBegin

CREATE SEQUENCE answer_comments_pkey_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

CREATE TABLE IF NOT EXISTS answer_comments (
    id               bigint DEFAULT nextval('answer_comments_pkey_seq'::regclass) NOT NULL,
    user_profile_id  bigint NOT NULL,
    answer_id        bigint NOT NULL,
    comment          text NOT NULL,
    created_at       timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    created_by       bigint,
    updated_at       timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    updated_by       bigint,
    is_deleted       boolean DEFAULT false,
    CONSTRAINT pk_answer_comments_id PRIMARY KEY (id)
);

-- +migrate StatementEnd
-- +migrate Down

DROP TABLE IF EXISTS answer_comments;
DROP SEQUENCE IF EXISTS answer_comments_pkey_seq;
