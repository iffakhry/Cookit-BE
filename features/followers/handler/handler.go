package handler

import (
	"alta-cookit-be/features/followers"
	"alta-cookit-be/middlewares"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type followHander struct {
	srv followers.FollowService
}

// Follow implements followers.FollowHandler
func (fh *followHander) Follow() echo.HandlerFunc {
	return func(c echo.Context) error {
		uID := c.Param("id")
		followingID, _ := strconv.Atoi(uID)
		id, _, err := middlewares.ExtractToken(c)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, map[string]interface{}{"message": "error from server"})
		}
		err = fh.srv.Follow(id, uint(followingID))
		if id == uint(followingID) {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "you cannot follow your self"})
		}
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "you already follow this user"})
		}
		return c.JSON(http.StatusCreated, map[string]interface{}{
			"message": "success follow this user",
		})
	}
}

// ShowAllFollower implements followers.FollowHandler
func (*followHander) ShowAllFollower() echo.HandlerFunc {
	panic("unimplemented")
}

// ShowAllFollowing implements followers.FollowHandler
func (*followHander) ShowAllFollowing() echo.HandlerFunc {
	panic("unimplemented")
}

// Unfollow implements followers.FollowHandler
func (*followHander) Unfollow() echo.HandlerFunc {
	panic("unimplemented")
}

func New(srv followers.FollowService) followers.FollowHandler {
	return &followHander{
		srv: srv,
	}
}