package schema

type CreateNewUser struct {
	RfidCardId   string `json:"rfid_card_id"`
	UserName     string `json:"user_name"`
	UserPhone    string `json:"user_phone"`
	UserBalance  string `json:"user_balance"`
	UserStatus   string `json:"user_status"`
	UserEmail    string `json:"user_email"`
	UserPassword string `json:"user_password"`
}
