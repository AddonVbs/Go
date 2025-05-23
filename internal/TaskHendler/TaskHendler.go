package taskhendler

import taskservers "BackEnd/internal/taskServers"

type TaskHendler struct {
	service taskservers.TaskServers
}

func new