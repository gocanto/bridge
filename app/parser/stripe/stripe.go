package stripe

type WebhookPayload struct {
	Id         string         `json:"id"`
	Object     string         `json:"object"`
	ApiVersion string         `json:"api_version"`
	Type       string         `json:"type"`
	CreatedAt  int64          `json:"created_at"`
	Request    WebhookRequest `json:"request"`
}

type WebhookRequest struct {
	Id             string `json:"id"`
	IdempotencyKey string `json:"idempotency_key"`
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
		Items              []struct {
			Object string `json:"object"`
		}
	}
}
