package amocrm

type CustomFieldValue struct {
	Value   string `json:"value"`   //Значение дополнительного поля
	Subtype string `json:"subtype"` //Тип изменяемого элемента дополнительного поля типа "адрес". Внимание, все типы, которые не были переданы, будут стёрты
}
