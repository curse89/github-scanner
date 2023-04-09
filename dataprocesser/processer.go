package dataprocesser

import (
	"fmt"
	"github-scanner/ghclient"
)

type ResultData struct {
	account string
	repos   []string
	err     error
}

func ProcessAccounts(accounts []string) {
	var allData []ResultData
	var channel chan ghclient.ReposList = make(chan ghclient.ReposList)
	for _, account := range accounts {
		go getData(channel, account)
	}
	for i := 0; i < len(accounts); i++ {
		var resultData ResultData
		reposList := <-channel
		resultData.account = accounts[i]
		if reposList.Error == nil {
			for _, repo := range reposList.Repos {
				resultData.repos = append(resultData.repos, repo.Name)
			}
		} else {
			resultData.err = fmt.Errorf("ошибка")
		}
		allData = append(allData, resultData)
	}
	printData(allData)
}

func getData(channel chan<- ghclient.ReposList, account string) {
	channel <- ghclient.GetUserReposList(account)
}

func printData(allData []ResultData) {
	fmt.Println(allData)
}
