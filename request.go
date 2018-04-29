package ebay

import (
	"github.com/bitbandit01/go-ebay/commands"
	"github.com/bitbandit01/go-ebay/config"
)

const xmlns = "urn:ebay:apis:eBLBaseComponents"

type Request struct {
	Config  *config.Config
	Command commands.Command
}
