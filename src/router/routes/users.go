package routes

import (
	"devbook-api/src/controllers"
	"net/http"
)

var usersRoutes = []Route{
	{
		Uri:                   "/users",
		Method:                http.MethodPost,
		Function:              controllers.CreateUser,
		RequireAuthentication: true,
	},
	{
		Uri:                   "/users",
		Method:                http.MethodGet,
		Function:              controllers.GetUsers,
		RequireAuthentication: true,
	},
	{
		Uri:                   "/users/{id}",
		Method:                http.MethodGet,
		Function:              controllers.GetUser,
		RequireAuthentication: true,
	},
	{
		Uri:                   "/users/{id}",
		Method:                http.MethodPut,
		Function:              controllers.UpdateUser,
		RequireAuthentication: true,
	},
	{
		Uri:                   "/users/{id}",
		Method:                http.MethodDelete,
		Function:              controllers.DeleteUser,
		RequireAuthentication: true,
	},
	{
		Uri:                   "/users/{id}/follow",
		Method:                http.MethodPost,
		Function:              controllers.FollowUser,
		RequireAuthentication: true,
	},
	{
		Uri:                   "/users/{id}/unfollow",
		Method:                http.MethodPost,
		Function:              controllers.UnfollowUser,
		RequireAuthentication: true,
	},
	{
		Uri:                   "/users/{id}/followers",
		Method:                http.MethodGet,
		Function:              controllers.GetFollowers,
		RequireAuthentication: true,
	},
	{
		Uri:                   "/users/{id}/following",
		Method:                http.MethodGet,
		Function:              controllers.GetFollowing,
		RequireAuthentication: true,
	},
	{
		Uri:                   "/users/{id}/update-password",
		Method:                http.MethodPost,
		Function:              controllers.UpdatePassword,
		RequireAuthentication: true,
	},
}
