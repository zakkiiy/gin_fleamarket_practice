package main

import (
	"fmt"
	"gin_fleamarket_practice/models"
	"gin_fleamarket_practice/repositories"

	"github.com/gin-gonic/gin"
)

func main() {
	items := []models.Item{
		{ID: 1, Name: "商品1", Price: 1000, Description: "説明1", SoldOut: false},
		{ID: 2, Name: "商品2", Price: 2000, Description: "説明2", SoldOut: true},
		{ID: 3, Name: "商品3", Price: 3000, Description: "説明3", SoldOut: false},
	}

	fmt.Println(items)
	fmt.Printf("%+v\n", items)

	repo := repositories.NewItemMemoryRepository(items)
	fmt.Println("れぽ")
	fmt.Println(repo)
	fmt.Printf("%+v\n", repo)

  r := gin.Default()
  r.GET("/ping", func(c *gin.Context) {
    c.JSON(200, gin.H{
      "message": "pong",
    })
  })

	r.GET("/items", func(c *gin.Context) {
		result, err := repo.FindAll()
		if err != nil {
			c.JSON(500, gin.H{"error": "internal server error"})
			return
		}
		c.JSON(200, gin.H{
			"items": result,
		})
	})
  r.Run() // 0.0.0.0:8080 でサーバーを立てます。
}
