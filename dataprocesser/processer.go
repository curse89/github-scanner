package dataprocesser

import (
	"fmt"
	"github-scanner/ghclient"
)

const LineBreak = "\r\n"

func ProcessAccounts(accounts []string) {
	var channel chan DataResult = make(chan DataResult)
	for i := 0; i < len(accounts); i++ {
		go getData(channel, accounts[i])
		resultData := <-channel
		printData(resultData)
	}
}

func getData(channel chan<- DataResult, account string) {
	var result DataResult
	result.account = account
	repos := ghclient.GetUserReposList(account, true)
	if repos.Error != nil {
		result.err = repos.Error
	} else {
		for _, repo := range repos.Repos {
			var repoResult RepoResult
			repoResult.name = repo.Name
			versions := ghclient.GetUserReleasesList(account, repo.Name)

			if versions.Error != nil {
				repoResult.versionsErr = versions.Error
			} else {
				repoResult.versions = versions.Versions
			}

			result.repos = append(result.repos, repoResult)
		}
	}

	channel <- result
}

func printData(result DataResult) {
	fmt.Println(result.account)
	fmt.Println("#############")
	if result.err == nil {
		if len(result.repos) != 0 {
			for _, repoData := range result.repos {
				fmt.Println(LineBreak)
				fmt.Println(repoData.name)
				fmt.Println("---------------")
				if repoData.versionsErr == nil {
					if len(repoData.versions) != 0 {
						for _, version := range repoData.versions {
							fmt.Println(version)
						}
					} else {
						fmt.Println("версионирование отсутствует")
					}
				} else {
					fmt.Println("ошибка в работе с версиями ", repoData.versionsErr)
				}
			}
		} else {
			fmt.Println("нет активных репозиториев")
		}
	} else {
		fmt.Println("ошибка в работе с репозиториями ", result.err)
	}
	fmt.Println(LineBreak, LineBreak)
}
