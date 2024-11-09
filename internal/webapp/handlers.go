package webapp

import (
	"encoding/base64"
	"net/http"
	"strings"

	"github.com/a-h/templ"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func handleRoot(a App, c *gin.Context) error {
	ua := c.GetHeader("User-Agent")
	if strings.Contains(strings.ToLower(ua), "healthcheck") {
		c.JSON(http.StatusNoContent, nil)
		return nil
	}

	redirectTo(c, "/toast")

	return nil
}

func handleError(a App, c *gin.Context) error {
	errMsg := c.Query("error")
	if errMsg == "" {
		errMsg = "Unknown error"
	} else {
		if decodedErrMsg, err := base64.URLEncoding.DecodeString(errMsg); err != nil {
			errMsg = "Decoding error"
		} else {
			errMsg = string(decodedErrMsg)
		}
	}
	err := render(c, errorPage(
		a.Settings().AppName(),
		errMsg,
		"Return to dashboard",
		templ.URL(appPath("/admin/dashboard"))))

	if err != nil {
		// We are already in an error context and have hit another error.
		// Log it and just try redirecting to the root page. Something is
		// really wrong if we get here.
		zap.L().Error("Error rendering error page", zap.Error(err))
		c.Redirect(http.StatusFound, appPath("/"))
	}

	return nil
}

func handleToast(a App, c *gin.Context) error {
	return render(c, toastPage(a.Settings().AppName()))
}

func handleAuthRoute(a App, c *gin.Context) error {
	c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
	return nil
}
