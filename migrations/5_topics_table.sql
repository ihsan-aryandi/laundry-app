-- +migrate Up
-- +migrate StatementBegin

CREATE SEQUENCE topics_pkey_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

CREATE TABLE IF NOT EXISTS topics (
    id               bigint DEFAULT nextval('topics_pkey_seq'::regclass) NOT NULL,
    topic            VARCHAR(256) NOT NULL,
    created_at       timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    created_by       bigint,
    updated_at       timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    updated_by       bigint,
    is_deleted       boolean DEFAULT false,
    CONSTRAINT pk_topics_id PRIMARY KEY (id)
);

CREATE UNIQUE INDEX topics_uq
    ON topics(topic)
WHERE is_deleted = FALSE;

INSERT INTO topics
    (topic, created_at, created_by, updated_at, updated_by)
VALUES
    ('Bisnis', CURRENT_TIMESTAMP, 1, CURRENT_TIMESTAMP, 1);

-- +migrate StatementEnd
-- +migrate Down

DROP TABLE IF EXISTS topics;
DROP SEQUENCE IF EXISTS topics_pkey_seq;
