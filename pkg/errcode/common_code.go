package errcode

// a collection for managing custom Error struct
var (
	Success                   = NewError(0, "成功")
	ServerError               = NewError(10000000, "服務內部錯誤")
	InvalidParams             = NewError(10000001, "導入參數錯誤")
	NotFound                  = NewError(10000002, "服務內部錯誤")
	UnauthorizedAuthNotExist  = NewError(10000003, "驗證失敗，找不到對應的 AppKey 和 AppSecret")
	UnauthorizedTokenError    = NewError(10000004, "驗證失敗，Token 錯誤")
	UnauthorizedTokenTimeout  = NewError(10000005, "驗證失敗，Token 逾時")
	UnauthorizedTokenGenetate = NewError(10000006, "驗證失敗，Token 產生失敗")
	TooManyRequests           = NewError(10000007, "驗證失敗，請求過多")
)
