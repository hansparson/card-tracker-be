package dbschema

import "time"

type AlarmBlastMode string
type GpsDevice string

const (
	RegisterMode AlarmBlastMode = "register"
	NearMode     AlarmBlastMode = "near"

	CardDevice  GpsDevice = "card"
	PhoneDevice GpsDevice = "phone"
)

type UserTracker struct {
	ID              int            `json:"id" gorm:"primaryKey;autoIncrement"`
	WalletNumber    string         `json:"wallet_number" gorm:"unique"`
	FirstName       string         `json:"first_name"`
	Surename        string         `json:"sure_name"`
	Email           string         `json:"email"`
	MobileNumber    string         `json:"mobile_number"`
	DateOfBirth     time.Time      `json:"dob"`
	Nationality     string         `json:"nationality"`
	NIK             string         `json:"nik"`
	Gender          string         `json:"gender"`
	BloodType       string         `json:"blood_type"`
	HistoryOfIllnes string         `json:"illness"`
	Alergic         string         `json:"alergic"`
	AlarmBlastMode  AlarmBlastMode `json:"alarm_blast_mode"`
	GpsDevice       GpsDevice      `json:"gps_device"`
	Visibility      bool           `json:"visibility"`
	RegistyNumber   string         `json:"registry_number"`
}
