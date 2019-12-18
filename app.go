package goRed_ui

import (
	_ "github.com/kakaisaname/account/core/accounts"
	_ "github.com/kakaisaname/goRed/core/envelopes"
	"github.com/kakaisaname/infra"
	"github.com/kakaisaname/infra/base"
	_ "goRed-ui/views"
)

func init() {
	infra.Register(&base.PropsStarter{})
	infra.Register(&base.DbxDatabaseStarter{})
	infra.Register(&base.ValidatorStarter{})
	infra.Register(&base.IrisServerStarter{})
	infra.Register(&infra.WebApiStarter{})
	infra.Register(&base.EurekaStarter{})
	infra.Register(&base.HookStarter{})
}
