package cron

import (
	"github.com/filecoin-project/lotus/chain/actors"
	"github.com/ipfs/go-cid"
	"golang.org/x/xerrors"

	builtin0 "github.com/filecoin-project/specs-actors/actors/builtin"

	builtin2 "github.com/filecoin-project/specs-actors/v2/actors/builtin"

	builtin3 "github.com/filecoin-project/specs-actors/v3/actors/builtin"

	builtin4 "github.com/filecoin-project/specs-actors/v4/actors/builtin"
)

func GetActorCodeID(av actors.Version) (cid.Cid, error) {
	switch av {

	case actors.Version0:
		return builtin0.CronActorCodeID, nil

	case actors.Version2:
		return builtin2.CronActorCodeID, nil

	case actors.Version3:
		return builtin3.CronActorCodeID, nil

	case actors.Version4:
		return builtin4.CronActorCodeID, nil

	}

	return cid.Undef, xerrors.Errorf("unknown actor version %d", av)
}

var (
	Address = builtin4.CronActorAddr
	Methods = builtin4.MethodsCron
)
