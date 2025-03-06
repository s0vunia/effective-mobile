-- +goose Up
-- +goose StatementBegin
CREATE TABLE verses (
    id SERIAL PRIMARY KEY,
    song_id INT NOT NULL,
    verse_number INT NOT NULL,
    text TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (song_id) REFERENCES songs(id) ON DELETE CASCADE
);
-- +goose StatementEnd

CREATE INDEX idx_verses_song_id ON verses(song_id);
CREATE INDEX idx_verses_verse_number ON verses(verse_number);
CREATE INDEX idx_verses_text ON verses(text);

-- +goose Down
-- +goose StatementBegin
DROP TABLE verses;
-- +goose StatementEnd
