package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"

	"github.com/atsushi-kitazawa/aws_cost_explorer_viewer/setting"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/costexplorer"
	"github.com/aws/aws-sdk-go-v2/service/costexplorer/types"
	"golang.org/x/net/context"
)

// target metrics
var metrics = []string{"BLENDED_COST", "UNBLENDED_COST", "USAGE_QUANTITY"}

// target priod
var start string
var end string

// summary metric each user
var metricSum []metricSummary

type costClient struct {
	region string
	client *costexplorer.Client
}

type metricSummary struct {
	user          string
	blendedCost   float64
	unBlendedCost float64
	usageQuantity float64
}

func main() {
	// parse program arguments
	start = os.Args[1]
	end = os.Args[2]

	// load settings
	regions, credentials := setting.LoadSettings()
	//fmt.Println(len(regions.R))
	//fmt.Println(len(credentials.C))
	//return

	// summarize metrics by user
	metricSum = make([]metricSummary, 0)

	// get cost
	for _, cred := range credentials.C {
		ms := metricSummary{cred.Name, 0, 0, 0}
		for _, region := range regions.R {
			client := getClient(cred, region)
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

			//fmt.Printf("[User:Resion]-[%s:%s]\n", cred.Name, region)
			for _, val := range res.ResultsByTime {
				for _, g := range val.Groups {
					//fmt.Println(g.Keys)
					//printMetrics(g.Metrics)
					summarizeMetric(g.Metrics, &ms)
				}
			}
		}
		metricSum = append(metricSum, ms)
	}
	for _, v := range metricSum {
		fmt.Println(v.user)
		fmt.Println(v.blendedCost)
		fmt.Println(v.unBlendedCost)
		fmt.Println(v.usageQuantity)
	}
}

func getClient(cred setting.Credential, region string) *costexplorer.Client {
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(cred.Apikey, cred.Secretkey, "")),
		config.WithRegion(region),
	)
	if err != nil {
		log.Fatal(err)
	}

	client := costexplorer.NewFromConfig(cfg)
	return client
}

func summarizeMetric(m map[string]types.MetricValue, s *metricSummary) {
	for k := range m {
		amount, _ := strconv.ParseFloat(*m[k].Amount, 64)
		switch k {
		case "BlendedCost":
			s.blendedCost = s.blendedCost + amount
		case "UnblendedCost":
			s.unBlendedCost = s.unBlendedCost + amount
		case "UsageQuantity":
			s.usageQuantity = s.usageQuantity + amount
		default:
		}
	}
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
