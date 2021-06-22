package helpers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/zeshanwd/go-rest-api/src/models"
)

func DecodeBody(req *http.Request) (models.Todo, bool) {
	var todo models.Todo
	erro := json.NewDecoder(req.Body).Decode(&todo)
	if erro != nil {
		return models.Todo{}, false
	}

	return todo, true
}

func IsValidDescription(description string) bool {
	desc := strings.TrimSpace(description)
	if len(desc) == 0 {
		return false
	}

	return true
}
