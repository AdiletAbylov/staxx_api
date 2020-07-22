package staxxapi

var serviceURL string
var servicePort string

// Init initializes with url and  port
func Init(url string, port string) {
	serviceURL = url
	servicePort = port
}

// ConnectionString returns string containing url and port to the service
func connectionString() string {
	return serviceURL + ":" + servicePort
}
