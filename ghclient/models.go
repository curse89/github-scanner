package ghclient

import (
	"encoding/json"
	"errors"
)

type DecodedType interface {
	UnmarshalJSON(data []byte) (err error)
	SetError(err error)
}

type ReposList struct {
	Repos []Repos
	Error error
}

type Repos struct {
	Name       string
	Visibility string
	Error      error
}

type ReleasesList struct {
	Versions []string
	Error    error
}

func NewErrorReposList(err error) (rList ReposList) {
	rList.Error = err

	return
}

func NewErrorReleasesList(err error) (releasesList ReleasesList) {
	releasesList.Error = err

	return
}

func (rList *ReposList) UnmarshalJSON(data []byte) (err error) {
	mdata := []Repos{}
	err = json.Unmarshal(data, &mdata)
	if err != nil {
		rList.Error = errors.New("ошибка десереализации данных")
	}
	rList.Repos = mdata

	return
}

func (releasesList *ReleasesList) UnmarshalJSON(data []byte) (err error) {
	mdata := []map[string]interface{}{}
	err = json.Unmarshal(data, &mdata)
	if err != nil {
		releasesList.Error = errors.New("ошибка десереализации данных")
	}
	for _, decValue := range mdata {
		if name, ok := decValue["name"].(string); ok {
			releasesList.Versions = append(releasesList.Versions, name)
		}
	}

	return
}

func (repo *Repos) UnmarshalJSON(data []byte) (err error) {
	mdata := map[string]interface{}{}
	err = json.Unmarshal(data, &mdata)
	if err == nil {
		if name, ok := mdata["name"].(string); ok {
			repo.Name = name
		}
		if visibility, ok := mdata["visibility"].(string); ok {
			repo.Visibility = visibility
		}
	}

	return
}

func (rList *ReposList) SetError(err error) {
	rList.Error = err
}

func (releasesList *ReleasesList) SetError(err error) {
	releasesList.Error = err
}
