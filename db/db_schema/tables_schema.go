package dbschema

import "time"

type UserLogStatus string
type TransactionLogStatus string

const (
	ACTIVE  UserLogStatus        = "ACTIVE"
	LOCK    UserLogStatus        = "LOCK"
	PLAYING UserLogStatus        = "PLAYING"
	TOPUP   TransactionLogStatus = "TOPUP"
	CASHOUT TransactionLogStatus = "CASHOUT"
	REFUND  TransactionLogStatus = "REFUND"
)

type User struct {
	ID                int        `json:"id"`
	RfidCardId        string     `json:"rfid_card_id" gorm:"unique"`
	UserName          string     `json:"user_name"`
	UserPhone         string     `json:"user_phone" gorm:"unique"`
	UserBalance       float64    `json:"user_balance"`
	UserStatus        string     `json:"user_status"`
	UserEmail         string     `json:"user_email"`
	UserPassword      string     `json:"user_password"`
	BalanceUpdateTime *time.Time `json:"balance_update_time"`
}

type HistoryTransactions struct {
	ID                int        `json:"id"`
	RfidCardId        string     `json:"rfid_card_id"`
	UserName          string     `json:"user_name"`
	UserPhone         string     `json:"user_phone"`
	UserBalance       float64    `json:"user_balance"`
	TransactionAmount float64    `json:"transaction_amount"`
	BalanceBefore     float64    `json:"balance_before"`
	TransactionStatus string     `json:"transaction_status"`
	TransactionTime   *time.Time `json:"transaction_time"`
	PlaygroundName    string     `json:"playground_name"`
}

type HistoryTopup struct {
	ID             int        `json:"id"`
	RfidCardId     string     `json:"rfid_card_id"`
	UserName       string     `json:"user_name"`
	UserPhone      string     `json:"user_phone"`
	UserBalance    float64    `json:"user_balance"`
	TopupAmount    float64    `json:"topup_amount"`
	BalanceBefore  float64    `json:"balance_before"`
	TopUpTime      *time.Time `json:"topup_time"`
	PlaygroundName string     `json:"playground_name"`
	WaitresName    string     `json:"waitres_name"`
}

type Admins struct {
	ID         int    `json:"id"`
	AdminID    string `json:"admins_id" gorm:"unique"`
	AdminName  string `json:"user_name"`
	AdminPhone string `json:"user_phone" gorm:"unique"`
	AdminEmail string `json:"user_email"`
	Password   string `json:"password"`
}
