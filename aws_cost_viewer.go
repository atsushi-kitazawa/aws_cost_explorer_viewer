package main

import (
	"fmt"
	"log"
	"os"
	"sort"

	"github.com/atsushi-kitazawa/aws_cost_explorer_viewer/setting"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/costexplorer"
	"github.com/aws/aws-sdk-go-v2/service/costexplorer/types"
	"golang.org/x/net/context"
)

var metrics = []string{"BLENDED_COST", "UNBLENDED_COST"}
//var metrics = []string{"BLENDED_COST", "UNBLENDED_COST", "USAGE_QUANTITY"}
//var metrics = []string{"NORMALIZED_USAGE_AMOUNT"}
var start string
var end string

func main() {
	// parse program arguments
	start = os.Args[1]
	end = os.Args[2]

	// load credentials
	credentials := setting.LoadCredential()

	// get client from credential
	clients := make([]*costexplorer.Client, 0)
	names := make([]string, 0)
	for _, cred := range credentials {
		clients = append(clients, getClient(cred))
		names = append(names, cred.Name)
	}

	// get cost
	for i, client := range clients {
		res, err := client.GetCostAndUsage(context.TODO(), &costexplorer.GetCostAndUsageInput{
			TimePeriod: &types.DateInterval{
				Start: aws.String(start),
				End:   aws.String(end),
			},
			Granularity: types.GranularityMonthly,
			Metrics:     metrics,
			GroupBy: []types.GroupDefinition{
				types.GroupDefinition{
					Type: types.GroupDefinitionTypeDimension,
					Key:  aws.String("SERVICE"),
				},
			},
		})
		if err != nil {
			log.Fatal(err)
		}

		// print cost
		fmt.Printf("User: %s\n", names[i])
		for _, val := range res.ResultsByTime {
			for _, g := range val.Groups {
				fmt.Println(g.Keys)
				printMetrics(g.Metrics)
			}
			fmt.Println(val.Total)
		}
	}
}

func getClient(cred setting.Credential) *costexplorer.Client {
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(cred.Apikey, cred.Secretkey, "")),
		config.WithRegion("ap-northeast-1"),
	)
	if err != nil {
		log.Fatal(err)
	}

	client := costexplorer.NewFromConfig(cfg)
	return client
}

func printMetrics(m map[string]types.MetricValue) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		fmt.Printf(" - %s %v %v\n", k, *m[k].Amount, *m[k].Unit)
	}
	return keys
}
