package uexchange

// Client - ..
type Client struct {
	APICredentials Credentials
}

// Credentials - ..
type Credentials struct {
	AccountPublicKey string `json:"PublicKey"` // Utopia Account public key
	Password         string `json:"password"`  // Exchange User password
	TwoFACode        string `json:"2fa_pin"`   // 2-factor-authorization code. optional
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

// APIPairsResponse - ..
type APIPairsResponse struct {
	Success bool                 `json:"success"`
	Result  []PairsDataContainer `json:"pairs"`
}

// PairsDataContainer - ..
type PairsDataContainer struct {
	Pair       PairData            `json:"pairs"`
	MarketData MarketDataContainer `json:"data_market"`
}

// PairData - ..
type PairData struct {
	ID              int     `json:"pair_id"`           // example: 25
	PairCode        string  `json:"pair"`              // example: crp_usdt
	PairTitle       string  `json:"pair_show"`         // example: CRP / USDT
	CoinsGroup      string  `json:"group"`             // example: crp
	Visible         bool    `json:"visible"`           // example: true
	Enabled         bool    `json:"enable"`            // example: true
	RoundDealAmount int     `json:"round_deal_amount"` // example: 3
	RoundDealPrice  int     `json:"round_deal_price"`  // example: 4
	MinAmount       float64 `json:"min_amount"`        // 1
	MinPrice        float64 `json:"min_price"`         // 0.001
	MaxPrice        float64 `json:"max_price"`         // 100
}

// MarketDataContainer - ..
type MarketDataContainer struct {
	Open      float64 `json:"open"`       // example: 0.1744
	Close     float64 `json:"close"`      // example: 0.1752
	High      float64 `json:"high"`       // example: 0.1766
	Low       float64 `json:"low"`        // example: 0.1553
	Volume    float64 `json:"volume"`     // example: 67044.815
	VolumeUSD float64 `json:"volume_usd"` // example: 1174.6252
	Value     float64 `json:"value"`      // example: 11346.6402207
	Rate      float64 `json:"rate"`       // example: 0.46
	DateNow   int64   `json:"date_now"`   // example: 1634566376377
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

// APIBookValueResponse - ..
type APIBookValueResponse struct {
	Success bool                   `json:"success"`
	Result  BookValueDataContainer `json:"result"`
}

// BookValueDataContainer - ..
type BookValueDataContainer struct {
	Sell []BookValueData `json:"book_sell"`
	Buy  []BookValueData `json:"book_buy"`
}

// BookValueData - ..
type BookValueData struct {
	Price  float64 `json:"price"`  // example: 0.1752
	Amount float64 `json:"amount"` // example: 1555
	Value  float64 `json:"value"`  // example: 272.436
}

// APICurrenciesListResponse - ..
type APICurrenciesListResponse struct {
	Success bool               `json:"success"`
	Result  CurrenciesListData `json:"result"`
}

// CurrenciesListData - ..
type CurrenciesListData map[string]CurrencyData

// APIOrdersResponse - response from get orders list
type APIOrdersResponse struct {
	Success bool                `json:"success"`
	Result  OrdersDataContainer `json:"result"`
}

// OrdersDataContainer - orders response data
type OrdersDataContainer struct {
	AllOrders     []OrderData    `json:"allorders"`
	Pagination    PaginationData `json:"paginator"`
	Filters       []string       `json:"filters"`
	MyOrdersCount int            `json:"my_order_count"`
}

// OrderData - ..
type OrderData struct {
	OrderID               int64   `json:"order_id"`       // example: 3747944
	Amount                float64 `json:"amount"`         // example: 1
	Price                 float64 `json:"price"`          // example: 3
	Value                 float64 `json:"value"`          // example: 3
	OriginalValue         float64 `json:"orig_value"`     // example: 3
	OriginalAmount        float64 `json:"orig_amount"`    // example: 1
	OrderRegistrationDate int64   `json:"date_reg"`       // example: 1527883807109
	TaskType              string  `json:"task"`           // example: sell
	Status                string  `json:"status"`         // example: open
	BaseTicker            string  `json:"cur"`            // example: crp
	QuoteTicker           string  `json:"ecur"`           // example: usdt
	ExecutedPrice         float64 `json:"price_executed"` // example: 0
	ExecutedValue         float64 `json:"value_executed"` // example: 0
}

// PaginationData - ..
type PaginationData struct {
	Count int `json:"count"`
	Page  int `json:"page"`
	Limit int `json:"limit"`
}

// APIOrdersHistoryResponse - ..
type APIOrdersHistoryResponse struct {
	Success bool                       `json:"success"`
	Result  OrdersHistoryDataContainer `json:"result"`
}

// OrdersHistoryDataContainer - ..
type OrdersHistoryDataContainer struct {
	Order   OrderData             `json:"order"`
	History []OrderData           `json:"history"`
	Pair    OrdersHistoryPairData `json:"pair"`
}

// OrdersHistoryPairData - ..
type OrdersHistoryPairData struct {
	BaseTicker  string `json:"cur"`
	QuoteTicker string `json:"ecur"`
}

// APITradeHistoryResponse - ..
type APITradeHistoryResponse struct {
	Success bool                      `json:"success"`
	Result  TradeHistoryDataContainer `json:"result"`
}

// TradeHistoryDataContainer - ..
type TradeHistoryDataContainer struct {
	Items []TradeHistoryData    `json:"items"`
	Pair  OrdersHistoryPairData `json:"pair"`
}

// TradeHistoryData - ..
type TradeHistoryData struct {
	RecordID   string  `json:"record_id"`   // UUID format, example: f914ea5d-32ae-4eac-a5e0-e94d55f7b4a7
	RecordType string  `json:"record_type"` // example: buy
	Price      float64 `json:"price"`       // example: 0.1752
	Amount     float64 `json:"amount"`      // example: 1
	Value      float64 `json:"value"`       // 0.1752
	CreatedAt  int64   `json:"created_at"`  // date, example: 1634566302532
}

// APIOperationsHistoryResponse - ..
type APIOperationsHistoryResponse struct {
	Success bool                           `json:"success"`
	Result  OperationsHistoryDataContainer `json:"result"`
}

// OperationsHistoryDataContainer - ..
type OperationsHistoryDataContainer struct {
	Items []TradeHistoryData `json:"items"`
}
