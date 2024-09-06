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
	}

	type User struct {
		models.BaseModel

		Name     string `gorm:"type:varchar(255);not null;index"`
		NickName string `gorm:"type:varchar(255);not null;index"`
		Email    string `gorm:"type:varchar(255);index;default:null"`
		Phone    string `gorm:"type:varchar(20);index;default:null"`
		Password string `gorm:"type:varchar(255)"`

		RoleID string `gorm:"type:bigint;index"`

		Role Role

		models.CommonTimestampsField
	}

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.AutoMigrate(&User{})
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.DropTable(&User{})
	}

	migrate.Add("2024_03_25_195341_add_users_table", up, down)
	// FIXME("操作错误，自动生成的 migrate.Add 名称，无需修改！！！")
}
