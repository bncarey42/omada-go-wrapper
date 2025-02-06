package main

import (
	"encoding/json"
	"fmt"
)

type OmadaService[E any] interface {
	Create(createEntity E) error
	GetList(page int32, pageSize int32) ([]E, error)
	GetOne(id string) (*E, error)
	Update() error // TODO
	Delete(id string) error
}

type EntityService[E any] struct {
	client  *OmadaClient
	baseUrl string
}

func NewOmadaService[E any](client *OmadaClient, baseUrl string) OmadaService[E] {
	return EntityService[E]{client, baseUrl}
}

func (s EntityService[E]) GetList(page int32, pageSize int32) ([]E, error) {
	url := s.client.BuildApiURL(s.baseUrl)
	query := map[string]string{"page": fmt.Sprint(page), "pageSize": fmt.Sprint(pageSize)}
	summeryList, err := HttpRequest[PaginatedApiData[E]]("GET", url, query, nil, s.client)
	if err != nil {
		return nil, fmt.Errorf("error getting list of entity summeries %s", err.Error())
	}
	return summeryList.Data, nil
}

func (s EntityService[E]) GetOne(id string) (*E, error) {
	url := s.client.BuildApiURL(fmt.Sprintf("%s/%s", s.baseUrl, id))
	site, err := HttpRequest[E]("GET", url, nil, nil, s.client)
	if err != nil {
		return nil, fmt.Errorf("error getting entity %s", err.Error())
	}
	return site, nil
}

func (s EntityService[E]) Create(createEntity E) error {
	url := s.client.BuildApiURL(s.baseUrl)
	jsonBodyStr, err := json.Marshal(createEntity)
	if err != nil {
		return fmt.Errorf("failed to marshal new site entity :: %s", err.Error())
	}
	body := []byte(jsonBodyStr)

	_, err = HttpRequest[struct{}]("POST", url, nil, body, s.client)
	return err
}

func (s EntityService[E]) Delete(id string) error {
	url := s.client.BuildApiURL(fmt.Sprintf("%s/%s", s.baseUrl, id))
	_, err := HttpRequest[struct{}]("DELETE", url, nil, nil, s.client)
	return err
}

func (s EntityService[E]) Update() error {
	return nil
}
