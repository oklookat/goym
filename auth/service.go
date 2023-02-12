package auth

import (
	"fmt"
	"os"

	"github.com/oklog/ulid/v2"
)

func wrapErrStr(err string) error {
	return fmt.Errorf(_errPrefix+"%s", err)
}

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
