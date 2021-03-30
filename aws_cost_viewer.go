package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strconv"
	"time"

	"github.com/atsushi-kitazawa/aws_cost_explorer_viewer/gui"
	"github.com/atsushi-kitazawa/aws_cost_explorer_viewer/model"
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

// verbose mode
var verbose bool

// summary metric each user
var metricSum []model.MetricSummary

type costClient struct {
	region string
	client *costexplorer.Client
}

func init() {
	flag.BoolVar(&verbose, "verbose", false, "Print a verbose message")
	flag.StringVar(&start, "start", "", "target start date")
	flag.StringVar(&end, "end", "", "target end date")
	flag.Parse()

	if verbose {
		log.SetOutput(os.Stderr)
	} else {
		log.SetOutput(ioutil.Discard)
	}
}

func main() {
	// set target period 
	if len(start) == 0 || len(end) == 0 {
	    now := time.Now()
	    nowyyyyMMdd := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, time.UTC)
	    start = nowyyyyMMdd.Format("2006-01-02")
	    end = nowyyyyMMdd.AddDate(0, 1, -1).Format("2006-01-02")
	}
	log.Printf("start=%s, end=%s", start, end)

	// load settings
	regions, credentials := setting.LoadSettings()
	log.Printf("[region] = %v", regions.R)
	log.Printf("[credential] = %v", credentials.C)

	// summarize metrics by user
	metricSum = make([]model.MetricSummary, 0)

	// get cost
	for _, cred := range credentials.C {
		ms := model.MetricSummary{cred.Name, 0, 0, 0}
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

			log.Printf("[User:Resion]-[%s:%s]\n", cred.Name, region)
			for _, val := range res.ResultsByTime {
				for _, g := range val.Groups {
					log.Println(g.Keys)
					printMetrics(g.Metrics)
					summarizeMetric(g.Metrics, &ms)
				}
			}
		}
		metricSum = append(metricSum, ms)
	}

	// display use tview
	//gui.InitTview()
	//gui.InitTable()
	//gui.DisplayMetricSummary(metricSum)

	// display table view
	gui.DisplayMetricSummaryTableView(metricSum)
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

func summarizeMetric(m map[string]types.MetricValue, s *model.MetricSummary) {
	for k := range m {
		amount, _ := strconv.ParseFloat(*m[k].Amount, 64)
		switch k {
		case "BlendedCost":
			s.BlendedCost = s.BlendedCost + amount
		case "UnblendedCost":
			s.UnBlendedCost = s.UnBlendedCost + amount
		case "UsageQuantity":
			s.UsageQuantity = s.UsageQuantity + amount
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
		log.Printf(" - %s %v %v\n", k, *m[k].Amount, *m[k].Unit)
	}
	return keys
}
