package echo

type ResponseBody struct {
	Type      string `json:"type"`
	MessageID int    `json:"msg_id"`
	InReplyTo int    `json:"in_reply_to"`
	Echo      string `json:"echo"`
}

type Response struct {
	Source      string `json:"src"`
	Destination string `json:"dest"`
	Body        ResponseBody
}
