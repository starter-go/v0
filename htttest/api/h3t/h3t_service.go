package h3t

type Service interface {
	GetUserAgent() UserAgent

	GetAgentContext() *AgentContext

	GetCodecManager() CodecManager

	NewTransaction(method, url string) *Transaction
}
