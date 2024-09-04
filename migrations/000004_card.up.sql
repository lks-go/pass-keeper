CREATE TABLE card (
    id                      SERIAL NOT NULL PRIMARY KEY,
    owner                   UUID NOT NULL,
    title                   VARCHAR(255) NOT NULL,
    encrypted_number        TEXT NOT NULL,
    encrypted_owner         TEXT NOT NULL,
    encrypted_exp_date      TEXT NOT NULL,
    encrypted_cvc_code      TEXT NOT NULL,
    created_at              TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at              TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX ON login_pass (owner);