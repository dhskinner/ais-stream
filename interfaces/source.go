package interfaces

import "ais-stream/models"

type Source interface {
	Name() string
	Stream() chan *models.Sentence
	PrintStats()
}
