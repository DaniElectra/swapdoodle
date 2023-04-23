package nex

import (
	"github.com/PretendoNetwork/nex-protocols-go/datastore"
	"github.com/PretendoNetwork/swapdoodle/globals"
)

func registerNEXProtocols() {
	_ = datastore.NewDataStoreProtocol(globals.NEXServer)
}
