package schema

import (
	"encoding/json"
	"net/url"

	"github.com/google/go-querystring/query"
)

// Преобразовать struct (НЕ указатель на struct) в url.Values.
//
// Доступно для структур, название которых заканчивается на "Params" и "Body".
//
// Но не всегда. В некоторых структурах есть дополнительные методы. Читайте доки (c).
//
// После получения Values можно сделать Encode(), и отправить GET или POST (request body).
func ParamsToValues(s any) (url.Values, error) {
	return query.Values(s)
}

// Приводит ID сущности к int64.
//
// Применяется, когда ID может быть как int, так и string.
//
// 1. Поле ID в структуре должно быть с типом int64 и тегом `json: "-"`.
//
// 2. Разместите эту функцию в UnmarshalJSON нужной структуры.
//
// ider должен находится в UnmarshalJSON() в виде var,
// сама функция должна создавать type alias(1) на структуру, делать демаршал в неё,
// результат копировать в текущую структуру, выставить поле ID.
//
// (1)type alias нужен, чтобы не было stack overflow, из-за бесконечного вызова UnmarshalJSON().
func unmarshalID(ider func(id ID, data []byte) error, data []byte) error {
	if len(data) == 0 || ider == nil {
		return nil
	}

	// если ID int: окей
	idInt := &struct {
		ID ID `json:"id"`
	}{}
	if err := json.Unmarshal(data, idInt); err == nil {
		return ider(idInt.ID, data)
	}

	// если ID строка: конвертируем в int
	idString := &struct {
		ID string `json:"id"`
	}{}
	if err := json.Unmarshal(data, idString); err != nil {
		// ID не int, и не строка.
		return err
	}
	var converted ID = 0
	err := converted.FromString(idString.ID)
	if err != nil {
		return err
	}
	return ider(converted, data)
}
