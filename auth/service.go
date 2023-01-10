package auth

import (
	"os"

	"github.com/oklog/ulid/v2"
)

// Получить имя устройства, и дописать к нему "/goym".
//
// Дописываем, чтобы в списке входов в аккаунт пользователя
// было видно, что вход выполнен через API.
func getHostname() (name string, err error) {
	if name, err = os.Hostname(); err != nil {
		return
	}
	name += "/goym"
	return
}

// Генерирует ULID.
func generateUlid() string {
	return ulid.Make().String()
}
