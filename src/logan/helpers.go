package logan

import "os"

const (
	SysLog = "syslog"
	Nats   = "nats"
)

func getStoreMethod() string {
	if retVal := os.Getenv("LOGAN_STORE_TYPE"); retVal != "" {
		return retVal
	}
	return SysLog
}

func getServiceName() string {
	if retVal := os.Getenv("LOGAN_SERVICE_NAME"); retVal != "" {
		return retVal
	}
	return "unknown"
}
