package cmd

import (
	"fmt"
	"net/http"
	"project-POS-APP-golang-be-team/pkg/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func ApiServer(config utils.Configuration, logger *zap.Logger, h *gin.Engine) {
	fmt.Println("Server running on port 8080")
	if err := http.ListenAndServe(":8080", h); err != nil {
		logger.Fatal("can't run service", zap.Error(err))
	}
}
