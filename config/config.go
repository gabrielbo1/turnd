package config

import "os"

type EnvironmentVariable string

const (
	//PublicIp - IP Address that TURN can be contacted by.
	PublicIp EnvironmentVariable = "PUBLIC_IP"
	//Port - Listening port.
	Port EnvironmentVariable = "PORT"
	//Users - List of username and password (e.g. "user=pass,user=pass")
	Users EnvironmentVariable = "USERS"
	//Realm - Realm (defaults to "pion.ly") WWW-Authenticate: Basic realm="WallyWorld"
	Realm EnvironmentVariable = "REALM"
)

func getVariable(name EnvironmentVariable, defaultValue string) string {
	if variable := os.Getenv(string(name)); variable != "" {
		return variable
	}
	return defaultValue
}

// EnvironmentVariableValue - Find to environment variable value
// or return default value of variable.
func EnvironmentVariableValue(variable EnvironmentVariable) string {
	switch variable {
	case PublicIp:
		return getVariable(PublicIp, "")
	case Port:
		return getVariable(Port, "3478")
	case Users:
		return getVariable(Users, "")
	case Realm:
		return getVariable(Realm, "pion.ly")
	}
	return ""
}
