package dataframe

import (
	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
)

func LoadFromMapStringFloat64(m map[string][]float64) dataframe.DataFrame {
	var list []series.Series
	for k, v := range m {
		list = append(list, series.New(v, series.Float, k))
	}
	return dataframe.New(list...)
}
