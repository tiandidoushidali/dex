package types

type GetNonceByAddressReq struct {
	WalletAddress string `json:"wallet_address"`
}

type GetNonceByAddressResp struct {
	Nonce string `json:"nonce"`
}

type WalletAddressLoginReq struct {
	WalletAddress string `json:"wallet_address" binding:"len=42,startwith=0x"`
	Signature     string `json:"signature" binding:"required"`
}

type JwtAuthResp struct {
	Token string `json:"token"`
}

type LoginReq struct {
	Username string `json:"username"`
}

type LoginResp struct {
	Token string `json:"token"`
}
