package main
import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

func GetHomePage(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Get Page",
	})
}

func PostHomePage(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Post Page",
	})
}

func Create(c *gin.Context) {
	// Redis integration
	Client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "", // no password
		DB: 0, // default DB
	})

	id:=c.Query("id")
	name:=c.Query("name")

	val, err := Client.HSet("test", id, name).Result()
    if err != nil {
        fmt.Println(err)
    }
	fmt.Println(val)
	
	c.JSON(200, gin.H{
		"message": "New User Added!",
	})
}

func Read(c *gin.Context) {
	// Redis integration
	Client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "", // no password
		DB: 0, // default DB
	})

	val, err := Client.HGetAll("test").Result()
    if err != nil {
        fmt.Println(err)
    }
	fmt.Println(val)
	
	c.JSON(200, gin.H{
		"message": val,
	})
}

func Update(c *gin.Context) {
	// Redis integration
	Client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "", // no password
		DB: 0, // default DB
	})

	id:=c.Query("id")
	name:=c.Query("name")

	val, err := Client.HSet("test", id, name).Result()
    if err != nil {
        fmt.Println(err)
    }
	fmt.Println(val)
	
	c.JSON(200, gin.H{
		"message": "User Updated!",
	})
}

func Delete(c *gin.Context) {
	// Redis integration
	Client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "", // no password
		DB: 0, // default DB
	})

	id:=c.Query("id")

	val, err := Client.HDel("test", id).Result()
    if err != nil {
        fmt.Println(err)
    }
	fmt.Println(val)
	
	c.JSON(200, gin.H{
		"message": "User Deleted!",
	})
}

func ReadOne(c *gin.Context) {
	// Redis integration
	Client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "", // no password
		DB: 0, // default DB
	})

	id:=c.Query("id")

	val, err := Client.HGet("test", id).Result()
    if err != nil {
        fmt.Println(err)
    }
	fmt.Println(val)
	
	c.JSON(200, gin.H{
		"message": val,
	})
}

func main() {
	
	// Redis integration
	// Client := redis.NewClient(&redis.Options{
	// 	Addr: "localhost:6379",
	// 	Password: "", // no password
	// 	DB: 0, // default DB
	// })
	r := gin.Default()
	r.GET("/", GetHomePage)
	r.POST("/", PostHomePage)
	r.POST("/create", Create)
	r.GET("/read", Read)
	r.POST("/update", Update)
	r.POST("/delete", Delete)
	r.GET("/read_one", ReadOne)
	r.Run()
}
