package utils

import (
	"strconv"
	"strings"

	"github.com/google/uuid"
)

func GenerateCliTokenUUID(userId int) string {
	newUUID := uuid.New()

	uuidString := strings.ReplaceAll((newUUID).String(), "", "")
	return strconv.Itoa(userId) + "clitoken" + uuidString
}
