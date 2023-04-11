package goapollo

import (
	"errors"
	"sync"
)

type App struct {
	AppId    string
	Clusters sync.Map
}

func (s *Server) AddApp(appId string) error {
	// Check if app already exists
	if _, ok := s.appMap.Load(appId); ok {
		return errors.New("app already exists")
	}

	// Add app
	s.appMap.Store(appId, &App{AppId: appId, Clusters: sync.Map{}})

	return nil
}

func (s *Server) DeleteApp(appId string) bool {
	// Delete app
	_, ok := s.appMap.LoadAndDelete(appId)

	return ok
}

func (s *Server) GetApps() []*App {
	// Get app list
	apps := make([]*App, 0)
	s.appMap.Range(func(key, value interface{}) bool {
		apps = append(apps, value.(*App))
		return true
	})

	return apps
}

func (s *Server) GetApp(appId string) (*App, bool) {
	// Get app
	appInterface, ok := s.appMap.Load(appId)
	if !ok {
		return nil, false
	}
	app := appInterface.(*App)

	return app, true
}
