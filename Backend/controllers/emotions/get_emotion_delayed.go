package emotions

import (
	"context"
	"io"
	pb "limbic/protos"
	"log"
	"mime/multipart"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func GrpcEmotionDelayed(client pb.SpeechEmotionRecognitionClient, file *multipart.FileHeader) string {
	log.Printf("here1")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	log.Printf("here2")
	stream, err := client.LoadData(ctx)
	if err != nil {
		log.Fatalf("%v.RecordRoute(_) = _, %v", client, err)
	}
	log.Printf("here3")
	fileContent, err := file.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer fileContent.Close()
	log.Printf("here4")

	nChunks := int64(0)
	buf := make([]byte, 4*1024)
	for {
		n, err := fileContent.Read(buf[:cap(buf)])
		buf = buf[:n]
		if n == 0 {
			if err == nil {
				continue
			}
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}

		stream.Send(&pb.Chunk{
			Content: buf,
		})
		nChunks++
		// process buf
		if err != nil && err != io.EOF {
			log.Fatal(err)
		}
	}

	reply, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("%v.CloseAndRecv() got error %v, want %v", stream, err, nil)
	}
	log.Printf("Route summary: %v", reply.Emotion)
	return reply.Emotion.String()
}

type Emotion struct {
	Value string `json:"value"`
}

func (h handler) EmotionDelayed(c *gin.Context) {
	file, _ := c.FormFile("file")

	client := pb.NewSpeechEmotionRecognitionClient(h.SerGrpcClient)
	emotion := Emotion{}
	emotion.Value = GrpcEmotionDelayed(client, file)

	c.JSON(http.StatusOK, emotion)
}
