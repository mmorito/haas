package main

import (
	"github.com/mnes/haas/src/services/imaging-study/openapi"
	sv "github.com/mnes/haas/src/services/imaging-study/server"
)

func (s *server) register() {
	// db := db.NewDB()
	// storage := storage.NewStorage()
	// healthcare := healthcare.NewHealthcare()
	// task := task.NewTask()

	openapi.RegisterHandlers(s.echo, sv.NewServer())
}
