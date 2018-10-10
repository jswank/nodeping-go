package nodeping

import "fmt"

type ChecksService service

type Check struct {
	Id     string `json:"_id"`
	Enable string `json:"enable"`
	Label  string `json:"label"`
	State  int64  `json:"state"`
}

type CheckMap map[string]*Check

func (s *ChecksService) List() ([]*Check, *Response, error) {
	u, err := addOptions("checks", nil)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	check_map := new(CheckMap)
	resp, err := s.client.Do(req, check_map)
	if err != nil {
		return nil, resp, err
	}

	var checks []*Check

	for _, c := range *check_map {
		checks = append(checks, c)
	}

	return checks, resp, err
}

func (s *ChecksService) Get(id string) (*Check, *Response, error) {
	u := fmt.Sprintf("checks/%v", id)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}
	check := new(Check)
	resp, err := s.client.Do(req, check)
	if err != nil {
		return nil, resp, err
	}

	return check, resp, err
}
