package models

import "net/http"

var Response401 = Error{
	Status:  http.StatusUnauthorized,
	Message: "No authorized",
}
