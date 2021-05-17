package character

import (
	"atlas-ncs/character/inventory"
	"atlas-ncs/rest/requests"
	"fmt"
)

const (
	charactersServicePrefix string = "/ms/cos/"
	charactersService              = requests.BaseRequest + charactersServicePrefix
	charactersResource             = charactersService + "characters/"
	charactersById                 = charactersResource + "%d"
	inventoryResource              = charactersById + "/inventories?type=%s&include=inventoryItems,equipmentStatistics"
	inventoryByItemId              = inventoryResource + "&itemId=%d"
)

func requestCharacter(characterId uint32) (*dataContainer, error) {
	ar := &dataContainer{}
	err := requests.Get(fmt.Sprintf(charactersById, characterId), ar)
	if err != nil {
		return nil, err
	}
	return ar, nil
}

func requestItemsForCharacter(characterId uint32, inventoryType string, itemId uint32) (*inventory.InventoryDataContainer, error) {
	ar := &inventory.InventoryDataContainer{}
	err := requests.Get(fmt.Sprintf(inventoryByItemId, characterId, inventoryType, itemId), ar)
	if err != nil {
		return nil, err
	}
	return ar, nil
}