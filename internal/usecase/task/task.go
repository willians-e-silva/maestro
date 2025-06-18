package usecase

import (
	"context"
	"fmt"
	"time"

	task "github.com/willians-e-silva/maestro/internal/domain/task"
	pb "github.com/willians-e-silva/maestro/internal/infra/grpc/task"
)

type TaskUsecase struct {
	TaskRepo task.TaskRepository
	pb.UnimplementedTaskServiceServer
}

func NewUserUsecase(ur task.TaskRepository) *TaskUsecase {
	return &TaskUsecase{TaskRepo: ur}
}

func (uc *TaskUsecase) CreateTask(ctx context.Context, req *pb.CreateTaskRequest) (*pb.TaskResponse, error) {

	task := &task.Task{
		Name:          req.GetName(),
		Description:   req.GetDescription(),
		DockerImage:   req.GetDockerImage(),
		DockerTag:     req.GetDockerTag(),
		Port:          req.GetPort(),
		EnvVars:       req.GetEnvVars(),
		Volumes:       req.GetVolumes(),
		NetworkMode:   req.GetNetworkMode(),
		RestartPolicy: req.GetRestartPolicy(),
		CpuLimit:      req.GetCpuLimit(),
		MemoryLimit:   req.GetMemoryLimit(),
		CreatedAt:     time.Now(),
	}

	createdTask, err := uc.TaskRepo.CreateTask(task)

	if err != nil {
		return nil, fmt.Errorf("falha ao criar tarefa: %w", err)
	}

	return &pb.TaskResponse{
		Uuid:        createdTask.UUID,
		Name:        createdTask.Name,
		Description: createdTask.Description,
	}, nil
}
