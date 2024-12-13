package model

type Menu struct {
	ID        int     `gorm:"primaryKey" json:"id" structs:"id"`
	Name      string  `gorm:"type:varchar(32);unique;not null" json:"name"`
	Path      string  `gorm:"type:varchar(32)" json:"path"`
	Level     int     `gorm:"default:1" json:"level"`
	Pid       int     // Parent menu ID
	RoleLevel int     `gorm:"default:1"` // Role level 1 代表只有管理员可以使用 2 代.231表管理员和用户都可以使用
	Children  []*Menu `gorm:"foreignKey:pid" json:"children"`
}

func (Menu) TableName() string {
	return "menus"
}

// BuildMenuTree 将 MenuModel 列表转换为树状结构
func BuildMenuTree(menuList []Menu) []*Menu {
	// 创建一个映射，用于快速查找每个节点的 ID 对应的索引
	menuMap := make(map[int]*Menu)
	for i := range menuList {
		menuMap[menuList[i].ID] = &menuList[i]
	}
	// 遍历每个节点，将其添加到父节点的 Children 字段中
	for i := range menuList {
		menu := &menuList[i]
		if menu.Pid != 0 {
			// 否则，将当前节点添加到父节点的 Children 字段中
			parent, ok := menuMap[menu.Pid]
			if ok {
				parent.Children = append(parent.Children, menu)
			}
		}
	}
	return menuList[0].Children
}
