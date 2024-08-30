package main

func isValidPort(port string) bool {
	if len(port) == 0 {
		return false
	}
	for _, digit := range port {
		if digit < '0' || digit > '9' {
			return false
		}
	}
	portNum := 0
	for _, digit := range port {
		portNum = portNum*10 + int(digit-'0')
	}
	if portNum < 1 || portNum > 65535 {
		return false
	}
	return true
}
