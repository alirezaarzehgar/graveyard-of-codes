package controller

import (
	"net/http"
	"strconv"

	"github.com/alirezaarzehgar/hi/model"
	"github.com/labstack/echo/v4"
)

func CreatePost(c echo.Context) error {
	title := c.FormValue("title")
	content := c.FormValue("content")
	user_id, _ := strconv.ParseInt(c.FormValue("user_id"), 10, 32)

	post := model.Post{
		Title:   title,
		Content: content,
		UserID:  int(user_id),
	}

	err := db.Create(&post).Error
	if err != nil {
		return echo.ErrBadRequest
	}

	return c.JSON(http.StatusOK, map[string]any{
		"post_id": post.ID,
	})
}

func GetPost(c echo.Context) error {
	post_id := c.Param("id")
	var post map[string]any

	r := db.Model(&model.Post{}).Find(&post, post_id)
	if r.Error != nil {
		return echo.ErrInternalServerError
	}
	if r.RowsAffected == 0 {
		return echo.ErrNotFound
	}

	return c.JSON(http.StatusOK, post)
}

func EditPost(c echo.Context) error {
	post_id, _ := strconv.ParseInt(c.Param("id"), 10, 32)
	title := c.FormValue("title")
	content := c.FormValue("content")

	err := db.Save(&model.Post{
		ID:      int(post_id),
		Title:   title,
		Content: content,
	}).Error
	if err != nil {
		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, map[string]any{
		"user_id": post_id,
	})
}

func DeletePost(c echo.Context) error {
	post_id, _ := strconv.ParseInt(c.Param("id"), 10, 32)

	r := db.Delete(&model.Post{}, post_id)
	if r.Error != nil {
		return echo.ErrInternalServerError
	}
	if r.RowsAffected == 0 {
		return echo.ErrNotFound
	}

	return c.JSON(http.StatusOK, map[string]any{
		"user_id": post_id,
	})
}
