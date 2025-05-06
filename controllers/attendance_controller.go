package controllers

import (
	"net/http"
	"time"
	"westay-go/config"
	"westay-go/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func StoreAttendance(c *gin.Context) {
	var existingAttendance models.Attendance
	var attendace models.Attendance
	if err := c.ShouldBindJSON(&attendace); err != nil {
		c.JSON(http.StatusBadRequest, models.AttendanceError{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}
	userAll, _ := c.Get("user")
	var userID uint = userAll.(models.UserResponse).ID
	//check apa sudah ada yang checkin
	config.DB.Where("user_id = ? and Date(ci_date) = Date(NOW())", userID).First(&existingAttendance)

	if existingAttendance.ID == uuid.Nil {
		existingAttendance.CiDate = time.Now().Local().UTC()
		existingAttendance.UserID = userAll.(models.UserResponse).ID
		config.DB.Create(&existingAttendance)
		c.JSON(http.StatusCreated, models.AttendanceResponse{
			Status:       http.StatusCreated,
			Message:      "Attendance Store",
			AttendanceID: existingAttendance.ID.String(),
		})
		return
	}

	// Jika sudah check-out hari ini, kembalikan error
	if !existingAttendance.CoDate.IsZero() {
		c.JSON(http.StatusForbidden, models.AttendanceError{
			Status:  http.StatusForbidden,
			Message: "You have already checked in and out today",
		})
		return
	}

	// If attendance exists, update check-out time
	updates := map[string]interface{}{
		"CoDate": time.Now().UTC(),
		"CoType": existingAttendance.CiType,
		"CoFoto": existingAttendance.CiFoto,
	}

	if err := config.DB.Model(&existingAttendance).Updates(updates).Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.AttendanceError{
			Status:  http.StatusInternalServerError,
			Message: "Failed to update attendance: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusAccepted, models.AttendanceResponse{
		Status:       http.StatusAccepted,
		Message:      "Attendance check-out updated successfully",
		AttendanceID: existingAttendance.ID.String(),
	})

}
