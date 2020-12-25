package server

import (
	"context"
	"testing"

	dLog "github.com/labstack/gommon/log"
	"github.com/stretchr/testify/assert"
)

func TestNewServer(t *testing.T) {
	s := New()
	ctx := context.TODO()
	defer s.Shutdown(ctx)

	assert.Equal(t, s.Server.HideBanner, true)
	assert.Equal(t, s.Server.Logger.Level(), dLog.INFO)
	assert.NotNil(t, s.Server.Logger)

	paths := []string{"/convert", "/ping", "/metrics", "/swagger/*"}
	for _, route := range s.Server.Routes() {
		assert.Contains(t, paths, route.Path)
	}
}
