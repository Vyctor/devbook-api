package routes

import (
	"devbook-api/src/controllers"
	"net/http"
)

var postsRoutes = []Route{
	{
		Uri:                   "/posts",
		Method:                http.MethodPost,
		Function:              controllers.CreatePost,
		RequireAuthentication: true,
	},
	{
		Uri:                   "/posts",
		Method:                http.MethodGet,
		Function:              controllers.FindPosts,
		RequireAuthentication: true,
	},
	{
		Uri:                   "/posts/{id}",
		Method:                http.MethodGet,
		Function:              controllers.FindPost,
		RequireAuthentication: true,
	},
	{
		Uri:                   "/posts/{id}",
		Method:                http.MethodPut,
		Function:              controllers.UpdatePost,
		RequireAuthentication: true,
	},
	{
		Uri:                   "/posts/{id}",
		Method:                http.MethodDelete,
		Function:              controllers.DeletePost,
		RequireAuthentication: true,
	},
	{
		Uri:                   "/users/{id}/posts",
		Method:                http.MethodGet,
		Function:              controllers.FindUserPosts,
		RequireAuthentication: true,
	},
	{
		Uri:                   "/posts/{id}/like",
		Method:                http.MethodPost,
		Function:              controllers.LikePost,
		RequireAuthentication: true,
	},
	{
		Uri:                   "/posts/{id}/unlike",
		Method:                http.MethodPost,
		Function:              controllers.Unlike,
		RequireAuthentication: true,
	},
}
