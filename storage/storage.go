package storage

import (
	"booking/models"
	"context"
)

type StorageRepoI interface {
	User() UserRepoI
	Role() RoleRepoI
	Doctor() DoctorRepoI
	Appointment() AppointmentRepo
	CloseDB()
}

type UserRepoI interface {
	Create(ctx context.Context, req *models.CreateUser) (*models.UserPrimaryKey, error)
	GetByID(ctx context.Context, req *models.UserPrimaryKey) (*models.User, error)
	GetByUserName(ctx context.Context, req *models.Login) (*models.User, error)
	GetList(ctx context.Context, req *models.GetListUserRequest) (resp *models.GetListUserResponse, err error)
	Update(ctx context.Context, req *models.UpdateUser) (int64, error)
	Delete(ctx context.Context, req *models.UserPrimaryKey) (int64, error)
}

type RoleRepoI interface {
	Create(ctx context.Context, req *models.CreateRole) (*models.PrimaryKey, error)
	GetByID(ctx context.Context, req *models.PrimaryKey) (*models.Role, error)
	GetList(ctx context.Context, req *models.GetListRoleRequest) (resp *models.GetListRoleResponse, err error)
	Update(ctx context.Context, req *models.UpdateRole) (int64, error)
	Delete(ctx context.Context, req *models.PrimaryKey) (int64, error)
	GetByName(ctx context.Context, req *models.Role) (*models.Role, error)
}


type DoctorRepoI interface {
	Create(ctx context.Context, req *models.CreateDoctor) (*models.Doctor, error)
	GetByID(ctx context.Context, id string) (*models.Doctor, error)
	GetList(ctx context.Context, req *models.GetListDoctorRequest) (*models.GetListDoctorResponse, error)
	Update(ctx context.Context, req *models.UpdateDoctor) (int64, error)
	Delete(ctx context.Context, id string) (int64, error)
}

type AppointmentRepo interface {
	Create(ctx context.Context, req *models.CreateAppointment) (*models.Appointment, error)
	GetByID(ctx context.Context, id int) (*models.Appointment, error)
	GetList(ctx context.Context, req *models.GetListAppointmentRequest) (*models.GetListAppointmentResponse, error)
	Update(ctx context.Context, req *models.UpdateAppointment) error
	Delete(ctx context.Context, id int) error
}