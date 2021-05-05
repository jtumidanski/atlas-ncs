package topic

import (
	"atlas-ncs/rest/requests"
	"atlas-ncs/retry"
	"fmt"
)

const (
	topicsServicePrefix string = "/ms/tds/"
	topicsService              = requests.BaseRequest + topicsServicePrefix
	topicById                  = topicsService + "topics/%s"
)

func RequestTopic(topic string) (*Model, error) {
	td := &DataContainer{}

	var get = func(attempt int) (bool, error) {
		err := requests.Get(fmt.Sprintf(topicById, topic), td)
		if err != nil {
			return true, err
		}
		return false, nil
	}

	err := retry.Try(get, 10)
	if err != nil {
		return nil, err
	}

	return &Model{name: td.Data.Attributes.Name}, nil
}
