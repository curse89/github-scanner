package ghclient

import (
	"fmt"
	"io"
	"net/http"
)

const BASE_PATH = "https://api.github.com/"

type token struct {
	Token string
	err   error
}

var Auth token

func init() {
	loadToken()
}

func GetUserReposList(userName string) ReposList {
	path := fmt.Sprintf("users/%s/repos", userName)
	result, err := makeGetRequest(path)
	if err != nil {
		return NewErrorReposList(err)
	}
	var reposList ReposList
	DeserializeResult(result, &reposList)

	return reposList
}

func GetUserReleasesList(owner string, repo string) ReleasesList {
	path := fmt.Sprintf("repos/%s/%s/releases", owner, repo)
	result, err := makeGetRequest(path)
	if err != nil {
		return NewErrorReleasesList(err)
	}
	var releasesList ReleasesList
	DeserializeResult(result, &releasesList)

	return releasesList
}

func makeGetRequest(path string) (string, error) {
	url := fmt.Sprintf("%s%s", BASE_PATH, path)
	request, _ := http.NewRequest("GET", url, nil)
	if Auth.err == nil {
		request.Header.Add("Authorization", fmt.Sprintf("token %s", Auth.Token))
	}
	client := &http.Client{}
	response, err := client.Do(request)

	if err != nil || response.StatusCode != 200 {
		fmt.Println(err)
		test, _ := io.ReadAll(response.Body)
		fmt.Println(test)
		return "", fmt.Errorf("%v: ошибка получения данных", path)
	}
	defer response.Body.Close()

	return readResponse(response)
}
