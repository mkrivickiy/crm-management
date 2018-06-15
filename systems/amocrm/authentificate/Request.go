package authentificate

import "github.com/mkrivickiy/crm-management/systems/amocrm"

type Request struct {
	amocrm.Request
}

func NewRequest() *Request {
	request := Request{}
	request.Host = "mkrivickiy.amocrm.ru"
	request.Protocol = "https"
	request.MethodURL = "private/api/auth.php"
	request.RequestType = "json"
	request.ContentType = "application/json; charset=utf-8"
	return &request
}
