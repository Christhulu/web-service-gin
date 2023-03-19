package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Http methods
// postRecipes adds a recipe from JSON received in the request body.
func postRecipes(c *gin.Context) {
	var newRecipe recipe

	// Call BindJSON to bind the received JSON to
	// newRecipe.
	if err := c.BindJSON(&newRecipe); err != nil {
		return
	}

	// Add the new recipe to the slice.
	recipes = append(recipes, newRecipe)
	c.IndentedJSON(http.StatusCreated, newRecipe)
}

// getRecipeByID locates the recipe whose ID value matches the id
// parameter sent by the client, then returns that recipe as a response.
func getRecipeByID(c *gin.Context) {
	id := c.Param("id")

	// Loop over the list of recipes, looking for
	// an recipe whose ID value matches the parameter.
	for _, a := range recipes {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "recipe not found"})
}

type recipe struct {
	ID     string `json:"id`
	Name   string `json:"name"`
	Meal   string `json:"meal"`
	Rating int16  `json:"rating"`
}

// getRecipes responds with the list of all recipes as JSON.
func getRecipes(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, recipes)
}

var recipes = []recipe{
	{ID: "1", Name: "Apple Bread", Meal: "Breakfast/Dessert", Rating: 5},
	{ID: "2", Name: "Skillet Cheeseburger Mac and Cheese", Meal: "Lunch/Dinner", Rating: 5},
	{ID: "3", Name: "Fresh Lemon Lime Cake", Meal: "Breakfast/Dessert", Rating: 5},
	{ID: "4", Name: "Soft Boiled Eggs ", Meal: "Breakfast", Rating: 3},
	{ID: "5", Name: "Nui Xao Bo", Meal: "Lunch/Dinner", Rating: 5},
	{ID: "6", Name: "Chicken Tetrazzini", Meal: "Lunch/Dinner", Rating: 4},
}

func main() {
	router := gin.Default()
	router.GET("/recipes", getRecipes)
	router.POST("/recipes", postRecipes)
	router.GET("/recipes/:id", getRecipeByID)
	router.Run("localhost:8080")
}
