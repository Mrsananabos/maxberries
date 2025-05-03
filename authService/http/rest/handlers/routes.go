package handlers

import (
	"authService/http/rest/handlers/auth"
	"authService/http/rest/handlers/permission"
	"authService/http/rest/handlers/role"
	"authService/http/rest/handlers/user"
	"authService/http/rest/middleware"
	"authService/internal/servicesStorage"
	"github.com/gin-gonic/gin"
)

func Register(gin *gin.Engine, services servicesStorage.ServicesStorage) {
	roleHandler := role.NewHandler(services)
	permissionHandler := permission.NewHandler(services)
	userHandler := user.NewHandler(services)
	authHandler := auth.NewHandler(services)

	jwrMiddleware := middleware.CreateJWTMiddleware(services.AuthService)

	gin.GET("/role", jwrMiddleware.PermissionCheckMiddleware("auth.getRoles"), roleHandler.GetAllRoles)
	gin.GET("/role/:id", jwrMiddleware.PermissionCheckMiddleware("auth.getRoleById"), roleHandler.GetRoleById)
	gin.POST("/role", jwrMiddleware.PermissionCheckMiddleware("auth.createRole"), roleHandler.CreateRole)
	gin.PUT("/role/:id", jwrMiddleware.PermissionCheckMiddleware("auth.editRole"), roleHandler.UpdateRole)
	gin.DELETE("/role/:id", jwrMiddleware.PermissionCheckMiddleware("auth.deleteRole"), roleHandler.DeleteRole)

	gin.GET("/permission", jwrMiddleware.PermissionCheckMiddleware("auth.getPermission"), permissionHandler.GetAllPermissions)
	gin.GET("/permission/:id", jwrMiddleware.PermissionCheckMiddleware("auth.getPermissionById"), permissionHandler.GetPermissionById)
	gin.POST("/permission", jwrMiddleware.PermissionCheckMiddleware("auth.createPermission"), permissionHandler.CreatePermission)
	gin.PUT("/permission/:id", jwrMiddleware.PermissionCheckMiddleware("auth.updatePermission"), permissionHandler.UpdatePermission)
	gin.DELETE("/permission/:id", jwrMiddleware.PermissionCheckMiddleware("auth.deletePermission"), permissionHandler.DeletePermission)

	gin.GET("/user", jwrMiddleware.PermissionCheckMiddleware("auth.getUsers"), userHandler.GetAll)
	gin.GET("/user/:id", jwrMiddleware.UserPermissionCheckMiddleware("auth.getUserById"), userHandler.GetById)
	gin.PUT("/user/:id", jwrMiddleware.UserPermissionCheckMiddleware("auth.updateUser"), userHandler.UpdateUser)
	gin.PATCH("/user/:id", jwrMiddleware.PermissionCheckMiddleware("auth.updateUserRole"), userHandler.UpdateUserRole)
	gin.DELETE("/user/:id", jwrMiddleware.UserPermissionCheckMiddleware("auth.deleteUser"), userHandler.DeleteUser)

	gin.GET("/auth/me", authHandler.Auth)
	gin.POST("/auth/register", authHandler.Register)
	gin.POST("/auth/login", authHandler.Login)
	gin.POST("/auth/refresh", authHandler.RefreshTokens)
}
