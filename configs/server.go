package configs

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type ServerConfig struct {
	Port           string
	GinHandler     *gin.Engine
	ReadTimeout    time.Duration
	WriteTimeout   time.Duration
	MaxHeaderBytes int
}

func (c *ServerConfig) Load() *http.Server {
	// set default config server
	if c.Port == "" {
		c.Port = ":8080"
	}

	if c.ReadTimeout == 0 {
		c.ReadTimeout = 10 * time.Second
	}

	if c.WriteTimeout == 0 {
		c.WriteTimeout = 10 * time.Second
	}

	if c.MaxHeaderBytes == 0 {
		c.MaxHeaderBytes = 1 << 20
	}

	s := &http.Server{
		Addr:           c.Port,
		Handler:        c.GinHandler,
		ReadTimeout:    c.ReadTimeout,
		WriteTimeout:   c.WriteTimeout,
		MaxHeaderBytes: c.MaxHeaderBytes,
	}

	return s
}
