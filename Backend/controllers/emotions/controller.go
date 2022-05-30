package emotions

import (
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

type handler struct {
	SerGrpcClient *grpc.ClientConn
}

func RegisterRoutes(r *gin.Engine, serGrpcClient *grpc.ClientConn) {
	h := &handler{
		SerGrpcClient: serGrpcClient,
	}

	routes := r.Group("/emotions")
	routes.POST("/delayed", h.EmotionDelayed)
	routes.POST("/realtime", h.EmotionRealtime)
}
