package performancedb

import (
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Service struct {
	DB *gorm.DB
}

// have the interface  here

func NewService(db *gorm.DB) *Service {
	return &Service{
		DB: db,
	}
}

func (s *Service) Stats() [5]time.Duration {

	var dbPerformance [5]time.Duration

	for i := 0; i < 5; i++ {
		// query database
		firstQueryStart := time.Now()
		rows, err := s.DB.Debug().Raw("select 1;").Rows()

		firstQueryEnd := time.Now()
		if err != nil {
			panic(err)
		}

		// put the connection back to the pool so
		// that it can be reused by next iteration
		rows.Close()

		dbPerformance[i] = firstQueryEnd.Sub(firstQueryStart)
		logrus.Info(fmt.Sprintf("query #%d took %s", i, firstQueryEnd.Sub(firstQueryStart).String()))
	}

	return dbPerformance
}
