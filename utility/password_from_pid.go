package utility

import "github.com/PretendoNetwork/swapdoodle/database"

func PasswordFromPID(pid uint32) string {
	user := database.GetNEXAccountByPID(pid)

	if user == nil {
		return ""
	}

	return user["password"].(string)
}
