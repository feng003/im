package constant

// 用户类型
const (
	UserTypeSystem      = 1 // 系统  配置的所有的店铺
	UserTypeDepartment  = 2 // 部门  看到部门员工或者自己的店铺
	UserTypeStore       = 3 // 店铺用户 店铺1对1
	UserTypeManageStore = 4 // 店铺管理员 管理多个店铺
	UserTypeStoreUser   = 5 // 店铺邀请码 单个店铺的邀请码数据
	UserTypeStoreStaff  = 6 // 店铺员工 单个店铺的多个员工数据
	UserTypeStoreShop   = 7 // 门店员工

	// 8区块链独立 7区块链聚合  6安心付聚合  5安心付独立

	MsAdmin        = 1 // admin后台
	MsStore        = 2 // store后台
	MsZhiMaJuHe    = 3 // 芝麻聚合
	MsZhiMaDuLi    = 4 // 芝麻独立
	MsCommerceDuLi = 5 // 安心付独立
	MsCommerceJuHe = 6 // 安心付聚合
	MsAntChainJuHe = 7 // 区块链聚合
	MsAntChainDuLi = 8 // 区块链独立
)
