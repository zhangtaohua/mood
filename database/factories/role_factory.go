package factories

import (
	"mood/app/models/role"
)

func MakeRoles(count int) []role.Role {

	var objs []role.Role

	// 设置唯一性，如 Role 模型的某个字段需要唯一，即可取消注释
	// faker.SetGenerateUniqueValues(true)

	roleModel := role.Role{
		Name:        "admin",
		Description: "管理员",
	}
	objs = append(objs, roleModel)

	roleModel = role.Role{
		Name:        "accountant",
		Description: "会计",
	}
	objs = append(objs, roleModel)

	roleModel = role.Role{
		Name:        "operation",
		Description: "运营",
	}

	objs = append(objs, roleModel)

	return objs
}
