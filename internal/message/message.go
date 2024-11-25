package message

type Type string

const (
	Echo       Type = "echo"
	EchoOk     Type = "echo_ok"
	Generate   Type = "generate"
	GenerateOk Type = "generate_ok"
)

type Message struct {
	Type      Type `json:"type"`
	MessageID uint `json:"msg_id"`
}
