package taskusecase

import (
	task "maestro/internal/domain/task"
	pb "maestro/internal/infra/grpc/task"
)

type TaskUsecase struct {
	TaskRepo task.TaskRepository
	pb.UnimplementedTaskServiceServer
}

func NewTaskUsecase(ur task.TaskRepository) *TaskUsecase {
	return &TaskUsecase{TaskRepo: ur}
}
