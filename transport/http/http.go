package http

// You should'nt need to touch this file.
// For middlewares use middlewares.go
// For routes use routes.go
// For creating http handler structs use handlers.go
import (
	"encoding/json"
	"io"

	"github.com/labstack/echo/v4"
)

// HTTPServer just serves the thing
type HTTPServer func(address string) error

func Routes(w io.Writer) {
	e := echo.New()
	configure(e)
	bs, _ := json.MarshalIndent(e.Routes(), "", "  ")
	w.Write(bs)
}

type RouteRegister interface {
	Register(e *echo.Echo)
}

func configure(e *echo.Echo) {
	// register global middlewares middlewares.go
	e.Use(middlewares()...)
	// register routes defined in routes.go
	for _, routeRegister := range RouteRegisterers {
		routeRegister.Register(e)
	}

}
func NewHTTPServer() HTTPServer {
	e := echo.New()

	return func(address string) error {
		configure(e)
		return e.Start(address)
	}

}
