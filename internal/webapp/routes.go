package webapp

import (
	"embed"
	"encoding/base64"
	"fmt"
	"io/fs"
	"net/http"

	"github.com/google/uuid"
	"github.com/mthurst0/tw30/internal/apperr"

	"github.com/mthurst0/tw30/internal/appenv"

	"github.com/a-h/templ"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

//go:embed static/*
var staticFS embed.FS

type App interface {
	Settings() appenv.Settings
}

func mailToAdminLink(adminEmail string) templ.SafeURL {
	return templ.URL(fmt.Sprintf("mailto:%s", adminEmail))
}

type handler func(App, *gin.Context) error

func addRoute(group *gin.RouterGroup, app App, method, path string, h handler) {
	group.Handle(method, path, func(c *gin.Context) {
		if err := h(app, c); err != nil {
			base64ErrMsg := base64.URLEncoding.EncodeToString([]byte(err.Error()))
			c.Header("HX-Redirect", appPath("/error?error="+base64ErrMsg))
			c.Status(apperr.ToHTTPStatus(err))
			err2 := render(c, errorPage(
				app.Settings().AppName(),
				err.Error(),
				"Return to dashboard",
				templ.URL(appPath("/admin/dashboard"))))
			if err2 != nil {
				// We are already in an error context and have hit another error.
				zap.L().Error("Error rendering error page", zap.Error(err2))
				c.Redirect(http.StatusFound, appPath("/admin/dashboard"))
			}
		}
		c.Abort()
	})
}

type routeAdder struct {
	app   App
	group *gin.RouterGroup
}

func (ra routeAdder) GET(path string, h handler) {
	addRoute(ra.group, ra.app, http.MethodGet, path, h)
}

func (ra routeAdder) POST(path string, h handler) {
	addRoute(ra.group, ra.app, http.MethodPost, path, h)
}

func (ra routeAdder) DELETE(path string, h handler) {
	addRoute(ra.group, ra.app, http.MethodDelete, path, h)
}

func AddRoutes(a App, baseRoutes *gin.Engine) {
	authGroup := baseRoutes.Group(appPath(""))
	authGroup.Use(setupAuthContext(a))
	noAuthGroup := baseRoutes.Group(appPath(""))

	noAuth := routeAdder{app: a, group: noAuthGroup}
	auth := routeAdder{app: a, group: authGroup}

	noAuth.GET("/", handleRoot)
	noAuth.GET("/error", handleError)
	noAuth.GET("/toast", handleToast)

	auth.GET("/auth/route", handleAuthRoute)

	// In dev mode, we serve the static files from the local filesystem to make
	// hot reloading easier during dev. In the production build, we serve via the
	// static filesystem that is bound into the binary.
	if appenv.DevMode() {
		noAuthGroup.StaticFS(
			"/static", http.Dir("internal/webapp/static"))
		baseRoutes.StaticFileFS(
			"/favicon.ico", "favicon.ico", http.Dir("internal/webapp/static"))
	} else {
		sfs, err := fs.Sub(staticFS, "static")
		if err != nil {
			zap.L().Fatal("Coder fail: could not initialize static file system", zap.Error(err))
		}
		noAuthGroup.StaticFS(
			"/static", http.FS(sfs))
		baseRoutes.StaticFileFS(
			"/favicon.ico", "static/favicon.ico", http.FS(sfs))
	}
}

func render(c *gin.Context, t templ.Component) error {
	if err := t.Render(c.Request.Context(), c.Writer); err != nil {
		zap.L().Error("Error rendering page", zap.Error(err))
		return err
	}
	return nil
}

func unauthorizedResponse(app App, c *gin.Context) {
	// This redirect ensures that the user is sent back to the login page
	// if they are in an HTMX context where responses are being handling
	// inline.
	c.Header("HX-Redirect", appPath("/login"))
	c.Status(http.StatusUnauthorized)
	// The message here is purposefully vague to avoid leaking information.
	// See the logs for diagnosing the issue.
	err := render(c, unauthorizedPage(
		app.Settings().AppName(), "Credentials are invalid or expired"))
	if err != nil {
		zap.L().Error("Error rendering unauthorized page", zap.Error(err))
	}
	c.Abort()
}

func setupAuthContext(app App) gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO
		c.Next()
	}
}

// CurrentUser represents the current logged-in user.
type CurrentUser struct {
	ClientIP string
	Username string
	ID       uuid.UUID
}

func resolveCurrentUser(a App, c *gin.Context) (*CurrentUser, error) {
	return &CurrentUser{
		ClientIP: c.ClientIP(),
		Username: "TODO",
	}, nil
}

func appPath(s string) string {
	return s
}

func redirectTo(c *gin.Context, path string) {
	c.Redirect(http.StatusFound, appPath(path))
}

func redirectToLogin(c *gin.Context) {
	redirectTo(c, "login")
}

func redirectToAppRoot(c *gin.Context) {
	redirectTo(c, "")
}
