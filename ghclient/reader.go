package ghclient

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func readResponse(response *http.Response) (string, error) {
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return string(body), nil
}

func DeserializeResult(result string, decodedVals DecodedType) {
	reader := strings.NewReader(result)
	decoder := json.NewDecoder(reader)
	err := decoder.Decode(&decodedVals)
	if err != nil {
		decodedVals.SetError(errors.New("ошибка обработки данных"))
	}
}
