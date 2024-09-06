package seeders

import (
	"fmt"
	"mood/database/factories"
	"mood/pkg/console"
	"mood/pkg/logger"
	"mood/pkg/seed"

	"gorm.io/gorm"
)

func init() {

	seed.Add("SeedRolesTable", func(db *gorm.DB) {

		roles := factories.MakeRoles(10)

		result := db.Table("roles").Create(&roles)

		if err := result.Error; err != nil {
			logger.LogIf(err)
			return
		}

		console.Success(fmt.Sprintf("Table [%v] %v rows seeded", result.Statement.Table, result.RowsAffected))
	})
}
