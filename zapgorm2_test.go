package zapgorm2_test

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
	"moul.io/zapgorm2"
)

func Example() {
	logger := zapgorm2.New(zap.L())
	logger.SetAsDefault() // optional: configure gorm to use this zapgorm.Logger for callbacks
	db, _ := gorm.Open(nil, &gorm.Config{Logger: logger})

	// do stuff normally
	var _ = db // avoid "unused variable" warn
}
