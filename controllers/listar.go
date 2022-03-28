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

func Listar(c *gin.Context) {

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

	cur, err := collection.Find(ctx, bson.D{})
	if err != nil {
		c.AbortWithError(400, err)
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var result bson.D
		err := cur.Decode(&result)
		if err != nil {
			c.AbortWithError(400, err)
		}
		c.JSON(200, gin.H{
			"result": result,
		})
	}
	if err := cur.Err(); err != nil {
		c.AbortWithError(400, err)
	}
}

func ListarUm(c *gin.Context) {

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

	var (
		opt    options.FindOneOptions
		result models.Player
	)

	filter := bson.M{"nome": name}
	opt.SetProjection(bson.M{"nome": 1, "ppg": 1})

	err = collection.FindOne(ctx, filter, &opt).Decode(&result)
	if err != nil {
		c.JSON(404, "Nenhum jogador encontrado. Por favor, busque por um jogador do TOP 10 Scorer da atual temporada (21/22)!")
	} else {

		c.JSON(200, gin.H{
			"mensagem":  "Jogador encontrado!",
			"resultado": result,
		})

	}
}
