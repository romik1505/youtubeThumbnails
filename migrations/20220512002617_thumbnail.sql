-- +goose Up
-- +goose StatementBegin

CREATE TABLE thumbnails (
    id TEXT PRIMARY KEY,
    thumbnail BLOB,
    created_at DATETIME DEFAULT (datetime('now'))
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE thumbnails;

-- +goose StatementEnd
