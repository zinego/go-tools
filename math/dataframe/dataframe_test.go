package dataframe

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestLoadFromMapStringFloat64(t *testing.T) {
	t.Run("correct", func(t *testing.T) {
		var m = map[string][]float64{
			"hello": {0.1, 0.2, 0.3, 0.4},
			"world": {0.2, 0.3, 0.3, 0.4},
		}
		r := LoadFromMapStringFloat64(m)
		f, _ := os.Create(time.Now().String() + ".csv")
		defer f.Close()
		assert.Nil(t, r.WriteCSV(f))
	})
}
