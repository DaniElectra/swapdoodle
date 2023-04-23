package nex

import (
	"fmt"
	"os"

	"github.com/PretendoNetwork/swapdoodle/globals"
	"github.com/PretendoNetwork/swapdoodle/utility"

	nex "github.com/PretendoNetwork/nex-go"
)

func StartNEXServer() {
	globals.NEXServer = nex.NewServer()
	globals.NEXServer.SetDefaultNEXVersion(&nex.NEXVersion{
		Major: 2,
		Minor: 4,
		Patch: 1,
	})
	globals.NEXServer.SetAccessKey("76f26496")
	globals.NEXServer.SetPasswordFromPIDFunction(utility.PasswordFromPID)

	globals.NEXServer.On("Data", func(packet *nex.HPPPacket) {
	request := packet.RMCRequest()

	fmt.Println("== Swapdoodle - HPP ==")
	fmt.Printf("Protocol ID: %#v\n", request.ProtocolID())
	fmt.Printf("Method ID: %#v\n", request.MethodID())
	fmt.Println("======================")
	})

	registerNEXProtocols()

	globals.NEXServer.HPPListen(os.Getenv("SERVER_ADDRESS"))
}
