// Package factories 存放工厂方法
package factories

import (
	"mood/app/models/user"
	"mood/pkg/helpers"

	"github.com/bxcodec/faker/v3"
	"github.com/spf13/cast"
)

func MakeUsers(times int) []user.User {

	var objs []user.User

	// 设置唯一值
	faker.SetGenerateUniqueValues(true)

	model1 := user.User{
		Name:     "admin",
		NickName: "admin",
		Email:    "admin@163.com",
		Phone:    "00015622888",
		Password: "$2a$14$oPzVkIdwJ8KqY0erYAYQxOuAAlbI/sFIsH0C0R4MPc.3JbWWSuaUe",
		RoleID:   "1",
	}

	objs = append(objs, model1)

	model1 = user.User{
		Name:     "yanghaibing",
		NickName: "yanghaibing",
		Email:    "yanghaibing@163.com",
		Phone:    "00015622999",
		Password: "$2a$14$oPzVkIdwJ8KqY0erYAYQxOuAAlbI/sFIsH0C0R4MPc.3JbWWSuaUe",
		RoleID:   "1",
	}

	objs = append(objs, model1)

	for i := 0; i < times; i++ {
		model := user.User{
			Name:     faker.Username(),
			NickName: cast.ToString(i+2) + "号", // faker.Name()
			Email:    faker.Email(),
			Phone:    helpers.RandomNumber(11),
			Password: "$2a$14$oPzVkIdwJ8KqY0erYAYQxOuAAlbI/sFIsH0C0R4MPc.3JbWWSuaUe",
			RoleID:   "3",
		}
		objs = append(objs, model)
	}

	return objs
}
