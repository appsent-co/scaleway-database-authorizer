package main

import (
	"fmt"
	"github.com/scaleway/scaleway-sdk-go/api/rdb/v1"
	"github.com/scaleway/scaleway-sdk-go/scw"
	"net"
	"os"
	"sync"
)

func main() {
	// Scaleway client config
	organizationId := os.Getenv("SCALEWAY_ORG_ID")
	accessKey := os.Getenv("SCALEWAY_ACCESS_KEY")
	secretKey := os.Getenv("SCALEWAY_SECRET_KEY")

	// Database config
	instanceId := os.Getenv("SCALEWAY_INSTANCE_ID")
	region, err := scw.ParseRegion(os.Getenv("SCALEWAY_DATABASE_REGION"))
	if err != nil {
		panic(err)
	}

	// Pod ip address
	podIp := net.ParseIP(os.Getenv("POD_IP"))
	podName := os.Getenv("POD_NAME")

	fmt.Printf("Authorizing ip address %s...", podIp)

	// Create a Scaleway client
	client, err := scw.NewClient(
		// Get your credentials at https://console.scaleway.com/account/credentials
		scw.WithDefaultOrganizationID(organizationId),
		scw.WithAuth(accessKey, secretKey),
	)
	if err != nil {
		panic(err)
	}

	databaseApi := rdb.NewAPI(client)
	_, err = databaseApi.AddInstanceACLRules(&rdb.AddInstanceACLRulesRequest{
		Region: region,
		InstanceID: instanceId,
		Rules: []*rdb.ACLRuleRequest{
			{
				IP: podIp,
				Description: fmt.Sprintf("K8S: %s", podName),
			},
		},
	})
	if err != nil {
		panic(err)
	}

	fmt.Println("Success.")

	blockForever()
}

// A deamonset will always restart.
// Let's block the program forever to avoid an infinite loop.
func blockForever() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	wg.Wait()
}
