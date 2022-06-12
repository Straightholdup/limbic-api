package main

import (
	"flag"
	"limbic/controllers/auth"
	"limbic/controllers/emotions"
	"limbic/controllers/users"
	"limbic/models"
	"log"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	dsn := "host=db user=root password=CSSE1810da dbname=limbic port=5432"
	db := models.Init(dsn)

	serverAddr := flag.String("addr", "ser_service:50052", "The server address in the format of host:port")
	conn, err := grpc.Dial(*serverAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()

	r := gin.Default()
	emotions.RegisterRoutes(r, conn)
	users.RegisterRoutes(r, db)
	auth.RegisterRoutes(r, db)

	r.Run(":80")
}
