package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	ctx := context.Background()

	// backend connections
	postgresUrl := os.Getenv("DATABASE_URL")

	dbpool, err := pgxpool.New(ctx, postgresUrl)
	if err != nil {
		log.Fatal("Cannot connect to postgres")
		panic(err)
	}
	defer dbpool.Close()

	// api
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{ "message": "pong"})
	})

	r.POST("/api/new", func(c *gin.Context) {
		newLink(c, dbpool)
	})


	r.GET("/:slug", func(c *gin.Context) {
		getLink(c, dbpool)
	})

	r.Run() // listen and serve on localhost:8080
}
// checks redis, then postgres for url. if found in postgres, upgrades it to redis.
// LEFTOFF, figure out better order for logic, work on redis failovering to postgres
func getLink(c *gin.Context, dbpool *pgxpool.Pool) {
    ctx := c.Request.Context()
    slug := c.Param("slug")
    println(slug)

	var url string

    err := dbpool.QueryRow(ctx, "SELECT url FROM url_mappings WHERE slug=$1", slug).Scan(&url)
    if err != nil {
        fmt.Println("unknown slug:" + slug)
        fmt.Println(err)
        c.JSON(404, gin.H{"error": "unknown slug"})
        return
    }

    // return redirect
    c.Redirect(302, "https://" + url)
}

func newLink(c *gin.Context, dbpool *pgxpool.Pool) {
    ctx := c.Request.Context()
    url := c.Query("url")
    
    // ensure fqdn

    // create uuid for slug
    uuidV7, err := uuid.NewV7()
    if err != nil {
        println(err)
        c.JSON(500, gin.H{"error": "internal"})
        return
    }

    // use last 6 of the uuid
    slug := strings.ToLower(uuidV7.String()[len(uuidV7.String())-6:])

    _, err = dbpool.Exec(ctx, "INSERT INTO url_mappings (slug, url, time) VALUES ($1, $2, NOW())", slug, url)
    if err != nil {
        c.JSON(500, gin.H{"error": "Database error", "details": err.Error()})
        return
    }

    c.JSON(200, gin.H{"slug" : slug})
}
