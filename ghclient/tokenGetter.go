package ghclient

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

const fileName = "auth.json"

func loadToken() {
	data, err := os.ReadFile(fileName)
	fmt.Println("Error: ", err)
	if err == nil {
		decoder := json.NewDecoder(strings.NewReader(string(data)))
		err = decoder.Decode(&Auth)
		if err != nil {
			Auth.err = err
		}
		if Auth.Token == "" {
			Auth.err = fmt.Errorf("в файле токен не обнаружен")
		}
	} else {
		Auth.err = err
	}
}
