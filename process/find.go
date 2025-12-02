package process

import (
	"strings"

	"github.com/jasontconnell/sitecore/data"
)

func FindOrpahs(items data.ItemMap, values map[string]bool) []data.ItemNode {
	var orphans []data.ItemNode
	for _, item := range items {
		id := strings.ReplaceAll(item.GetId().String(), "-", "")

		if _, ok := values[id]; !ok {
			orphans = append(orphans, item)
		}
	}
	return orphans
}
