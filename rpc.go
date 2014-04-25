package main

import (
	"errors"

	"github.com/flynn/strowger/types"
)

type Router struct {
	s Server
}

func (r *Router) AddHTTPRoute(config *strowger.Config, res *struct{}) error {
	switch config.Type {
	case strowger.RouteHTTP:
		if err := r.s.AddHTTPDomain(config.HTTPDomain, config.Service, config.HTTPSCert, config.HTTPSKey); err != nil {
			return err
		}
	default:
		return errors.New("unsupported frontend type")
	}
	return nil
}

func (r *Router) RemoveRoute(config *strowger.Config, res *struct{}) error {
	switch config.Type {
	case strowger.RouteHTTP:
		if err := r.s.RemoveHTTPDomain(config.HTTPDomain); err != nil {
			return err
		}
	default:
		return errors.New("unsupported frontend type")
	}
	return nil
}

func (r *Router) UpdateRoute(config *strowger.Config, res *struct{}) error {
	switch config.Type {
	case strowger.RouteHTTP:
		if err := r.s.RemoveHTTPDomain(config.HTTPDomain); err != nil {
			return err
		}
	default:
		return errors.New("unsupported frontend type")
	}
	return nil
}

func (r *Router) ListRoutes(config *strowger.Config, res *struct{}) error {
	// List all? Only a specific type?
	return nil
}
