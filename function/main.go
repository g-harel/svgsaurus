package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/g-harel/svgsaurus"
)

// Handler responds to http requests.
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	svg, err := new(svgsaurus.Config).FromQuery(request.QueryStringParameters).Render()
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	return events.APIGatewayProxyResponse{
		Body:       string(svg),
		StatusCode: 200,
		Headers: map[string]string{
			"Access-Control-Allow-Origin": "*",
			"Content-Type":                "image/svg+xml",
			"Cache-Control":               "public, max-age=900, s-maxage=900",
		},
	}, nil

}

func main() {
	lambda.Start(Handler)
}
