package req

type SubscribeRequest struct {
	UserId    string `json:"userId"`
	AccountId string `json:"accountId"`
}
