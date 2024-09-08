CREATE TABLE binary_data (
    id                      SERIAL NOT NULL PRIMARY KEY,
    owner                   UUID NOT NULL,
    title                   VARCHAR(255) NULL,
    created_at              TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at              TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX ON binary_data (owner);