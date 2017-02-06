package main // import "vk8s.org/kapt/services/agentd"

import (
	"vk8s.io/kapt/services/agent/agentlib"
)

func main() {
	agentlib.Welcome("Hello from AgentDaemon")
}
