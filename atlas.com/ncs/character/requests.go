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

func requestCharacter(characterId uint32) requests.Request[attributes] {
	return requests.MakeGetRequest[attributes](fmt.Sprintf(charactersById, characterId))
}

func requestAllItemsForCharacter(characterId uint32) requests.Request[itemAttributes] {
	return requests.MakeGetRequest[itemAttributes](fmt.Sprintf(characterItems, characterId))
}

func requestItemsForCharacter(characterId uint32, itemId uint32) requests.Request[itemAttributes] {
	return requests.MakeGetRequest[itemAttributes](fmt.Sprintf(characterItem, characterId, itemId))
}
