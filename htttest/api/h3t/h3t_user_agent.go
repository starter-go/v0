package h3t

type UserAgent interface {
	GetAgentContext() *AgentContext

	Execute(c *Transaction) error
}
