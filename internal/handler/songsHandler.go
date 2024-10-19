package handler

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/kolllaka/EffectiveMobile/internal/model"
	"github.com/kolllaka/EffectiveMobile/internal/utils"
)

// @Summary		Получить все песни
// @Tags			API для работы с песнями
// @Description	Получить все песни
// @ID				api-songs-get
// @Produce		json
// @Param			field	query		string			false	"sort field: group, song, releaseDate, text or link"
// @Param			sort	query		string			false	"ASC or DESC"
// @Param			page	query		int				false	"count songs on page"
// @Param			limit	query		int				false	"page number"
// @Success		200		{object}	[]model.Song	"ok"
// @Failure		400,404	{object}	errorResponse
// @Failure		500		{object}	errorResponse
// @Failure		default	{object}	errorResponse
// @Router			/api/songs [get]
func (s *server) getSongsHandler(c *gin.Context) {
	param := model.QueryParam{
		Field: strings.ToLower(c.Query("field")),
		Sort:  strings.ToUpper(c.Query("sort")),
		Page:  utils.ATOIwithDeafult(c.Query("page"), 1),
		Limit: utils.ATOIwithDeafult(c.Query("limit"), s.conf.Song.SongOnPage),
	}

	songs, err := s.service.GetSongs(c, param)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "Internal server error")

		return
	}

	c.JSON(http.StatusOK, songs)
}

// @Summary		Создать Песню
// @Tags			API для работы с песнями
// @Description	Создать Песню
// @ID				api-songs-post
// @Accept			json
// @Produce		json
// @Param			newSong	body		model.AddSong	true	"название песни(song) и группы(group)"
// @Success		200		{object}	statusResponse	"ok"
// @Failure		400,404	{object}	errorResponse
// @Failure		500		{object}	errorResponse
// @Failure		default	{object}	errorResponse
// @Router			/api/songs [post]
func (s *server) createSongHandler(c *gin.Context) {
	var song model.AddSong
	if err := c.BindJSON(&song); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Bad request")

		return
	}

	id, err := s.service.CreateSong(c, song)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "Internal server error")

		return
	}

	c.JSON(http.StatusCreated, statusResponse{id, "created"})

}

// @Summary		Изменить Песню по id
// @Tags			API для работы с песнями
// @Description	Изменить Песню по id
// @ID				api-songs-id-put
// @Accept			json
// @Produce		json
// @Param			id		path		string			false	"song id"
// @Param			newSong	body		model.Song		true	"новые поля песни"
// @Success		200		{object}	statusResponse	"ok"
// @Failure		400,404	{object}	errorResponse
// @Failure		500		{object}	errorResponse
// @Failure		default	{object}	errorResponse
// @Router			/api/songs/{id} [put]
func (s *server) changeSongHandler(c *gin.Context) {
	var song model.Song

	id := c.Param("id")
	if err := c.BindJSON(&song); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Bad request")

		return
	}
	song.Id = id

	if err := s.service.ChangeSong(c, song); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "Internal server error")

		return
	}

	c.JSON(http.StatusCreated, statusResponse{id, "changed"})
}

// @Summary		Удалить Песню по id
// @Tags			API для работы с песнями
// @Description	Удалить Песню по id
// @ID				api-songs-id-delete
// @Accept			json
// @Produce		json
// @Param			id		path		string			false	"song id"
// @Success		200		{object}	statusResponse	"ok"
// @Failure		400,404	{object}	errorResponse
// @Failure		500		{object}	errorResponse
// @Failure		default	{object}	errorResponse
// @Router			/api/songs/{id} [delete]
func (s *server) deleteSongHandler(c *gin.Context) {
	id := c.Param("id")
	if err := s.service.DeleteSong(c, id); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "Internal server error")

		return
	}
	c.JSON(http.StatusOK, statusResponse{id, "deleted"})
}

// @Summary		Получить текст песни Песню по id
// @Tags			API для работы с песнями
// @Description	Получить текст песни Песню по id
// @ID				api-songs-id-get
// @Accept			json
// @Produce		json
// @Param			id		path		string									false	"song id"
// @Success		200		{object}	handler.getTextOfSongHandler.responce	"ok"
// @Failure		400,404	{object}	errorResponse
// @Failure		500		{object}	errorResponse
// @Failure		default	{object}	errorResponse
// @Router			/api/songs/{id} [get]
func (s *server) getTextOfSongHandler(c *gin.Context) {
	type Page struct {
		PageCount int    `json:"page_count,omitempty"`
		Text      string `json:"text,omitempty"`
	} //	@name	Page
	type responce struct {
		Pages []Page `json:"pages,omitempty"`
	}
	var pages []Page

	id := c.Param("id")
	textOfSongs, err := s.service.GetTextOfSong(c, id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "Internal server error")

		return
	}

	songCouplets := strings.Split(textOfSongs, "\n\n")
	for i, couplet := range songCouplets {
		pages = append(pages, Page{
			PageCount: i + 1,
			Text:      couplet,
		})
	}

	c.JSON(http.StatusOK, responce{Pages: pages})
}
