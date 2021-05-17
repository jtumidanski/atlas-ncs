package inventory

import (
	"atlas-ncs/rest/response"
	"strconv"
)

const (
	ItemAttributesType      string = "com.atlas.cos.rest.attribute.ItemAttributes"
	EquipmentAttributesType string = "com.atlas.cos.rest.attribute.EquipmentAttributes"
	EquipmentStatisticsType string = "com.atlas.cos.rest.attribute.EquipmentStatisticsAttributes"
)

var equipmentIncludes = []response.ConditionalMapperProvider{
	transformItemAttributes,
	transformEquipmentAttributes,
	transformEquipmentStatistics,
}

type InventoryDataContainer struct {
	data     response.DataSegment
	included response.DataSegment
}

type InventoryData struct {
	Id         string              `json:"id"`
	Type       string              `json:"type"`
	Attributes InventoryAttributes `json:"attributes"`
}

type InventoryAttributes struct {
	Type     string `json:"type"`
	Capacity byte   `json:"capacity"`
}

func (c *InventoryDataContainer) UnmarshalJSON(data []byte) error {
	d, i, err := response.UnmarshalRoot(data, response.MapperFunc(EmptyInventoryData), equipmentIncludes...)
	if err != nil {
		return err
	}

	c.data = d
	c.included = i
	return nil
}

func (c *InventoryDataContainer) Data() *InventoryData {
	if len(c.data) >= 1 {
		return c.data[0].(*InventoryData)
	}
	return nil
}

func (c *InventoryDataContainer) DataList() []InventoryData {
	var r = make([]InventoryData, 0)
	for _, x := range c.data {
		r = append(r, *x.(*InventoryData))
	}
	return r
}

func (c *InventoryDataContainer) GetIncludedEquippedItems() []EquipmentData {
	var e = make([]EquipmentData, 0)
	for _, x := range c.included {
		if val, ok := x.(*EquipmentData); ok && val.Attributes.Slot < 0 {
			e = append(e, *val)
		}
	}
	return e
}

func (c *InventoryDataContainer) GetIncludedEquips() []EquipmentData {
	var e = make([]EquipmentData, 0)
	for _, x := range c.included {
		if val, ok := x.(*EquipmentData); ok && val.Attributes.Slot >= 0 {
			e = append(e, *val)
		}
	}
	return e
}

func (c *InventoryDataContainer) GetEquipmentStatistics(id int) *EquipmentStatisticsAttributes {
	for _, x := range c.included {
		if val, ok := x.(*EquipmentStatisticsData); ok {
			eid, err := strconv.Atoi(val.Id)
			if err == nil && eid == id {
				return &val.Attributes
			}
		}
	}
	return nil
}

func (c *InventoryDataContainer) GetIncludedItems() []ItemData {
	var e = make([]ItemData, 0)
	for _, x := range c.included {
		if val, ok := x.(*ItemData); ok && val.Attributes.Slot >= 0 {
			e = append(e, *val)
		}
	}
	return e
}

func EmptyItemData() interface{} {
	return &ItemData{}
}

type ItemData struct {
	Id         string         `json:"id"`
	Type       string         `json:"type"`
	Attributes ItemAttributes `json:"attributes"`
}

type ItemAttributes struct {
	ItemId   uint32 `json:"itemId"`
	Quantity uint16 `json:"quantity"`
	Slot     int16  `json:"slot"`
}

func transformItemAttributes() (string, response.ObjectMapper) {
	return response.UnmarshalData(ItemAttributesType, EmptyItemData)
}

func transformEquipmentAttributes() (string, response.ObjectMapper) {
	return response.UnmarshalData(EquipmentAttributesType, EmptyEquipmentData)
}

func transformEquipmentStatistics() (string, response.ObjectMapper) {
	return response.UnmarshalData(EquipmentStatisticsType, EmptyEquipmentStatisticsData)
}

func EmptyInventoryData() interface{} {
	return &InventoryData{}
}

type EquipmentData struct {
	Id         string              `json:"id"`
	Type       string              `json:"type"`
	Attributes EquipmentAttributes `json:"attributes"`
}

type EquipmentAttributes struct {
	EquipmentId int   `json:"equipmentId"`
	Slot        int16 `json:"slot"`
}

func EmptyEquipmentData() interface{} {
	return &EquipmentData{}
}

type EquipmentStatisticsData struct {
	Id         string                        `json:"id"`
	Type       string                        `json:"type"`
	Attributes EquipmentStatisticsAttributes `json:"attributes"`
}

type EquipmentStatisticsAttributes struct {
	ItemId        uint32 `json:"itemId"`
	Strength      uint16 `json:"strength"`
	Dexterity     uint16 `json:"dexterity"`
	Intelligence  uint16 `json:"intelligence"`
	Luck          uint16 `json:"luck"`
	Hp            uint16 `json:"hp"`
	Mp            uint16 `json:"mp"`
	WeaponAttack  uint16 `json:"weaponAttack"`
	MagicAttack   uint16 `json:"magicAttack"`
	WeaponDefense uint16 `json:"weaponDefense"`
	MagicDefense  uint16 `json:"magicDefense"`
	Accuracy      uint16 `json:"accuracy"`
	Avoidability  uint16 `json:"avoidability"`
	Hands         uint16 `json:"hands"`
	Speed         uint16 `json:"speed"`
	Jump          uint16 `json:"jump"`
	Slots         byte   `json:"slots"`
}

func EmptyEquipmentStatisticsData() interface{} {
	return &EquipmentStatisticsData{}
}