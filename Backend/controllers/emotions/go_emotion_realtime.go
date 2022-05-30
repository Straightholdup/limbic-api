package emotions

import (
	"github.com/gin-gonic/gin"
	pb "limbic/protos"
	"net/http"
)

func (h handler) EmotionRealtime(c *gin.Context) {
	file, _ := c.FormFile("file")

	client := pb.NewSpeechEmotionRecognitionClient(h.SerGrpcClient)
	emotion := Emotion{}
	emotion.Value = GrpcEmotionDelayed(client, file)

	//for _, file := range files {
	//	log.Println(file.Filename)
	//
	//	// Upload the file to specific dst.
	//	c.SaveUploadedFile(file, "media/"+file.Filename)
	//}
	c.JSON(http.StatusOK, emotion)
}
