package server

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	dLog "github.com/labstack/gommon/log"
	"github.com/mxschmitt/playwright-go"
	"github.com/starptech/fay/internals/controller"
	"github.com/swaggo/echo-swagger"
	"golang.org/x/net/context"
)

type Server struct {
	browser *playwright.Browser
	pw      *playwright.Playwright
	Server  *echo.Echo
}

func New() *Server {
	pw, err := playwright.Run()
	if err != nil {
		log.Fatalln("Could not run playwright")
	}
	browser, err := pw.Chromium.Launch()
	if err != nil {
		log.Fatalln("Could not launch browser")
	}

	e := echo.New()
	e.Use(middleware.BodyLimit("2M"))
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.HideBanner = true
	e.Logger.SetLevel(dLog.INFO)

	httpController := controller.Http{
		Browser: browser,
	}

	e.POST("/convert", httpController.ConvertHTML)
	e.GET("/ping", httpController.Ping)
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	return &Server{
		browser: browser,
		pw:      pw,
		Server:  e,
	}
}

func (s *Server) Shutdown(ctx context.Context) error {
	if err := s.browser.Close(); err != nil {
		return err
	}
	if err := s.pw.Stop(); err != nil {
		return err
	}
	if err := s.Server.Shutdown(ctx); err != nil {
		return err
	}
	return nil
}
