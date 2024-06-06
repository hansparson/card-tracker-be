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

func TopUp(ctx *gin.Context) {
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

	rfid_card_id := payload["rfid_card_id"].(string)
	admin_name := payload["admin_name"].(string)
	topup_value := payload["topup_value"].(float64)

	var user dbschema.User
	/// Querry data from Databases
	result := db.Where("rfid_card_id = ?", rfid_card_id).First(&user)
	if result.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"response_status": "RFID_CARD_NOT_FOUND"})
		return
	}

	total_balance := user.UserBalance + topup_value

	//// Update Data
	db.Where("rfid_card_id = ?", rfid_card_id).Updates(dbschema.User{
		UserBalance:       total_balance,
		BalanceUpdateTime: &currentTime,
	})

	// Add History Data Topup
	TopHistory := dbschema.HistoryTopup{
		RfidCardId:    rfid_card_id,
		UserName:      user.UserName,
		UserPhone:     user.UserPhone,
		UserBalance:   total_balance,
		TopupAmount:   topup_value,
		BalanceBefore: user.UserBalance,
		TopUpTime:     &currentTime,
		WaitresName:   admin_name,
	}
	db.Create(&TopHistory)

	ctx.JSON(200, gin.H{
		"response_status":  "SUCCESS",
		"response_message": "Saldo Telah Ditambahkan",
		"balance_before":   user.UserBalance,
		"total_balance":    total_balance,
		"rfid_card_id":     rfid_card_id,
		"user_name":        user.UserName,
	})
	return
}

func Transaction(ctx *gin.Context) {
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

	rfid_card_id := payload["rfid_card_id"].(string)
	admin_name := payload["admin_name"].(string)
	topup_value := payload["topup_value"].(float64)

	var user dbschema.User
	/// Querry data from Databases
	result := db.Where("rfid_card_id = ?", rfid_card_id).First(&user)
	if result.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"response_status": "RFID_CARD_NOT_FOUND"})
		return
	}

	total_balance := user.UserBalance + topup_value

	//// Update Data
	db.Where("rfid_card_id = ?", rfid_card_id).Updates(dbschema.User{
		UserBalance:       total_balance,
		BalanceUpdateTime: &currentTime,
	})

	// Add History Data Topup
	TopHistory := dbschema.HistoryTopup{
		RfidCardId:    rfid_card_id,
		UserName:      user.UserName,
		UserPhone:     user.UserPhone,
		UserBalance:   total_balance,
		TopupAmount:   topup_value,
		BalanceBefore: user.UserBalance,
		TopUpTime:     &currentTime,
		WaitresName:   admin_name,
	}
	db.Create(&TopHistory)

	ctx.JSON(200, gin.H{
		"response_status":  "SUCCESS",
		"response_message": "Saldo Telah Ditambahkan",
		"balance_before":   user.UserBalance,
		"total_balance":    total_balance,
		"rfid_card_id":     rfid_card_id,
		"user_name":        user.UserName,
	})
	return
}
