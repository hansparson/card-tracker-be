package services

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"rfid_payment/db"
	dbschema "rfid_payment/db/db_schema"
	"time"

	"github.com/gin-gonic/gin"
)

func formatJSON(data interface{}) string {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err.Error()
	}
	return string(jsonData)
}

func CreateNewUser(ctx *gin.Context) {
	db := db.ConnectGorm()
	currentTime := time.Now()
	// signature := ctx.GetHeader("signature")
	// external_id := ctx.GetHeader("external-id")
	// fmt.Println(signature)
	// fmt.Println(external_id)

	// Read the request body
	body, _ := ioutil.ReadAll(ctx.Request.Body)
	body_string := string(body)
	println(body_string)

	// Unmarshal the JSON payload into a map
	var payload map[string]interface{}
	if err := json.Unmarshal(body, &payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"response_status": "INVALID_PAYLOAD"})
		return
	}

	newUser := dbschema.User{
		RfidCardId:        payload["rfid_card_id"].(string),
		UserName:          payload["user_name"].(string),
		UserPhone:         payload["user_phone"].(string),
		UserBalance:       payload["user_balance"].(float64),
		UserStatus:        "ACTIVE",
		UserEmail:         payload["user_email"].(string),
		UserPassword:      payload["user_password"].(string),
		BalanceUpdateTime: &currentTime,
	}
	result := db.Create(&newUser)

	if result.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"response_status": "USER_ALREADY_EXIST"})
		return
	}

	ctx.JSON(200, gin.H{
		"response_status":  "SUCCESS",
		"response_message": "User Telah Ditambahkan",
		"response_data":    payload})
	return
}

func DeleteUser(ctx *gin.Context) {
	db := db.ConnectGorm()
	// signature := ctx.GetHeader("signature")
	// external_id := ctx.GetHeader("external-id")
	// fmt.Println(signature)
	// fmt.Println(external_id)

	// Read the request body
	body, _ := ioutil.ReadAll(ctx.Request.Body)
	body_string := string(body)
	println(body_string)

	// Unmarshal the JSON payload into a map
	var payload map[string]interface{}
	if err := json.Unmarshal(body, &payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"response_status": "INVALID_PAYLOAD"})
		return
	}

	var real_value string
	var fieldName string

	value1, ok1 := payload["rfid_card_id"].(string)
	value2, ok2 := payload["user_phone"].(string)

	if ok1 && value1 != "" {
		real_value = value1
		fieldName = "rfid"
	} else if ok2 && value2 != "" {
		real_value = value2
		fieldName = "phone"
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{"response_status": "INVALID_PAYLOAD"})
		return
	}

	if fieldName == "rfid" {
		result := db.Where("rfid_card_id = ?", real_value).Delete(&dbschema.User{})
		if result.Error != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"response_status": "DATABASE_ERROR"})
			return
		}

		if result.RowsAffected == 0 {
			ctx.JSON(http.StatusBadRequest, gin.H{"response_status": "USER_NOT_FOUND"})
			return
		}
	} else {
		result := db.Where("user_phone = ?", real_value).Delete(&dbschema.User{})
		if result.Error != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"response_status": "DATABASE_ERROR"})
			return
		}

		if result.RowsAffected == 0 {
			ctx.JSON(http.StatusBadRequest, gin.H{"response_status": "USER_NOT_FOUND"})
			return
		}
	}

	ctx.JSON(http.StatusOK, gin.H{
		"response_status":  "SUCCESS",
		"response_message": "User was successfully deleted",
		"response_data":    payload,
	})
	return

}

func NewAdmin(ctx *gin.Context) {
	db := db.ConnectGorm()
	// signature := ctx.GetHeader("signature")
	// external_id := ctx.GetHeader("external-id")
	// fmt.Println(signature)
	// fmt.Println(external_id)

	// Read the request body
	body, _ := ioutil.ReadAll(ctx.Request.Body)
	body_string := string(body)
	println(body_string)

	// Unmarshal the JSON payload into a map
	var payload map[string]interface{}
	if err := json.Unmarshal(body, &payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"response_status": "INVALID_PAYLOAD"})
		return
	}

	newAdmin := dbschema.Admins{
		AdminID:    payload["admin_id"].(string),
		AdminName:  payload["admin_name"].(string),
		AdminPhone: payload["admin_phone"].(string),
		AdminEmail: payload["admin_email"].(string),
		Password:   payload["password"].(string),
	}

	result := db.Create(&newAdmin)

	if result.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"response_status": "ADMIN_ALREADY_EXIST"})
		return
	}

	ctx.JSON(200, gin.H{
		"response_status":  "SUCCESS",
		"response_message": "User Telah Ditambahkan",
		"response_data":    payload})
	return
}

func DeleteAdmin(ctx *gin.Context) {
	db := db.ConnectGorm()
	// signature := ctx.GetHeader("signature")
	// external_id := ctx.GetHeader("external-id")
	// fmt.Println(signature)
	// fmt.Println(external_id)

	// Read the request body
	body, _ := ioutil.ReadAll(ctx.Request.Body)
	body_string := string(body)
	println(body_string)

	// Unmarshal the JSON payload into a map
	var payload map[string]interface{}
	if err := json.Unmarshal(body, &payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"response_status": "INVALID_PAYLOAD"})
		return
	}

	var real_value string
	var fieldName string

	value1, ok1 := payload["admin_id"].(string)
	value2, ok2 := payload["admin_phone"].(string)

	if ok1 && value1 != "" {
		real_value = value1
		fieldName = "admin_id"
	} else if ok2 && value2 != "" {
		real_value = value2
		fieldName = "admin_phone"
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{"response_status": "INVALID_PAYLOAD"})
		return
	}

	if fieldName == "admin_id" {
		result := db.Where("admin_id = ?", real_value).Delete(&dbschema.Admins{})
		if result.Error != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"response_status": "DATABASE_ERROR"})
			return
		}

		if result.RowsAffected == 0 {
			ctx.JSON(http.StatusBadRequest, gin.H{"response_status": "USER_NOT_FOUND"})
			return
		}
	} else {
		result := db.Where("admin_phone = ?", real_value).Delete(&dbschema.Admins{})
		if result.Error != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"response_status": "DATABASE_ERROR"})
			return
		}

		if result.RowsAffected == 0 {
			ctx.JSON(http.StatusBadRequest, gin.H{"response_status": "USER_NOT_FOUND"})
			return
		}
	}

	ctx.JSON(http.StatusOK, gin.H{
		"response_status":  "SUCCESS",
		"response_message": "Admin was successfully deleted",
		"response_data":    payload,
	})
	return
}
