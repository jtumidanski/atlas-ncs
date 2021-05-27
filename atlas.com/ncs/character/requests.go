package character

import (
	"atlas-ncs/rest/requests"
	"fmt"
)

const (
	charactersServicePrefix string = "/ms/cos/"
	charactersService              = requests.BaseRequest + charactersServicePrefix
	charactersResource             = charactersService + "characters/"
	charactersById                 = charactersResource + "%d"
	characterItems                 = charactersById + "/items"
	characterItem                  = characterItems + "?itemId=%d"
)

func requestCharacter(characterId uint32) (*dataContainer, error) {
	ar := &dataContainer{}
	err := requests.Get(fmt.Sprintf(charactersById, characterId), ar)
	if err != nil {
		return nil, err
	}
	return ar, nil
}

func requestAllItemsForCharacter(characterId uint32) (*ItemListDataContainer, error) {
	ar := &ItemListDataContainer{}
	err := requests.Get(fmt.Sprintf(characterItems, characterId), ar)
	if err != nil {
		return nil, err
	}
	return ar, nil
}

func requestItemsForCharacter(characterId uint32, itemId uint32) (*ItemListDataContainer, error) {
	ar := &ItemListDataContainer{}
	err := requests.Get(fmt.Sprintf(characterItem, characterId, itemId), ar)
	if err != nil {
		return nil, err
	}
	return ar, nil
}
