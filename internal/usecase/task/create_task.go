package taskusecase

import (
	"context"
	"fmt"
	"time"

	task "maestro/internal/domain/task"
	pb "maestro/internal/infra/grpc/task"

	"github.com/google/uuid"
)

func (uc *TaskUsecase) CreateTask(ctx context.Context, req *pb.CreateTaskRequest) (*pb.CreateTaskResponse, error) {
	task := &task.Task{
		Id:            uuid.New().String(),
		Name:          req.GetName(),
		Description:   req.GetDescription(),
		DockerImage:   req.GetDockerImage(),
		DockerTag:     req.GetDockerTag(),
		Port:          req.GetPort(),
		NetworkMode:   req.GetNetworkMode(),
		RestartPolicy: req.GetRestartPolicy(),
		CpuLimit:      req.GetCpuLimit(),
		MemoryLimit:   req.GetMemoryLimit(),
		CreatedAt:     time.Now(),
	}

	createdTask, err := uc.TaskRepo.CreateTask(task)
	if err != nil {
		return nil, fmt.Errorf("failed to create task: %w", err)
	}

	return &pb.CreateTaskResponse{
		Id:          createdTask.Id,
		Name:        createdTask.Name,
		Description: createdTask.Description,
	}, nil
}
