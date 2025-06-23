package taskusecase

import (
	"context"
	"fmt"
	"time"

	task "maestro/internal/domain/task"
	pb "maestro/internal/infra/grpc/task"

	"github.com/google/uuid"
)

type TaskUsecase struct {
	TaskRepo task.TaskRepository
	pb.UnimplementedTaskServiceServer
}

func NewTaskUsecase(ur task.TaskRepository) *TaskUsecase {
	return &TaskUsecase{TaskRepo: ur}
}

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

func (uc *TaskUsecase) GetAllTasks(ctx context.Context, req *pb.GetAllTasksRequest) (*pb.GetAllTasksResponse, error) {
	tasks, err := uc.TaskRepo.GetAllTasks()

	if err != nil {
		return nil, fmt.Errorf("failed to get tasks: %w", err)
	}

	var pbTasks []*pb.Task
	for _, t := range tasks {
		pbTasks = append(pbTasks, &pb.Task{
			Id:            t.Id,
			Name:          t.Name,
			Description:   t.Description,
			DockerImage:   t.DockerImage,
			DockerTag:     t.DockerTag,
			Port:          t.Port,
			NetworkMode:   t.NetworkMode,
			RestartPolicy: t.RestartPolicy,
			CpuLimit:      t.CpuLimit,
			MemoryLimit:   t.MemoryLimit,
			CreatedAt:     t.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt:     t.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	return &pb.GetAllTasksResponse{
		Tasks: pbTasks,
	}, nil
}
