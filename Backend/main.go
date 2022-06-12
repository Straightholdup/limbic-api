package main

import (
	"github.com/gin-gonic/gin"
	"limbic/controllers/auth"
	"limbic/controllers/payment"
	"limbic/controllers/users"
	"limbic/models"
)

func main() {
	dsn := "host=localhost user=root password=CSSE1810da dbname=limbic port=5432"
	db := models.Init(dsn)

	//serverAddr := flag.String("addr", "localhost:50052", "The server address in the format of host:port")
	//conn, err := grpc.Dial(*serverAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	//if err != nil {
	//	log.Fatalf("fail to dial: %v", err)
	//}
	//defer conn.Close()

	r := gin.Default()
	//emotions.RegisterRoutes(r, conn)
	users.RegisterRoutes(r, db)
	auth.RegisterRoutes(r, db)
	payment.RegisterRoutes(r, db)
	//r.GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":8080")
}
