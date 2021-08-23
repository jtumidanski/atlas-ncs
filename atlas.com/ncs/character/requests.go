package character

import (
	"atlas-ncs/rest/requests"
	"fmt"
	"github.com/sirupsen/logrus"
)

const (
	charactersServicePrefix string = "/ms/cos/"
	charactersService              = requests.BaseRequest + charactersServicePrefix
	charactersResource             = charactersService + "characters/"
	charactersById                 = charactersResource + "%d"
	characterItems                 = charactersById + "/items"
	characterItem                  = characterItems + "?itemId=%d"
)

func requestCharacter(l logrus.FieldLogger) func(characterId uint32) (*dataContainer, error) {
	return func(characterId uint32) (*dataContainer, error) {
		ar := &dataContainer{}
		err := requests.Get(l)(fmt.Sprintf(charactersById, characterId), ar)
		if err != nil {
			return nil, err
		}
		return ar, nil
	}
}

func requestAllItemsForCharacter(l logrus.FieldLogger) func(characterId uint32) (*ItemListDataContainer, error) {
	return func(characterId uint32) (*ItemListDataContainer, error) {
		ar := &ItemListDataContainer{}
		err := requests.Get(l)(fmt.Sprintf(characterItems, characterId), ar)
		if err != nil {
			return nil, err
		}
		return ar, nil
	}
}

func requestItemsForCharacter(l logrus.FieldLogger) func(characterId uint32, itemId uint32) (*ItemListDataContainer, error) {
	return func(characterId uint32, itemId uint32) (*ItemListDataContainer, error) {
	ar := &ItemListDataContainer{}
	err := requests.Get(l)(fmt.Sprintf(characterItem, characterId, itemId), ar)
	if err != nil {
		return nil, err
	}
	return ar, nil
	}
}
