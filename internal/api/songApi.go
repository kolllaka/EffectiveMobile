package api

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/kolllaka/EffectiveMobile/internal/model"
)

func (m *methods) GetSongInfo(ctx context.Context, songName model.AddSong) (model.Song, error) {
	value := url.Values{}
	value.Set("group", songName.Group)
	value.Set("song", songName.Song)
	apiUrl := url.URL{
		Scheme:   m.conf.SongApi.Scheme,
		Host:     m.conf.SongApi.Host,
		Path:     m.conf.SongApi.PathsInfo,
		RawQuery: value.Encode(),
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, apiUrl.String(), nil)
	if err != nil {
		return model.Song{}, err
	}
	client := http.DefaultClient
	res, err := client.Do(req)
	if err != nil {
		return model.Song{}, err
	}

	var song model.Song

	if err := json.NewDecoder(res.Body).Decode(&song); err != nil {
		return model.Song{}, err
	}

	return song, nil
}
