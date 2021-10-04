package client

import "flag"

var (
	backendURIFlag = flag.String("backend", "http://localhost:8080", "Backend API URI")
	helpFlag = flag.Bool("help", false, "Display help prompt")
)

func main() {
	flag.Parse()
}
