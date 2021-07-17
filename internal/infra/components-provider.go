package infra

import (
	"crawler/internal/components"
	"crawler/internal/components/frontier"
)

func ProvideFrontier() components.Frontier {
	return frontier.NewFrontier()
}
