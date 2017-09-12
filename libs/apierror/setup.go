package apierror

import "net/http"

var errorDictionary map[string]int

// Setup APIError with error dictionary (mapping error code with http status)
func Setup(dictionary map[string]int) {
	if dictionary != nil {
		errorDictionary = dictionary
		errorDictionary[CodeInternalServerError] = http.StatusInternalServerError
	}
}
