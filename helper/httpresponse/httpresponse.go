package httpresponse

const (
	conflictCode = 7000
	noConflict   = 6000
	updated      = 2002
	created      = 2001
	delted       = 2004
	found        = 2003
)

type BaseDataResponse struct {
	Code        int    `json:"code"`
	MessageKeys string `json:"messageKeys"`
	Message     string `json:"message"`
	Result      struct {
		Data interface{} `json:"data,omitempty"`
	} `json:"result"`
	PageSize   int `json:"pageSize,omitempty"`
	PageNumber int `json:"pageNmber,omitempty"`
}
type BaseDataResponseWithoutData struct {
	Code        int         `json:"code"`
	MessageKeys string      `json:"messageKeys"`
	Message     string      `json:"message"`
	Result      interface{} `json:"result"`
	PageSize    int         `json:"pageSize,omitempty"`
	PageNumber  int         `json:"pageNmber,omitempty"`
}

func NewBaseResponse(data interface{}, message string, httpVerb string) *BaseDataResponse {
	var response BaseDataResponse
	response.Message = message
	switch httpVerb {
	case "GET":
		response.Code = found
		response.MessageKeys = "found successfully"
	case "POST":
		response.Code = created
		response.MessageKeys = "created successfully"
	case "DELETE":
		response.Code = delted
		response.MessageKeys = "deleted successfully"
	case "PUT":
		response.Code = updated
		response.MessageKeys = "updated successfully"
	}

	response.Result.Data = data
	return &response
}

func NewBaseResponseWithoutData(data interface{}, message string, httpVerb string) *BaseDataResponseWithoutData {
	var response BaseDataResponseWithoutData
	response.Message = message
	switch httpVerb {
	case "GET":
		response.Code = found
		response.MessageKeys = "found successfully"
	case "POST":
		response.Code = created
		response.MessageKeys = "created successfully"
	case "DELETE":
		response.Code = delted
		response.MessageKeys = "deleted successfully"
	case "PUT":
		response.Code = updated
		response.MessageKeys = "updated successfully"
	}

	response.Result = data
	return &response
}

// NewNotificationResponse ..
func NewNotificationResponse() *BaseDataResponse {
	var response BaseDataResponse
	response.Code = updated
	response.Message = "status updated successfully"
	response.MessageKeys = "updated successfully"
	return &response
}
