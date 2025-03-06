-- +goose Up
-- +goose StatementBegin
CREATE TABLE groups (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

CREATE UNIQUE INDEX idx_groups_name ON groups(name);

-- +goose Down
-- +goose StatementBegin
DROP TABLE groups;
-- +goose StatementEnd
