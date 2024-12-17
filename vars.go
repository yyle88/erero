package erero

import (
	"github.com/yyle88/erero/eecho"
	"github.com/yyle88/erero/erzap"
)

var op = eecho.NewErrorsEcho(erzap.NewZapED(3))
