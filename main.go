package main

import (
	"log"
	"net/http"
	"os"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/gin-gonic/gin"
)

// Company model
type Company struct {
	ID        bson.ObjectId `bson:"_id" json:"id"`
	Name      string        `bson:"name" json:"name" binding:"required"`
	Address   string        `bson:"address" json:"address" binding:"required"`
	City      string        `bson:"city" json:"city" binding:"required"`
	Zipcode   string        `bson:"zipcode" json:"zipcode" binding:"required"`
	Country   string        `bson:"country" json:"country" binding:"required"`
	Email     string        `bson:"email" json:"email"`
	Phone     string        `bson:"phone" json:"phone"`
	Owners    []string      `bson:"owners" json:"owners"`
	Directors []string      `bson:"directors" json:"directors"`
	Revisions []Company     `bson:"revisions" json:"revisions"`
}

var (
	mongoSession *mgo.Session
	dbName       string
)

func main() {

	// Connect to MongoDB
	var err error
	mongoSession, err = mgo.Dial(os.Getenv("MONGODB_URI"))
	dbName = os.Getenv("MONGODB_DB")

	if err != nil {
		log.Fatalf("Could not connect to MongoDB: %q", err)
	}

	router := gin.Default()

	// Load template
	router.LoadHTMLFiles("templates/index.html")

	// Service static files
	router.Static("/static", "./static")

	// Serve index.html
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	// API endpoints
	apiGroup := router.Group("/api")
	{
		apiGroup.GET("/companies", companiesList)
		apiGroup.GET("/companies/:id", companiesDetail)
		apiGroup.POST("/companies/:id", companiesUpdate)
		apiGroup.DELETE("/companies/:id", companiesDelete)
		apiGroup.POST("/companies", companiesCreate)
	}

	// Run until stopped
	router.Run(":" + os.Getenv("PORT"))
}

func companiesList(c *gin.Context) {
	session := mongoSession.Copy()
	defer session.Close()
	collection := session.DB(dbName).C("companies")

	// Get all Company objects
	var companies []Company
	err := collection.Find(nil).All(&companies)
	if err != nil {
		log.Printf("Error: %q", err)
		c.JSON(http.StatusInternalServerError, gin.H{"status": "Internal server error"})
		return
	}

	// If no companies are found, we create an empty array
	if companies == nil {
		companies = make([]Company, 0)
	}

	c.JSON(http.StatusOK, companies)
}

func companiesDetail(c *gin.Context) {
	var company Company
	id := c.Param("id")

	// Check if the id is valid - if not, return 404
	if !bson.IsObjectIdHex(id) {
		c.JSON(http.StatusNotFound, gin.H{"status": "Not found"})
		return
	}

	session := mongoSession.Copy()
	defer session.Close()
	collection := session.DB(dbName).C("companies")

	// Get the object or 404
	err := collection.FindId(bson.ObjectIdHex(id)).One(&company)
	if err != nil {
		if err == mgo.ErrNotFound {
			c.JSON(http.StatusNotFound, gin.H{"status": "Not found"})
		} else {
			log.Printf("Error: %q", err)
			c.JSON(http.StatusInternalServerError, gin.H{"status": "Internal server error"})
		}

		return
	}

	c.JSON(http.StatusOK, company)
}

func companiesUpdate(c *gin.Context) {
	id := c.Param("id")

	// Check if the id is valid - if not, return 404
	if !bson.IsObjectIdHex(id) {
		c.JSON(http.StatusNotFound, gin.H{"status": "Not found"})
		return
	}

	session := mongoSession.Copy()
	defer session.Close()
	collection := session.DB(dbName).C("companies")

	// Check if the request object fits in the Company model
	var company Company
	if c.BindJSON(&company) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "Bad request"})
	}

	// Find the old company object
	var revision Company
	err := collection.FindId(bson.ObjectIdHex(id)).One(&revision)
	if err != nil {
		if err == mgo.ErrNotFound {
			c.JSON(http.StatusNotFound, gin.H{"status": "Not found"})
		} else {
			log.Printf("Error: %q", err)
			c.JSON(http.StatusInternalServerError, gin.H{"status": "Internal server error"})
		}
		return
	}

	// Create a new revision
	revisions := revision.Revisions
	revision.Revisions = make([]Company, 0)
	revisions = append(revisions, revision)
	company.Revisions = revisions

	// Update the company
	err = collection.UpdateId(bson.ObjectIdHex(id), company)
	if err != nil {
		if err == mgo.ErrNotFound {
			c.JSON(http.StatusNotFound, gin.H{"status": "Not found"})
		} else {
			log.Printf("Error: %q", err)
			c.JSON(http.StatusInternalServerError, gin.H{"status": "Internal server error"})
		}
		return
	}

	c.JSON(http.StatusOK, company)
}

func companiesDelete(c *gin.Context) {
	id := c.Param("id")

	// Check if the id is valid - if not, return 404
	if !bson.IsObjectIdHex(id) {
		c.JSON(http.StatusNotFound, gin.H{"status": "Not found"})
		return
	}

	session := mongoSession.Copy()
	defer session.Close()
	collection := session.DB(dbName).C("companies")

	// Delete the object
	err := collection.RemoveId(bson.ObjectIdHex(id))
	if err != nil {
		if err == mgo.ErrNotFound { // If the object is not found
			c.JSON(http.StatusNotFound, gin.H{"status": "Not found"})
		} else {
			log.Printf("Error: %q", err)
			c.JSON(http.StatusInternalServerError, gin.H{"status": "Internal server error"})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "OK"})
}

func companiesCreate(c *gin.Context) {

	// Check if the request object fits in the Company model
	var company Company
	if c.BindJSON(&company) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "Bad request"})
		return
	}

	// Generate the id
	company.ID = bson.NewObjectId()
	session := mongoSession.Copy()
	defer session.Close()
	collection := session.DB(dbName).C("companies")

	// Save the object to MongoDB
	err := collection.Insert(company)
	if err != nil {
		log.Printf("Error: %q", err)
		c.JSON(http.StatusInternalServerError, gin.H{"status": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, company)
}
