CREATE TABLE data_user_pass (
    id          UUID NOT NULL PRIMARY KEY DEFAULT gen_random_uuid(),
    owner       UUID NOT NULL,
    login       VARCHAR(255) NOT NULL,
    password    VARCHAR(255) NOT NULL,
    title       VARCHAR(255) NOT NULL,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX ON data_user_pass(owner);