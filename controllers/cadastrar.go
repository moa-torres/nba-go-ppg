package controllers

import (
	"context"
	"fmt"
	"log"
	"os"
	"reflect"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/moacirtorres/nba-go-ppg/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Cadastrar(c *gin.Context) {

	var player models.Player
	err := c.ShouldBindJSON(&player)
	if err != nil {
		c.JSON(500, gin.H{
			"mensagem": "Não foi possível bindar o JSON",
		})

		return
	}

	if player.Nome == "" {
		c.AbortWithError(400, err)
	}

	if player.Ppg == 0 {
		c.AbortWithError(400, err)
	}

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

	var (
		opt    options.FindOneOptions
		result models.Player
	)

	filter := bson.M{"nome": player.Nome}
	opt.SetProjection(bson.M{"nome": 1})

	err = collection.FindOne(ctx, filter, &opt).Decode(&result)
	if err != nil {
		_, err = collection.InsertOne(ctx, player)
		if err != nil {
			c.JSON(500, gin.H{
				"mensagem": "Falha ao conectar com a database e/ou a collection",
			})

			return
		}

		c.JSON(200, gin.H{
			"mensagem": "Cadastro realizado com sucesso!",
		})
	} else {
		if reflect.TypeOf(result) == reflect.TypeOf(player) {
			c.JSON(500, gin.H{
				"mensagem": "Já existe este jogador na database",
			})

			return
		}
	}

}
