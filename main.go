package main

import (
	config "builder/restful-api-gogin/application/database"
	route "builder/restful-api-gogin/application/routes"
	util "builder/restful-api-gogin/utils"
	"log"

	helmet "github.com/danielkov/gin-helmet"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

/*
// album represents data about a record album.

	type album struct {
		ID     string  `json:"id"`
		Title  string  `json:"title"`
		Artist string  `json:"artist"`
		Price  float64 `json:"price"`
	}

// getAlbums responds with the list of all albums as JSON.

	func getAlbums(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, albums)
	}

// albums slice to seed record album data.

	var albums = []album{
		{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
		{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
		{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
		{ID: "4", Title: "The Modern Sound of Betty Carter", Artist: "Betty Carter", Price: 49.99},
	}

// postAlbums adds an album from JSON received in the request body.

	func postAlbums(c *gin.Context) {
		var newAlbum album

		// Call BindJSON to bind the received JSON to
		// newAlbum.
		if err := c.BindJSON(&newAlbum); err != nil {
			return
		}

		// Add the new album to the slice.
		albums = append(albums, newAlbum)
		c.IndentedJSON(http.StatusCreated, newAlbum)
	}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.

	func getAlbumByID(c *gin.Context) {
		id := c.Param("id")

		// Loop over the list of albums, looking for
		// an album whose ID value matches the parameter.
		for _, a := range albums {
			if a.ID == id {
				c.IndentedJSON(http.StatusOK, a)
				return
			}
		}
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
	}
*/
func main() {

	router := SetupRouter()

	/**
	*@ Run Server
	 */
	log.Fatal(router.Run("localhost:" + util.GodotEnv("GO_PORT")))

	/*
		|@ handleNotFound
	*/
	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	// App Run
	//router.Run("localhost:8080")
}

func SetupRouter() *gin.Engine {
	/**
	@description Setup Database Connection
	*/
	db := config.Connection()
	/**
	@description Init Router
	*/
	router := gin.Default()
	/**
	@description Setup Mode Application
	*/
	if util.GodotEnv("GO_ENV") != "production" && util.GodotEnv("GO_ENV") != "test" {
		gin.SetMode(gin.DebugMode)
	} else if util.GodotEnv("GO_ENV") == "test" {
		gin.SetMode(gin.TestMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	/**
	@description Setup Middleware
	*/
	router.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"*"},
		AllowHeaders:  []string{"*"},
		AllowWildcard: true,
	}))
	router.Use(helmet.Default())
	router.Use(gzip.Gzip(gzip.BestCompression))
	/**
	@description Init All Route
	*/
	route.InitRoutes(db, router)

	return router
}
