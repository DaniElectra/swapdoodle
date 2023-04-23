package utility

import (
	"github.com/PretendoNetwork/swapdoodle/database"

	"github.com/PretendoNetwork/nex-go"
)

func PasswordFromPID(pid uint32) (string, uint32) {
	user := database.GetNEXAccountByPID(pid)

	if user == nil {
		return "", nex.Errors.RendezVous.InvalidUsername
	}

	return user["password"].(string), 0
}
