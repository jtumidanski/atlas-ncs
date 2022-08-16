package portal

import (
	"atlas-ncs/model"
	"atlas-ncs/rest/requests"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
	"strconv"
)

func getId(m Model) (uint32, error) {
	return m.Id(), nil
}

func ByNamePortalIdProvider(l logrus.FieldLogger, span opentracing.Span) func(mapId uint32, name string) model.IdProvider[uint32] {
	return func(mapId uint32, name string) model.IdProvider[uint32] {
		return model.ProviderToIdProviderAdapter(ByNameProvider(l, span)(mapId, name), getId)
	}
}

func RandomPortalProvider(l logrus.FieldLogger, span opentracing.Span) func(mapId uint32) model.Provider[Model] {
	return func(mapId uint32) model.Provider[Model] {
		return model.SliceProviderToProviderAdapter(InMapProvider(l, span)(mapId), model.RandomPreciselyOneFilter[Model])
	}
}

func RandomPortalIdProvider(l logrus.FieldLogger, span opentracing.Span) func(mapId uint32) model.IdProvider[uint32] {
	return func(mapId uint32) model.IdProvider[uint32] {
		return model.ProviderToIdProviderAdapter(RandomPortalProvider(l, span)(mapId), getId)
	}
}

func InMapProvider(l logrus.FieldLogger, span opentracing.Span) func(mapId uint32) model.SliceProvider[Model] {
	return func(mapId uint32) model.SliceProvider[Model] {
		return requests.SliceProvider[attributes, Model](l, span)(requestAll(mapId), makePortal)
	}
}

func ByNameProvider(l logrus.FieldLogger, span opentracing.Span) func(mapId uint32, portalName string) model.Provider[Model] {
	return func(mapId uint32, portalName string) model.Provider[Model] {
		return requests.Provider[attributes, Model](l, span)(requestByName(mapId, portalName), makePortal)
	}
}

func makePortal(body requests.DataBody[attributes]) (Model, error) {
	id, err := strconv.ParseUint(body.Id, 10, 32)
	if err != nil {
		return Model{}, err
	}

	attr := body.Attributes
	return Model{
		id:   uint32(id),
		name: attr.Name,
	}, nil
}
