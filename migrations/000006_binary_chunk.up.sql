CREATE TABLE binary_data_chunk (
    binary_data_id          INT NOT NULL,
    encrypted_chunk         TEXT NOT NULL,
    order_number            INT NOT NULL,
    created_at              TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX ON binary_data_chunk (binary_data_id);