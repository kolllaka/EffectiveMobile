package service

import (
	"context"

	"github.com/kolllaka/EffectiveMobile/internal/model"
)

type MusicService interface {
	GetSongs(ctx context.Context, param model.QueryParam) ([]model.Song, error)
	CreateSong(ctx context.Context, newSong model.AddSong) (string, error)
	ChangeSong(ctx context.Context, song model.Song) error
	DeleteSong(ctx context.Context, id string) error
	GetTextOfSong(ctx context.Context, id string) (string, error)
}

func (s *service) GetSongs(ctx context.Context, param model.QueryParam) ([]model.Song, error) {
	return s.store.GetSongs(ctx, param)
}
func (s *service) CreateSong(ctx context.Context, newSong model.AddSong) (string, error) {
	song, err := s.api.GetSongInfo(ctx, newSong)
	if err != nil {
		return "", err
	}

	return s.store.CreateSong(ctx, song)
}
func (s *service) ChangeSong(ctx context.Context, song model.Song) error {
	return s.store.ChangeSong(ctx, song)
}
func (s *service) DeleteSong(ctx context.Context, id string) error {
	return s.store.DeleteSong(ctx, id)
}
func (s *service) GetTextOfSong(ctx context.Context, id string) (string, error) {
	return s.store.GetTextOfSong(ctx, id)
}
