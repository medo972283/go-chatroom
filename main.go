package main

import (
	"context"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"time"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/medo972283/go-chatroom/controllers"
	"github.com/medo972283/go-chatroom/controllers/api"
)

const (
	PORT = ":3000"
)

// Template subdirectories
var tmplDirs = []string{
	"templates/*.html",
	"templates/*/*.html",
}

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

// Glob all directories and returns all the matching file names
func globAllFiles(dirs []string) []string {
	tmplFileName := []string{}

	for _, dir := range dirs {
		// Get the names of files matching the dir pattern
		files, err := filepath.Glob(dir)
		if err != nil {
			log.Fatalf("filepath.Glob err: %v", err)
		}
		tmplFileName = append(tmplFileName, files...)
	}

	return tmplFileName
}

// Entry point
func main() {
	e := echo.New()

	/* Attach middleware service */

	// Add logger middleware to log the information about each HTTP request.
	e.Use(middleware.Logger())
	// Add recover middleware to recover from panics anywhere in the chain, prints stack trace and handles the control to the centralized HTTPErrorHandler.
	e.Use(middleware.Recover())
	// Add HTTP session middleware.
	e.Use(session.Middleware(sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))))

	/* Compile and register the front-end resources */

	// Registers the route with path prefix to serve static files from the provided root directory.
	e.Static("/", "assets")

	// Get templates
	templates := globAllFiles(tmplDirs)

	// Pre-compile all templates
	t := &Template{
		templates: template.Must(template.ParseFiles(templates...)),
	}

	// Register templates
	e.Renderer = t

	/* Register router handler */

	// attach router handler
	controllers.AttachHandler(e)

	/* Server Execution */

	// Start websocket's clinet management
	go api.Manager.Start()

	// Listen port & start server
	go func() {
		if err := e.Start(PORT); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatalf("shutting down the server %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	// Block program until quit channel recevive the interrupt signal
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
