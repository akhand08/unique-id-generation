package protocols

type RequestBody struct {
	Type      string   `json:"type"`
	MessageID int      `json:"msg_id"`
	Echo      string   `json:"echo"`
	NodeID    string   `json:"node_id"`
	NodeIDs   []string `json:"node_ids"`
}

type InitResponseBody struct {
	Type      string `json:"type"`
	InReplyTo int    `json:"in_reply_to"`
}
type InitResponse struct {
	Source      string `json:"src"`
	Destination string `json:"dest"`
	Body        InitResponseBody
}

type Request struct {
	Source      string `json:"src"`
	Destination string `json:"dest"`
	Body        RequestBody
}
