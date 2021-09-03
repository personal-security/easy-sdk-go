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
var payloadField = "payload"

type RespondApi struct {
	respondHashMap map[string]interface{}
	respondCode    int
}

func (respond *RespondApi) Create(status bool, message string) {
	respond.respondHashMap = map[string]interface{}{}
	respond.respondHashMap[statusField] = status
	respond.respondHashMap[messageField] = message
	respond.respondCode = http.StatusOK
}

// Optional func
func (respond *RespondApi) SetCode(code int) {
	respond.respondCode = code
}

// Optional func
func (respond *RespondApi) SetStatus(status bool) {
	respond.respondHashMap[statusField] = status
}

// Optional func
func (respond *RespondApi) SetMessage(message string) {
	respond.respondHashMap[messageField] = message
}

// Optional func
func (respond *RespondApi) SetPayload(data map[string]interface{}) {
	respond.respondHashMap[payloadField] = data
}

func GenerateApiError(w http.ResponseWriter, message string, data map[string]interface{}, code int) {
	resp := &RespondApi{}
	resp.Create(false, message)
	resp.SetCode(code)
	resp.Respond(w)
}

func GenerateApiErrorJson(message string, data map[string]interface{}) string {
	resp := &RespondApi{}
	resp.Create(false, message)
	return resp.ReturnJson()
}

func GenerateApiRespond(w http.ResponseWriter, status bool, message string, data map[string]interface{}) {
	resp := &RespondApi{}
	resp.Create(status, message)
	if data != nil {
		resp.SetPayload(data)
	}
	resp.Respond(w)
}

func GenerateApiRespondJson(status bool, message string, data map[string]interface{}) string {
	resp := &RespondApi{}
	resp.Create(status, message)
	if data != nil {
		resp.SetPayload(data)
	}
	return resp.ReturnJson()
}

// func (respond *RespondApi) SetItem(item map[string]interface{}) {
// 	payload := map[string]interface{}{}
// 	payload["item"] = item

// 	respond.respondHashMap[payloadField] = payload
// }

// func (respond *RespondApi) SetItems(items map[string]interface{}) {
// 	payload := map[string]interface{}{}
// 	payload["items"] = items

// 	respond.respondHashMap[payloadField] = payload
// }

// func (respond *RespondApi) AddCustomField(fieldName string, items map[string]interface{}) {
// 	if fieldName != "" {
// 		payload := map[string]interface{}{}
// 		payload[fieldName] = items

// 		respond.respondHashMap[payloadField] = payload
// 	}
// }

func (respond *RespondApi) ReturnHashMap() map[string]interface{} {
	return respond.respondHashMap
}

func (respond *RespondApi) ReturnJson() string {
	json, err := json.Marshal(respond.respondHashMap)
	if err != nil {
		//log.Fatal(err)
		return ""
	}
	return string(json)
}

// Return to browser
func (respond *RespondApi) Respond(w http.ResponseWriter) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(respond.respondCode)
	json.NewEncoder(w).Encode(respond.ReturnHashMap())
}
