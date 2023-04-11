package goapollo

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ServicesConfig struct {
	AppName     string `json:"appName"`
	InstanceId  string `json:"instanceId"`
	HomepageUrl string `json:"homepageUrl"`
}

func (s *Server) getServicesConfig(c *gin.Context) {
	// Return result
	c.JSON(http.StatusOK, []*ServicesConfig{
		{
			AppName:     "MyAppName",
			InstanceId:  "goapollo",
			HomepageUrl: fmt.Sprintf("http://localhost:%d", s.config.Port),
		},
	})
}
