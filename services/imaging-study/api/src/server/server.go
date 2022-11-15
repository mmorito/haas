package server

import "github.com/mnes/haas/services/imaging-study/openapi"

type server struct {
}

func NewServer() openapi.ServerInterface {
	return &server{}
}
