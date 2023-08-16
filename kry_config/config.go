package kry_config

// SetTokenForShopIdHandle 设置店铺的 token 的回调函数
type SetTokenForShopIdHandle func(shopId int64, token string)

// GetTokenForShopIdHandle 获取店铺的 token 的回调函数
type GetTokenForShopIdHandle func(shopId int64) string

// PrintApiLogHandle 打印 API 日志的回调函数
type PrintApiLogHandle func()

type Config struct {
	Domain                  string                  // 域名
	AppKey                  string                  // 应用标识
	SecretKey               string                  // 应用密钥
	Version                 string                  // 版本号, 默认为 2.0
	SetTokenForShopIdHandle SetTokenForShopIdHandle // 设置ShopId对应的Token
	GetTokenForShopIdHandle GetTokenForShopIdHandle // 获取ShopId对应的Token
}
