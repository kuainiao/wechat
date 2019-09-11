package pay

const (
	Fail       = "FAIL"
	Success    = "SUCCESS"
	HMACSHA256 = "HMAC-SHA256"
	MD5        = "MD5"
	Sign       = "sign"

	microPayURL                = "https://api.mch.weixin.qq.com/pay/micropay"
	unifiedOrderURL            = "https://api.mch.weixin.qq.com/pay/unifiedorder"
	orderQueryURL              = "https://api.mch.weixin.qq.com/pay/orderquery"
	reverseURL                 = "https://api.mch.weixin.qq.com/secapi/pay/reverse"
	closeOrderURL              = "https://api.mch.weixin.qq.com/pay/closeorder"
	refundURL                  = "https://api.mch.weixin.qq.com/secapi/pay/refund"
	refundQueryURL             = "https://api.mch.weixin.qq.com/pay/refundquery"
	downloadBillURL            = "https://api.mch.weixin.qq.com/pay/downloadbill"
	reportURL                  = "https://api.mch.weixin.qq.com/payitil/report"
	shortURL                   = "https://api.mch.weixin.qq.com/tools/shorturl"
	authCodeToOpenidURL        = "https://api.mch.weixin.qq.com/tools/authcodetoopenid"
	getSignkeyURL              = "https://api.mch.weixin.qq.com/sandboxnew/pay/getsignkey"
	microPaySandboxURL         = "https://api.mch.weixin.qq.com/sandboxnew/pay/micropay"
	unifiedOrderSandboxURL     = "https://api.mch.weixin.qq.com/sandboxnew/pay/unifiedorder"
	orderQuerySandboxURL       = "https://api.mch.weixin.qq.com/sandboxnew/pay/orderquery"
	reverseSandboxURL          = "https://api.mch.weixin.qq.com/sandboxnew/secapi/pay/reverse"
	closeOrderSandboxURL       = "https://api.mch.weixin.qq.com/sandboxnew/pay/closeorder"
	refundSandboxURL           = "https://api.mch.weixin.qq.com/sandboxnew/secapi/pay/refund"
	refundQuerySandboxURL      = "https://api.mch.weixin.qq.com/sandboxnew/pay/refundquery"
	downloadBillSandboxURL     = "https://api.mch.weixin.qq.com/sandboxnew/pay/downloadbill"
	reportSandboxURL           = "https://api.mch.weixin.qq.com/sandboxnew/payitil/report"
	shortSandboxURL            = "https://api.mch.weixin.qq.com/sandboxnew/tools/shorturl"
	authCodeToOpenidSandboxURL = "https://api.mch.weixin.qq.com/sandboxnew/tools/authcodetoopenid"

	transferURL     = "https://api.mch.weixin.qq.com/mmpaymkttransfers/promotion/transfers"
	transferBankURL = "https://api.mch.weixin.qq.com/mmpaysptrans/pay_bank"
)