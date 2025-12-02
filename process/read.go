package process

import (
	"fmt"

	"github.com/jasontconnell/sitecore/api"
	"github.com/jasontconnell/sitecore/data"
)

func LoadItems(connstr, protobufLocation, path string) (data.ItemMap, error) {
	var pitems []data.ItemNode
	var perr error
	if protobufLocation != "" {
		pitems, perr = api.ReadProtobuf(protobufLocation)
		if perr != nil {
			return nil, fmt.Errorf("reading protobuf %w", perr)
		}
	}

	items, err := api.LoadItems(connstr)
	if err != nil {
		return nil, fmt.Errorf("reading items %w", err)
	}

	items = append(items, pitems...)

	_, m := api.LoadItemMap(items)

	m = api.FilterItemMap(m, []string{path})

	return m, nil
}
