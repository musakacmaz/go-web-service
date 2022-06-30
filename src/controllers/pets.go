package controllers

import (
	"net/http"

	_ "go-pet-api/docs"

	"github.com/gin-gonic/gin"
)

type pet struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Status string `json:"status"`
	Type   string `json:"type"`
	Breed  string `json:"breed"`
}

var pets = []pet{
	{ID: "1", Name: "dog", Status: "available", Type: "dog", Breed: "labrador"},
	{ID: "2", Name: "cat", Status: "available", Type: "cat", Breed: "persian"},
	{ID: "3", Name: "bird", Status: "available", Type: "bird", Breed: "parrot"},
}

// @Summary get all items in the pet list
// @ID get-pets
// @Produce json
// @Success 200 {object} pet
// @Router /pets [get]
func GetPets(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, pets)
}

// @Summary add a new item to the pet list
// @ID create-pet
// @Produce json
// @Param data body pet true "Pet Data"
// @Success 200 {object} pet
// @Router /pets [post]
func CreatePet(ctx *gin.Context) {
	var newPet pet

	if err := ctx.ShouldBindJSON(&newPet); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	pets = append(pets, newPet)
	ctx.IndentedJSON(http.StatusCreated, newPet)
}

// @Summary get a pet by ID
// @ID get-pet-by-id
// @Produce json
// @Param id path string true "Pet ID"
// @Success 200 {object} pet
// @Router /pets/{id} [get]
func GetPetsByID(ctx *gin.Context) {
	id := ctx.Param("id")
	for _, pet := range pets {
		if pet.ID == id {
			ctx.IndentedJSON(http.StatusOK, pet)
			return
		}
	}
	ctx.IndentedJSON(http.StatusNotFound, gin.H{"error": "Pet not found"})
}
