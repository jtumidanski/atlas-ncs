package quest

import "github.com/sirupsen/logrus"

func Progress(l logrus.FieldLogger) func(characterId uint32, questId uint32) string {
	return func(characterId uint32, questId uint32) string {
		//TODO
		return ""
	}
}

func ProgressInt(l logrus.FieldLogger) func(characterId uint32, questId uint32, infoNumber int) int {
	return func(characterId uint32, questId uint32, infoNumber int) int {
		//TODO
		return 0
	}
}

func SetProgress(l logrus.FieldLogger) func(characterId uint32, questId uint32, infoNumber int, progress uint32) {
	return func(characterId uint32, questId uint32, infoNumber int, progress uint32) {
		//TODO
	}
}

func SetProgressString(l logrus.FieldLogger) func(characterId uint32, questId uint32, progress string) {
	return func(characterId uint32, questId uint32, progress string) {
		//TODO
	}
}

func IsActive(l logrus.FieldLogger) func(characterId uint32, questId uint32) bool {
	return func(characterId uint32, questId uint32) bool {
		//TODO
		return false
	}
}

func AnyActive(l logrus.FieldLogger) func(characterId uint32, questId ...uint32) bool {
	return func(characterId uint32, questId ...uint32) bool {
		for _, q := range questId {
			active := IsActive(l)(characterId, q)
			if active {
				return true
			}
		}
		return false
	}
}

func ForceComplete(l logrus.FieldLogger) func(characterId uint32, questId uint32) {
	return func(characterId uint32, questId uint32) {
		//TODO
	}
}

func IsNotStarted(l logrus.FieldLogger) func(characterId uint32, questId uint32) bool {
	return func(characterId uint32, questId uint32) bool {
		//TODO
		return true
	}
}

func Start(l logrus.FieldLogger) func(characterId uint32, questId uint32) {
	return func(characterId uint32, questId uint32) {
		//TODO
	}
}

func IsCompleted(l logrus.FieldLogger) func(characterId uint32, questId uint32) bool {
	return func(characterId uint32, questId uint32) bool {
		//TODO
		return false
	}
}

func IsStarted(l logrus.FieldLogger) func(characterId uint32, questId uint32) bool {
	return func(characterId uint32, questId uint32) bool {
		//TODO
		return false
	}
}

func CompleteViaNPC(l logrus.FieldLogger) func(characterId uint32, questId uint32, npcId uint32) {
	return func(characterId uint32, questId uint32, npcId uint32) {
		//TODO
	}
}

func Complete(l logrus.FieldLogger) func(characterId uint32, questId uint32) {
	return func(characterId uint32, questId uint32) {
		//TODO
	}
}

