package main

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	rolInfra "github.com/Romieb26/ApIsistema_permisos/src/rol/infrastructure"
	usuarioInfra "github.com/Romieb26/ApIsistema_permisos/src/user/infrastructure"
	tutoradoInfra "github.com/Romieb26/ApIsistema_permisos/src/tutorados/infrastructure"
	docenteInfra "github.com/Romieb26/ApIsistema_permisos/src/docente/infrastructure"
	permissionInfra "github.com/Romieb26/ApIsistema_permisos/src/permission/infrastructure"

	"github.com/Romieb26/ApIsistema_permisos/src/core"
)

func main() {
	// Inicializar la conexión a la base de datos
	core.InitDB()

	// Crear router Gin
	router := gin.Default()

	// Configuración CORS abierta para cualquier front
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * 3600,
	}))

	// Middleware Logger y Recovery
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// Servir archivos estáticos desde la carpeta ./uploads
	router.Static("/uploads", "./uploads")

	// Inicializar rutas de cada módulo

	// Roles
	rolRouter := rolInfra.NewRolRouter(router)
	rolRouter.Run()

	// Usuarios
	usuarioRouter := usuarioInfra.NewUserRouter(router)
	usuarioRouter.Run()

	// Tutorados
	tutoradoRouter := tutoradoInfra.NewTutoradosRouter(router)
	tutoradoRouter.Run()

	// Docentes
	docenteRouter := docenteInfra.NewDocenteRouter(router)
	docenteRouter.Run()

	// Permisos
	permissionRouter := permissionInfra.NewPermissionRouter(router)
	permissionRouter.Run()

	// Iniciar servidor en puerto 8000
	log.Println("API inicializada en http://localhost:8000")
	log.Println("- Rutas de roles: /roles")
	log.Println("- Rutas de usuarios: /usuarios")
	log.Println("- Rutas de tutorados: /tutorados")
	log.Println("- Rutas de docentes: /docentes")
	log.Println("- Rutas de permisos: /permissions")
	log.Println("- Acceso a imágenes: http://localhost:8000/uploads/<archivo>")

	if err := router.Run(":8000"); err != nil {
		log.Fatal("Error al iniciar el servidor:", err)
	}
}
