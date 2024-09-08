CREATE TABLE login_pass (
    id                  SERIAL NOT NULL PRIMARY KEY,
    owner               UUID NOT NULL,
    title               VARCHAR(255) NOT NULL,
    encrypted_login     TEXT NOT NULL,
    encrypted_password  TEXT NOT NULL,
    created_at          TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at          TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX ON login_pass (owner);