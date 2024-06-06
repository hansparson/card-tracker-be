package schema

// Success Response
type Response struct {
	Message string      `json:"message" example:"GET_SUCCESS"`
	Status  int         `json:"status" example:"201"`
	Payload interface{} `json:"payload,omitempty"`
	Error   interface{} `json:"error,omitempty"`
}

type Success struct {
	Message string `json:"message" example:"GET_SUCCESS"`
	Status  int    `json:"status" example:"201"`
}

// Failed Response
type Message struct {
	Message string `json:"message"`
}
type Failed struct {
	Message_Action string  `json:"message_action" example:"GET_SUCCESS"`
	Status         int     `json:"status" example:"201"`
	Message_Data   Message `json:"message_data"`
}
