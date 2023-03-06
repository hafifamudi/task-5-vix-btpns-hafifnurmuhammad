package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	photoRes "github.com/hafifamudi/task-5-vix-btpns-hafifnurmuhammad/app/photo"
	"github.com/hafifamudi/task-5-vix-btpns-hafifnurmuhammad/helpers"
	"github.com/hafifamudi/task-5-vix-btpns-hafifnurmuhammad/models"
	"gorm.io/gorm"
)

type photoController struct {
	db *gorm.DB
}

func NewPhotoController(db *gorm.DB) *photoController {
	return &photoController{db}
}

func (h *photoController) Get(c *gin.Context) {
	var userPhoto models.Photo
	err := h.db.Preload("User").Find(&userPhoto).Error

	if err != nil {
		response := helpers.ApiResponse(http.StatusBadRequest, "error", nil, "Failed to Get Your Photo")
		c.JSON(http.StatusBadRequest, response)
		return
	}

	if userPhoto.PhotoURL == "" {
		response := helpers.ApiResponse(http.StatusBadRequest, "error", nil, "Please Upload Your Photo")
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := photoRes.FormatPhoto(&userPhoto, "")
	response := helpers.ApiResponse(http.StatusOK, "success", formatter, "Successfully Fetch User Photo")
	c.JSON(http.StatusOK, response)
}

func (h *photoController) Create(c *gin.Context) {
	var userPhoto models.Photo
	var countPhoto int64
	currentUser := c.MustGet("currentUser").(models.User)

	h.db.Model(&userPhoto).Where("user_id = ?", currentUser.ID).Count(&countPhoto)
	if countPhoto > 0 {
		data := gin.H{
			"is_uploaded": false,
		}
		response := helpers.ApiResponse(http.StatusBadRequest, "error", data, "You Already Have a Photo")
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var input models.Photo
	err := c.ShouldBind(&input)
	if err != nil {
		errors := helpers.FormatValidationError(err)
		errorMessages := gin.H{"errors": errors}

		response := helpers.ApiResponse(http.StatusUnprocessableEntity, "error", errorMessages, "Failed to Upload User Photo")
		c.JSON(http.StatusBadRequest, response)
		return
	}

	file, err := c.FormFile("photo_profile")
	if err != nil {
		errors := helpers.FormatValidationError(err)
		errorMessages := gin.H{"errors": errors}

		response := helpers.ApiResponse(http.StatusUnprocessableEntity, "error", errorMessages, "Failed to Upload User Photo")
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	extension := file.Filename
	path := "static/images/" + uuid.New().String() + extension

	err = c.SaveUploadedFile(file, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}

		response := helpers.ApiResponse(http.StatusUnprocessableEntity, "error", data, "Failed to Upload User Photo")
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	h.InsertPhoto(input, path, currentUser.ID)

	data := gin.H{"is_uploaded": true}
	response := helpers.ApiResponse(http.StatusOK, "success", data, "Photo Profile Successfully Uploaded")
	c.JSON(http.StatusOK, response)
}

func (h *photoController) InsertPhoto(userPhoto models.Photo, fileLocation string, currUserID int) error {
	savePhoto := models.Photo{
		UserID:   currUserID,
		Title:    userPhoto.Title,
		Caption:  userPhoto.Caption,
		PhotoURL: fileLocation,
	}

	err := h.db.Debug().Create(&savePhoto).Error
	if err != nil {
		return err
	}
	return nil
}

func (h *photoController) Update(c *gin.Context) {
	var userPhoto models.Photo
	currentUser := c.MustGet("currentUser").(models.User)

	err := h.db.Where("user_id = ?", currentUser.ID).Find(&userPhoto).Error
	if err != nil {
		response := helpers.ApiResponse(http.StatusBadRequest, "error", err, "Photo Profile Failed to Update")
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var input models.Photo
	err = c.ShouldBind(&input)
	if err != nil {
		response := helpers.ApiResponse(http.StatusBadRequest, "error", err, "Photo Profile Failed to Update")
		c.JSON(http.StatusBadRequest, response)
		return
	}

	file, err := c.FormFile("update_profile")
	if err != nil {
		data := gin.H{"is_uploaded": false}

		response := helpers.ApiResponse(http.StatusUnprocessableEntity, "error", data, "Failed to Update User Photo")
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	extension := file.Filename
	path := "static/images/" + uuid.New().String() + extension

	err = c.SaveUploadedFile(file, path)
	if err != nil {
		response := helpers.ApiResponse(http.StatusBadRequest, "error", err, "Photo Profile Failed to Upload")
		c.JSON(http.StatusBadRequest, response)
		return
	}

	h.UpdatePhoto(input, &userPhoto, path)

	data := photoRes.FormatPhoto(&userPhoto, "regular")
	response := helpers.ApiResponse(http.StatusOK, "success", data, "Photo Profile Successfully Updated")
	c.JSON(http.StatusOK, response)
}

func (h *photoController) UpdatePhoto(oldPhoto models.Photo, newPhoto *models.Photo, path string) error {
	newPhoto.Title = oldPhoto.Title
	newPhoto.Caption = oldPhoto.Caption
	newPhoto.PhotoURL = path

	err := h.db.Save(&newPhoto).Error
	if err != nil {
		return err
	}

	return nil
}

func (h *photoController) Delete(c *gin.Context) {
	var userPhoto models.Photo
	currentUser := c.MustGet("currentUser").(models.User)

	err := h.db.Where("user_id = ?", currentUser.ID).Delete(&userPhoto).Error
	if err != nil {
		data := gin.H{
			"is_deleted": false,
		}

		response := helpers.ApiResponse(http.StatusBadRequest, "error", data, "Failed to delete user photo")
		c.JSON(http.StatusBadRequest, response)
		return
	}

	data := gin.H{
		"is_deleted": true,
	}

	response := helpers.ApiResponse(http.StatusOK, "success", data, "User Photo Successfully Deleted")
	c.JSON(http.StatusOK, response)
}
