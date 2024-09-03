CREATE TABLE data (
    id                  UUID NOT NULL PRIMARY KEY DEFAULT gen_random_uuid(),
    owner               UUID NOT NULL,
    title               VARCHAR(255) NOT NULL,
    encrypted_payload   TEXT NOT NULL,
    part                SMALLINT NULL,
    created_at          TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at          TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX ON data (owner);