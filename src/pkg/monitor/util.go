package monitor

import (
	"github.com/labstack/gommon/log"
	"github.com/tidwall/gjson"
)

func ExtractorPath(data string, paths []string) (tar [][]string) {
	ret := gjson.Parse(data)

	for _, path := range paths {
		var item []string
		val := ret.Get(path)

		if val.IsArray() {
			val.ForEach(func(key, value gjson.Result) bool {
				item = append(item, value.String())
				return true
			})
		} else {
			item = append(item, val.String())
		}
		if len(item) == 0 {
			log.Errorf("len(item) == 0  %s  || %s", path, data)
		}
		tar = append(tar, item)

	}
	return
}
