package controllers

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/moacirtorres/nba-go-ppg/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Atualizar(c *gin.Context) {

	var player models.Player
	err := c.ShouldBindJSON(&player)
	if err != nil {
		c.JSON(500, gin.H{
			"mensagem": "Não foi possível bindar o JSON",
		})

		return
	}

	name := c.Param("jogador")

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

	filter := bson.M{"nome": name}
	update := bson.D{
		{"$set",
			bson.D{
				{"nome", player.Nome},
				{"ppg", player.Ppg},
			},
		},
	}

	result, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		c.AbortWithError(400, err)
	}

	c.JSON(200, gin.H{
		"resultado": result,
	})

}
