package amocrm

type CustomField struct {
	ID     int                `json:"id"`     //Уникальный идентификатор заполняемого дополнительного поля
	Values []CustomFieldValue `json:"values"` //Массив значений
}
