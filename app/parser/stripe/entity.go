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
		Status             string `json:"status"`
		// --- Items
		Items struct {
			Object     string `json:"object"`
			HasMore    bool   `json:"has_more"`
			TotalCount int64  `json:"total_count"`
			URL        string `json:"url"`
			// --- Data
			Data []struct {
				Id           string         `json:"id"`
				Object       string         `json:"object"`
				Created      int64          `json:"created"`
				Quantity     int64          `json:"quantity"`
				Subscription string         `json:"subscription"`
				Metadata     map[string]any `json:"metadata"`
				// --- Plan
				Plan struct {
					Id            string         `json:"id"`
					Object        string         `json:"object"`
					Active        bool           `json:"active"`
					Amount        int64          `json:"amount"`
					AmountDecimal string         `json:"amount_decimal"`
					BillingScheme string         `json:"billing_scheme"`
					Created       int64          `json:"created"`
					Currency      string         `json:"currency"`
					Interval      string         `json:"interval"`
					IntervalCount int64          `json:"interval_count"`
					Product       string         `json:"product"`
					UsageType     string         `json:"usage_type"`
					Metadata      map[string]any `json:"metadata"`
				} `json:"plan"`
				// --- Price
				Price struct {
					Id                string         `json:"id"`
					Object            string         `json:"object"`
					Active            bool           `json:"active"`
					Amount            int64          `json:"amount"`
					AmountDecimal     string         `json:"amount_decimal"`
					BillingScheme     string         `json:"billing_scheme"`
					Created           int64          `json:"created"`
					Currency          string         `json:"currency"`
					Product           string         `json:"product"`
					TaxBehavior       string         `json:"tax_behavior"`
					Type              string         `json:"type"`
					UnitAmount        int64          `json:"unit_amount"`
					UnitAmountDecimal string         `json:"unit_amount_decimal"`
					Metadata          map[string]any `json:"metadata"`
					// --- Recurring
					Recurring struct {
						Interval      string `json:"interval"`
						IntervalCount int64  `json:"interval_count"`
						UsageType     string `json:"usage_type"`
					} `json:"recurring"`
				} `json:"price"`
			} `json:"data"`
		} `json:"items"`
	} `json:"object"`
	// ---- Previous Attributes
	PreviousAttrs struct {
		CurrentPeriodEnd   int64  `json:"current_period_end"`
		CurrentPeriodStart int64  `json:"current_period_start"`
		LatestInvoice      string `json:"latest_invoice"`
		Status             string `json:"status"`
	} `json:"previous_attributes"`
}
