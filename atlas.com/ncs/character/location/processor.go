package location

import (
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

func SaveLocation(l logrus.FieldLogger, span opentracing.Span) func(characterId uint32, location string) {
	return func(characterId uint32, location string) {
		_, err := saveLocation(l, span)(characterId, location)
		if err != nil {
			l.WithError(err).Errorf("Unable to save location %s for character %d.", location, characterId)
			return
		}
	}
}
