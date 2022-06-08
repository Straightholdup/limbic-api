package emotions

import (
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"limbic/controllers/auth"
)

type handler struct {
	SerGrpcClient *grpc.ClientConn
}

func RegisterRoutes(r *gin.Engine, serGrpcClient *grpc.ClientConn) {
	h := &handler{
		SerGrpcClient: serGrpcClient,
	}

	routes := r.Group("/emotions")
	routes.POST("/delayed", auth.IsAuthenticated(), h.EmotionDelayed)
	routes.POST("/realtime", auth.IsAuthenticated(), h.EmotionRealtime)
}
