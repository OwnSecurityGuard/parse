package monitor

import (
	"github.com/labstack/gommon/log"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type ExtractUtil interface {
	GetValue(data string, path string) []string
	GetValues(data string, path []string) [][]string
	GetValueByBytes(data []byte, path string) []string
	GetValuesByBytes(data []byte, path []string) [][]string

	SetValue(srcData string, path string, value string) (string, error)

	Filter(src string, data []string, filters []*Filter) []string
}

type GJsonExtractUtil struct {
}

func (G *GJsonExtractUtil) GetValueByBytes(data []byte, path string) []string {
	return G.GetValue(string(data), path)
}

func (G *GJsonExtractUtil) GetValuesByBytes(data []byte, path []string) [][]string {
	return G.GetValues(string(data), path)
}

func NewGJsonExtractUtil() *GJsonExtractUtil {
	return &GJsonExtractUtil{}
}

func (G *GJsonExtractUtil) GetValue(data string, path string) (tar []string) {
	ret := gjson.Parse(data)
	val := ret.Get(path)
	if val.IsArray() {
		val.ForEach(func(key, value gjson.Result) bool {
			tar = append(tar, value.String())
			return true
		})
	} else {
		tar = append(tar, val.String())
	}

	return tar
}

func (G *GJsonExtractUtil) GetValues(data string, paths []string) (tar [][]string) {
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

func (G *GJsonExtractUtil) Filter(src string, data []string, filters []*Filter) []string {
	var (
		srcRet     = gjson.Parse(src)
		filterData []string
	)
	for i := range data {
		ret := gjson.Parse(data[i])
		for _, filter := range filters {
			var actualVal string
			if len(filter.ExpectKey) == 0 {
				actualVal = filter.ExpectVal // ret.Get(filter.Path).String()
			} else {
				actualVal = srcRet.Get(filter.ExpectKey).String()
			}

			respVal := ret.Get(filter.Path).String()
			if len(respVal) > 0 && actualVal == respVal { // todo 暂时先用== ，  后续看filters op 字段来比较
				filterData = append(filterData, data[i])
			}
		}
	}
	return filterData
}

func (G *GJsonExtractUtil) SetValue(srcData string, path string, value string) (string, error) {
	return sjson.Set(srcData, path, value)
}
