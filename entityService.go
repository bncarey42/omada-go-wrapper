package main

import (
	"encoding/json"
	"fmt"
)

type EntityService[C any, S any, E any] struct {
	client  *OmadaClient
	baseUrl string
}

func NewEnytityService[C any, S any, E any](client *OmadaClient, baseUrl string) EntityService[C, S, E] {
	return EntityService[C, S, E]{client, baseUrl}
}

func (s *EntityService[C, S, E]) GetEntitySummaryList(page int32, pageSize int32) ([]S, error) {
	url := s.client.BuildApiURL(s.baseUrl)
	query := map[string]string{"page": fmt.Sprint(page), "pageSize": fmt.Sprint(pageSize)}
	summeryList, err := HttpRequest[PaginatedApiData[S]]("GET", url, query, nil, s.client)
	if err != nil {
		return nil, fmt.Errorf("error getting list of entity summeries %s", err.Error())
	}
	return summeryList.Data, nil
}

func (s *EntityService[C, S, E]) GetEntityInfo(id string) (*E, error) {
	url := s.client.BuildApiURL(fmt.Sprintf("%s/%s", s.baseUrl, id))
	site, err := HttpRequest[E]("GET", url, nil, nil, s.client)
	if err != nil {
		return nil, fmt.Errorf("error getting entity %s", err.Error())
	}
	return site, nil
}

func (s *EntityService[C, S, E]) CreateNewEntity(createEntity C) error {
	url := s.client.BuildApiURL(s.baseUrl)
	jsonBodyStr, err := json.Marshal(createEntity)
	if err != nil {
		return fmt.Errorf("failed to marshal new site entity :: %s", err.Error())
	}
	body := []byte(jsonBodyStr)

	_, err = HttpRequest[struct{}]("POST", url, nil, body, s.client)
	return err
}

func (s *EntityService[C, S, E]) DeleteEntity(id string) error {
	url := s.client.BuildApiURL(fmt.Sprintf("%s/%s", s.baseUrl, id))
	_, err := HttpRequest[struct{}]("DELETE", url, nil, nil, s.client)
	return err
}
