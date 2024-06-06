package logger

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func GenerateAPICallID() string {
	// Get current Unix timestamp
	nowTime := time.Now().Unix()

	// Get current time components
	current := time.Now()
	hour := current.Hour()
	minute := current.Minute()
	second := current.Second()

	// Create the start ID by concatenating timestamp and time components
	startID := strconv.FormatInt(nowTime, 10) +
		strconv.Itoa(hour) +
		strconv.Itoa(minute) +
		strconv.Itoa(second)

	// Generate a random number between 1 and 10,000,000
	randomNum := rand.Intn(10000000) + 1

	// Create the invoice code
	invoiceCode := fmt.Sprintf("API_CALL_%s_%d", startID, randomNum)

	return invoiceCode
}

func formatJSON(data interface{}) string {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err.Error()
	}
	return string(jsonData)
}

func PrintLogInfo(ctx *gin.Context, title string, apicall string, host string, endpoint string, response ...string) {
	// Extract and print headers
	headersJSON := formatJSON(ctx.Request.Header)

	// Extract and print payload (request body)
	bodyBytes, _ := ioutil.ReadAll(ctx.Request.Body)
	bodyString := string(bodyBytes)

	// Extract and print query parameters
	queryParams := ctx.Request.URL.Query()
	paramsJSON := formatJSON(queryParams)

	var responseValue string
	if len(response) > 0 {
		responseValue = response[0]
	} else {
		responseValue = ""
	}

	options := DataLogger{
		apicall:  apicall,
		state:    title,
		body:     bodyString,
		params:   paramsJSON,
		headers:  headersJSON,
		host:     host,
		endpoint: endpoint,
		response: responseValue,
	}

	MyLogger(options)
}

type DataLogger struct {
	apicall  string
	state    string
	body     string
	params   string
	headers  string
	host     string
	endpoint string
	response string
}

func MyLogger(options DataLogger) {
	string_logger := ""

	currentTime := time.Now()
	datetimeStr := currentTime.Format("2006-01-02 15:04:05")
	string_logger = fmt.Sprintf("[%s]", datetimeStr)

	if options.apicall != "" {
		string_logger = string_logger + "||" + options.apicall
	}
	fmt.Println(string_logger + "||" + options.state)

	if options.host != "" {
		string_logger = string_logger + "||" + options.host
	}

	if options.endpoint != "" {
		string_logger = string_logger + "||" + options.endpoint
	}

	if options.headers != "" {
		string_logger = string_logger + "||HEADERS=" + options.headers
	}

	if options.body != "" {
		string_logger = string_logger + "||PAYLOADS=" + formatJSON(json.RawMessage(options.body))
	}

	if options.params != "" {
		string_logger = string_logger + "||PARAMS=" + options.params
	}

	if options.response != "" {
		string_logger = string_logger + "||RESPONSE=" + options.response
	}

	fmt.Println(string_logger)

	if options.state == "RESPONSE_END" {
		fmt.Println("")
	}

}
