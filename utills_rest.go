package easysdk

import (
	"encoding/json"
	"net/http"
)

// Message show message in request
func Message(status bool, message string) map[string]interface{} {
	return map[string]interface{}{"status": status, "message": message}
}

// Respond add headers to respond
func Respond(w http.ResponseWriter, data map[string]interface{}) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

// new api

var statusField = "status"
var messageField = "message"

type RespondAnswer struct {
	respondAnswer map[string]interface{}
}

func (respond *RespondAnswer) Create(status bool, message string) {
	respond.respondAnswer = map[string]interface{}{}
	respond.respondAnswer[statusField] = status
	respond.respondAnswer[messageField] = message
}

// Optional func
func (respond *RespondAnswer) SetStatus(status bool) {
	respond.respondAnswer[statusField] = status
}

// Optional func
func (respond *RespondAnswer) SetMessage(message string) {
	respond.respondAnswer[messageField] = message
}

func (respond *RespondAnswer) AddItem(item map[string]interface{}) {
	respond.respondAnswer["item"] = item
}

func (respond *RespondAnswer) AddItems(items map[string]interface{}) {
	respond.respondAnswer["items"] = items
}

func (respond *RespondAnswer) AddCustomField(fieldName string, items map[string]interface{}) {
	if fieldName != "" {
		respond.respondAnswer[fieldName] = items
	}
}

func (respond *RespondAnswer) Return() map[string]interface{} {
	return respond.respondAnswer
}

// Return to browser
func (respond *RespondAnswer) Respond(w http.ResponseWriter) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(respond.Return())
}
