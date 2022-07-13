package httperrors

type JsonErrorMessage struct {
	Error   bool   `json:"error"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func ProvidedIdIsnNotAnInteger() JsonErrorMessage {
	return JsonErrorMessage{
		Error:   true,
		Code:    400,
		Message: "Provided ID is not an integer.",
	}
}

func CarNotFound() JsonErrorMessage {
	return JsonErrorMessage{
		Error:   true,
		Code:    404,
		Message: "Car with given ID not found.",
	}
}
