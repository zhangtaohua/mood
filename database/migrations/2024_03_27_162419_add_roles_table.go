package migrations

import (
	"database/sql"
	"mood/app/models"
	"mood/pkg/migrate"

	"gorm.io/gorm"
)

func init() {

	type Role struct {
		models.BaseModel

		Name        string `gorm:"type:varchar(255);not null;index"`
		Description string `gorm:"type:varchar(511);default:null"`

		models.CommonTimestampsField
	}

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.AutoMigrate(&Role{})
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.DropTable(&Role{})
	}

	migrate.Add("2024_03_27_162419_add_roles_table", up, down)
}
