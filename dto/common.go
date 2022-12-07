package dto

type RequestPayload struct {
	Page     int `json:"page" form:"page" binding:"required"`
	PageSize int `json:"page_size" form:"page_size" binding:"required"`
}

type Response struct {
	ErrCode        uint64 `json:"err_code"`
	ErrMsg         string `json:"err_msg"`
	Data           any    `json:"data,omitempty"`
	HttpStatusCode int    `json:"-"`
}

type ReqByID struct {
	ID uint64 `json:"id" form:"id" binding:"required"`
}

type ReqByIDs struct {
	IDS []uint64 `json:"ids" form:"ids" binding:"required,gte=1"`
}

type OTPReq struct {
	OTP string `json:"otp" form:"otp" binding:"required"`
}

type PassphraseReq struct {
	Passphrase string `json:"passphrase" binding:"required,checkphrase=12"`
	Network    string `json:"network" binding:"required,oneof='ERC20' 'TRC20'"`
}

type TransferReq struct {
	ID          uint64  `json:"id" binding:"required"`
	Amount      float64 `json:"amount"`
	FromAddress string  `json:"from_address"`
	ToAddress   string  `json:"to_address"`
	PrivateKey  string  `json:"private_key"`
	IP          string  `json:"ip"`
	Area        string  `json:"area"`
}

type PageReq struct {
	Page     int `json:"page" form:"page" binding:"required"`
	PageSize int `json:"page_size" form:"page_size" binding:"required"`
}

type Transaction struct {
	TxID string `json:"txId"`
}
