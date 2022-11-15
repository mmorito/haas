package main

import (
	"github.com/mnes/haas/services/imaging-study/openapi"
	sv "github.com/mnes/haas/services/imaging-study/src/server"
)

func (s *server) register() {
	// db := db.NewDB()
	// storage := storage.NewStorage()
	// healthcare := healthcare.NewHealthcare()
	// task := task.NewTask()

	openapi.RegisterHandlers(s.echo, sv.NewServer())
}
