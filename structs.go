package uexchange

// Client - ..
type Client struct {
	APICredentials Credentials
}

// Credentials - ..
type Credentials struct {
	AccountPublicKey string `json:"PublicKey"` // Utopia Account public key
	Password         string `json:"password"`  // Exchange User password
	TwoFACode        string `json:"2fa_pin"`   // 2-factor-authorization code
}

// APIPlainResponse - ..
type APIPlainResponse struct {
	Success bool   `json:"success"`
	Result  string `json:"result"`
}

// APIAuthResponse - ..
type APIAuthResponse struct {
	Success bool                   `json:"success"`
	Result  APIAuthResultContainer `json:"result"`
}

// APIAuthResultContainer - ..
type APIAuthResultContainer struct {
	UserSession UserSessionData `json:"user_session"`
	AuthToken   string          `json:"auth_token"`
}

// APIBalanceResponse - ..
type APIBalanceResponse struct {
	Success bool            `json:"success"`
	Result  BalanceResponse `json:"result"`
}

// BalanceResponse - ..
type BalanceResponse struct {
	AllBalance []BalanceData `json:"allbalance"`
	UserID     string        `json:"user_id"`
}

// BalanceData - ..
type BalanceData struct {
	ID       string       `json:"id"`
	Currency CurrencyData `json:"currency"`
	Reserve  float64      `json:"reserve"`
	Balance  float64      `json:"balance"`
}

// APITradeResponse - ..
type APITradeResponse struct {
	Success  bool  `json:"success"`
	OrderID  int64 `json:"order_id"`
	DaemonID int64 `json:"daemon_id"`
}

// CurrencyData - ..
type CurrencyData struct {
	ID            string  `json:"id"`
	Name          string  `json:"name"`             // example: crp
	FullName      string  `json:"fullname"`         // example: Utopia Crypton
	AppName       string  `json:"appname"`          // example: crypton
	Icon          string  `json:"icon"`             // example: crp
	Round         int     `json:"round"`            // round precision, example: 8
	DepositFee    float64 `json:"deposit_fee"`      // example: 0
	DepositFeePRO float64 `json:"withdraw_fee_pro"` // example: 0.1
	MinWithdraw   float64 `json:"withdraw_min"`     // example: 5
	AddressSize   int     `json:"address_size"`     // example: 64
	MinFee        float64 `json:"min_fee"`          // example: 0.00000001
	Enable        bool    `json:"enable"`           // example: true
	Visible       bool    `json:"show"`             // example: true
}

// UserSessionData - ..
type UserSessionData struct {
	User       UserData       `json:"user"`
	APISession APISessionData `json:"session"`
}

// UserData - ..
type UserData struct {
	ID     string `json:"id"`     // exchange user ID -- UUID format
	Name   string `json:"name"`   // username
	Status string `json:"status"` // example: active
	Lang   string `json:"lang"`   // user language
}

// APISessionData - ..
type APISessionData struct {
	ID string `json:"id"` // session ID -- UUID format
}
