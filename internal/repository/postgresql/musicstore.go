package db

import (
	"context"
	"fmt"

	"github.com/kolllaka/EffectiveMobile/internal/model"
)

const (
	songsTable = "songs"
)

func (s *store) GetSongs(ctx context.Context, param model.QueryParam) ([]model.Song, error) {
	var songs []model.Song

	stmt := fmt.Sprintf("SELECT id, song_group, song_name, release_date, song_text, song_link FROM %s", songsTable)
	if param.Field != "" {
		var field string

		switch param.Field {
		case "group":
			field = "song_group"
		case "song":
			field = "song_name"
		case "releaseDate":
			field = "release_date"
		case "text":
			field = "song_text"
		case "link":
			field = "song_link"
		default:
			goto NEXT
		}

		if param.Sort == "DESC" {
			stmt = fmt.Sprintf("%s ORDER BY %s DESC", stmt, field)
		} else {
			stmt = fmt.Sprintf("%s ORDER BY %s ASC", stmt, field)
		}
	}
NEXT:

	rows, err := s.db.Query(ctx, stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var song model.Song
		if err := rows.Scan(&song.Id, &song.Group, &song.Song, &song.ReleaseDate, &song.Text, &song.Link); err != nil {
			return nil, err
		}

		songs = append(songs, song)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return songs, nil
}
func (s *store) CreateSong(ctx context.Context, newSong model.Song) (string, error) {
	var id string

	stmt := fmt.Sprintf("INSERT INTO %s (song_group, song_name, release_date, song_text, song_link) VALUES ($1, $2, $3, $4, $5) RETURNING id", songsTable)
	if err := s.db.QueryRow(ctx, stmt, newSong.Group, newSong.Song, newSong.ReleaseDate, newSong.Text, newSong.Link).Scan(&id); err != nil {
		return "", err
	}

	return id, nil
}
func (s *store) ChangeSong(ctx context.Context, song model.Song) error {
	stmt := fmt.Sprintf("UPDATE %s SET song_group = $2, song_name= $3, release_date= $4, song_text= $5, song_link= $6 WHERE id = $1", songsTable)
	res, err := s.db.Exec(ctx, stmt,
		song.Id,
		song.Group,
		song.Song,
		song.ReleaseDate,
		song.Text,
		song.Link,
	)
	if err != nil {
		return err
	}

	rowsAffected := res.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("song with id %s not found", song.Id)
	}

	return nil
}
func (s *store) DeleteSong(ctx context.Context, id string) error {
	stmt := fmt.Sprintf("DELETE FROM %s WHERE id = $1", songsTable)
	res, err := s.db.Exec(ctx, stmt, id)
	if err != nil {
		return err
	}

	rowsAffected := res.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("song with id %s not found", id)
	}

	return nil
}
func (s *store) GetTextOfSong(ctx context.Context, id string) (string, error) {
	var songText string

	stmt := fmt.Sprintf("SELECT song_text FROM %s WHERE id = $1", songsTable)
	if err := s.db.QueryRow(ctx, stmt, id).Scan(&songText); err != nil {
		return "", err
	}

	return songText, nil
}
