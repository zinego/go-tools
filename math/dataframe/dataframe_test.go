package dataframe

import (
	"testing"

	"github.com/zinego/go-tools/utils/log"
)

func TestLoadFromMapStringFloat64(t *testing.T) {
	t.Run("correct", func(t *testing.T) {
		var m = map[string][]float64{
			"hello": {0.1, 0.2, 0.3, 0.4},
			"world": {0.2, 0.3, 0.3, 0.4},
		}
		r := LoadFromMapStringFloat64(m)
		log.Info(r.String())
	})
	t.Run("test", func(t *testing.T) {
		test(1, 2, 4)
		test(1, 2, 342, 423, 423, 423, 4, 234, 23)
	})
}

func test(i ...int) {
	for _, v := range i {
		log.Info(v)
	}
}
