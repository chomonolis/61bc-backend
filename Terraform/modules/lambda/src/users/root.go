package users

import (
	"fmt"
	"strings"

	"github.com/aws/aws-lambda-go/events"
)

func Root(request events.APIGatewayProxyRequest) (any, int, error) {
	path := request.Path
	trm := strings.Trim(path, "/")
	pathArray := strings.Split(trm, "/")
	if len(pathArray) < 1 || pathArray[0] != "users" {
		return nil, 400, fmt.Errorf("path is not allowed")
	}

	pathArray = pathArray[1:]
	if len(pathArray) == 0 {
		method := request.HTTPMethod
		switch method {
		case "POST":
			return post(request)
		}
		return nil, 400, fmt.Errorf("method %s is not allowed", method)
	}
	return nil, 400, fmt.Errorf("resource is not allowed")
}
