package controllers

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Deletar(c *gin.Context) {
	jogador := c.Param("jogador")

	_ = godotenv.Load()

	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().
		ApplyURI(os.Getenv("MONGO_URL")).
		SetServerAPIOptions(serverAPIOptions)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Conectado ao cluster MongoDB...")

	collection := client.Database("playerDatabase").Collection("player")

	filter := bson.M{"nome": jogador}

	result, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		c.AbortWithError(400, err)
	}

	c.JSON(200, gin.H{
		"mensagem":  "Jogador deletado com sucesso!",
		"resultado": result,
	})

}
