package mergers

//go:generate go run github.com/v2fly/v2ray-core/v5/common/errors/errorgen

import (
    "log"
	"strings"

	core "github.com/v2fly/v2ray-core/v5"
	"github.com/v2fly/v2ray-core/v5/common"
	"github.com/v2fly/v2ray-core/v5/infra/conf/json"
)

func init() {
    log.Println("infra/conf/mergers/mergers.go:15 init")
	common.Must(RegisterMerger(makeMerger(
		core.FormatJSON,
		[]string{".json", ".jsonc"},
		nil,
	)))
	common.Must(RegisterMerger(makeMerger(
		core.FormatTOML,
		[]string{".toml"},
		json.FromTOML,
	)))
	common.Must(RegisterMerger(makeMerger(
		core.FormatYAML,
		[]string{".yml", ".yaml"},
		json.FromYAML,
	)))
	common.Must(RegisterMerger(
		&Merger{
			Name:       core.FormatAuto,
			Extensions: nil,
			Merge:      Merge,
		}),
	)
}

// Merger is a configurable format merger for V2Ray config files.
type Merger struct {
	Name       string
	Extensions []string
	Merge      MergeFunc
}

// MergeFunc is a utility to merge V2Ray config from external source into a map and returns it.
type MergeFunc func(input interface{}, m map[string]interface{}) error

var (
	mergersByName = make(map[string]*Merger)
	mergersByExt  = make(map[string]*Merger)
)

// RegisterMerger add a new Merger.
func RegisterMerger(format *Merger) error {
	if _, found := mergersByName[format.Name]; found {
		return newError(format.Name, " already registered.")
	}
    log.Printf("infra/conf/mergers/mergers.go:59 RegisterMerger reg %v\n", format.Name)
	mergersByName[format.Name] = format

	for _, ext := range format.Extensions {
		lext := strings.ToLower(ext)
		if f, found := mergersByExt[lext]; found {
			return newError(ext, " already registered to ", f.Name)
		}
        log.Printf("infra/conf/mergers/mergers.go:67 RegisterMerger reg ext %v\n", lext)
		mergersByExt[lext] = format
	}

	return nil
}
