package api

import (
	"database/sql"

	"github.com/gin-gonic/gin"

	"github.com/asadbekGo/golang_crud/api/handler"
)

func SetUpApi(r *gin.Engine, db *sql.DB) {

	handlerV1 := handler.NewHandlerV1(db)

	r.POST("/user", handlerV1.Create)
	r.GET("/user/:id", handlerV1.GetById)
	r.GET("/user", handlerV1.GetList)
	r.PUT("/user/:id", handlerV1.Update)
	r.PATCH("/user/:id", handlerV1.Patch)
	r.DELETE("/user/:id", handlerV1.Delete)
}
