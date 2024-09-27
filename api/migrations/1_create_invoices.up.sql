CREATE TABLE invoices (
    id BIGSERIAL PRIMARY KEY,
    data BYTEA NOT NULL,
    created_at TIMESTAMP
);
