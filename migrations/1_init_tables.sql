-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS songs
(
	id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
	song_group VARCHAR(40) NOT NULL,
	song_name VARCHAR(40) NOT NULL,
	release_date VARCHAR(11) NOT NULL,
	song_text TEXT NOT NULL,
	song_link VARCHAR(200) NOT NULL,
	created_at TIMESTAMPTZ DEFAULT now()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS songs;
-- +goose StatementEnd
