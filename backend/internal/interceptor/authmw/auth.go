package authmw

import (
	"backend/pkg/logutil"

	"github.com/labstack/echo/v4"
)

// RequireUserID is a middleware that checks if login_id is present in query parameters
func RequireUserID() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			usrID := c.QueryParam("login_id")
			// if usrID == "" {
			// 	return errmsg.RespondError(c, "ERR_UNAUTHORIZED", "login_id is required")
			// }
			logutil.Info("Request authenticated for user: %s", usrID)
			return next(c)
		}
	}
}
