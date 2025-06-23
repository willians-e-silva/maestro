package taskusecase

import (
	"context"
	"fmt"

	pb "maestro/internal/infra/grpc/task"
)

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
