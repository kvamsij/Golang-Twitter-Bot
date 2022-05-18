package middleware

import (
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type Credentials struct {
	Token  string `json:"token"`
	IsProd bool   `json:"isProd"`
}

func getCredentials() Credentials {
	var credentials Credentials
	filePath, _ := filepath.Abs("./.env")
	data, err := os.ReadFile(filePath)
	if err != nil {
		panic("Error occured while loading config data.")
	}
	envVariable := strings.Split(string(data), "\n")

	for _, variable := range envVariable {
		if strings.Split(variable, "=")[0] == "BearerToken" {
			credentials.Token = strings.Split(variable, "=")[1]
		}
		credentials.IsProd, _ = strconv.ParseBool(strings.Split(variable, "=")[1])
	}
	return credentials
}
