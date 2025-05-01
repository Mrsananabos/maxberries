package servicesStorage

import (
	"authService/configs"
	auth "authService/internal/auth/service"
	jwt "authService/internal/jwt/service"
	permissionRepo "authService/internal/permission/repository"
	permission "authService/internal/permission/service"
	roleRepo "authService/internal/role/repository"
	role "authService/internal/role/service"
	userRepo "authService/internal/user/repository"
	user "authService/internal/user/service"
	redisCli "authService/pkg/redis"
	"gorm.io/gorm"
)

type ServicesStorage struct {
	RoleService       role.Service
	PermissionService permission.Service
	UserService       user.Service
	AuthService       auth.Service
}

func NewServicesStorage(config configs.Config, db *gorm.DB) (ServicesStorage, error) {
	redisClient, err := redisCli.Connect(config.Redis)
	if err != nil {
		return ServicesStorage{}, err
	}
	roleService := role.NewService(roleRepo.NewRepository(db))
	permissionService := permission.NewService(permissionRepo.NewRepository(db))
	userService := user.NewService(userRepo.NewRepository(db), roleService)
	jwtService := jwt.NewService(config.JWTConfig, userService, roleService, redisClient)
	authService := auth.NewService(userService, roleService, jwtService)

	return ServicesStorage{
		RoleService:       roleService,
		PermissionService: permissionService,
		UserService:       userService,
		AuthService:       authService,
	}, nil

}
