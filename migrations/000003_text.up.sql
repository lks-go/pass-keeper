CREATE TABLE text (
    id                  SERIAL NOT NULL PRIMARY KEY,
    owner               UUID NOT NULL,
    title               VARCHAR(255) NOT NULL,
    encrypted_text      TEXT NOT NULL,
    created_at          TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at          TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX ON text (owner);