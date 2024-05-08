package stripe

type WebhookPayload struct {
	Id         string      `json:"id"`
	Object     string      `json:"object"`
	ApiVersion string      `json:"api_version"`
	Created    int64       `json:"created"`
	Type       string      `json:"type"`
	Data       WebhookData `json:"data"`
	Request    struct {
		Id             string `json:"id"`
		IdempotencyKey string `json:"idempotency_key"`
	} `json:"request"`
}

type WebhookData struct {
	Object struct {
		Id                 string `json:"id"`
		Object             string `json:"object"`
		CollectionMethod   string `json:"collection_method"`
		Created            int64  `json:"created"`
		Currency           string `json:"currency"`
		CurrentPeriodEnd   int64  `json:"current_period_end"`
		CurrentPeriodStart int64  `json:"current_period_start"`
		Customer           string `json:"customer"`
		Description        string `json:"description"`
		Quantity           int64  `json:"quantity"`
		StartDate          int64  `json:"start_date"`
		Status             bool   `json:"status"`
		Items              []struct {
			Object     string `json:"object"`
			HasMore    bool   `json:"has_more"`
			TotalCount int64  `json:"total_count"`
			URL        string `json:"url"`
			Data       []struct {
				Row struct {
					Id       string `json:"id"`
					Object   string `json:"object"`
					Created  int64  `json:"created"`
					Metadata any    `json:"metadata"`
					Plan     struct {
						Id            string `json:"id"`
						Object        string `json:"object"`
						Active        bool   `json:"active"`
						Amount        int64  `json:"amount"`
						AmountDecimal int64  `json:"amount_decimal"`
						BillingSchema string `json:"billing_schema"`
						Created       int64  `json:"created"`
						Currency      string `json:"currency"`
						Interval      string `json:"interval"`
						IntervalCount int64  `json:"interval_count"`
						Metadata      any    `json:"metadata"`
						Product       string `json:"product"`
						UsageType     string `json:"usage_type"`
					} `json:"plan"` // --- end Object.Items.Data.Row.Plan
					Price struct {
						Id            string `json:"id"`
						Object        string `json:"object"`
						Active        bool   `json:"active"`
						BillingSchema string `json:"billing_schema"`
						Created       int64  `json:"created"`
						Currency      string `json:"currency"`
						Metadata      any    `json:"metadata"`
						Product       string `json:"product"`
						Recurring     struct {
							Interval      string `json:"interval"`
							IntervalCount int64  `json:"interval_count"`
							UsageType     string `json:"usage_type"`
						} `json:"recurring"`
						TaxBehaviour      string `json:"tax_behaviour"`
						Type              string `json:"type"`
						UnitAmount        int64  `json:"unit_amount"`
						UnitAmountDecimal int64  `json:"unit_amount_decimal"`
					} `json:"price"` // --- end Object.Items.Data.Row.Price
					Quantity     int64  `json:"quantity"`
					Subscription string `json:"subscription"`
				} `json:"row"` // ---> end: Object.Items.Data.Row
			} `json:"data"` // ---> end: Object.Items.Data
		} `json:"items"` // ---> end: Object.Items
	} `json:"object"` // ---> end: Object
	PreviousAttrs struct {
		CurrentPeriodEnd   int64  `json:"current_period_end"`
		CurrentPeriodStart int64  `json:"current_period_start"`
		LatestInvoice      string `json:"latest_invoice"`
		Status             string `json:"status"`
	} `json:"previous_attrs"` // ---> end: PreviousAttrs
}
