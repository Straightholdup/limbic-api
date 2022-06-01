package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"limbic/controllers/emotions"
	"limbic/controllers/users"
	"limbic/models"
	"log"
)

func main() {
	dsn := "host=0.0.0.0 user=root password=CSSE1810da dbname=limbic port=5432"
	db := models.Init(dsn)

	serverAddr := flag.String("addr", "192.168.1.107:50052", "The server address in the format of host:port")
	conn, err := grpc.Dial(*serverAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()

	r := gin.Default()

	emotions.RegisterRoutes(r, conn)
	users.RegisterRoutes(r, db)

	r.Run(":8080")
}
