package model

import (
	//"database/sql"
	"github.com/CourseComment/conf"
	_ "github.com/go-sql-driver/mysql"
	//"os"
	"time"
)

var (
	timepara   int       = conf.Timepara
	timeLayout string    = conf.TimeLayout
	timeStart  time.Time = conf.TimeStart
)

type sortelement struct {
	Id           idtype
	T            time.Time
	Super_number int
	Score        float64
}

func caculateScore(data []sortelement, super_number int) {
	for i := range data {
		data[i].Score = 0
		data[i].Score += float64(data[i].Super_number) / float64(super_number) * 100 // super
		data[i].Score += data[i].T.Sub(timeStart).Seconds() / float64(timepara)
	}
}

func sorting(data []sortelement) {
	//I need  a quick-sork in go
	//but now i use a simple one
	for i := range data {
		for j := i + 1; j < len(data); j++ {
			if data[i].Score < data[j].Score {
				data[i], data[j] = data[j], data[i]
			}
		}
	}

}

func TradeOffSorting(data []sortelement, super_number int) {
	caculateScore(data, super_number)
	sorting(data)
}
