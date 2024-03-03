package pkg


import (
"github.com/aws/aws-sdk-go-v2/aws"
"github.com/aws/aws-sdk-go-v2/config"
"github.com/aws/aws-sdk-go-v2/service/eks"
)


type NodegroupUpdater struct {
	Regions  []string
	Clusters map[string]string
	Nodegroups []string
	Config map[string]aws.Config
}

func NewNodegroupUpdater(regions []string, clusters []string, nodegroups []string) *NodegroupUpdater {
	updater := &NodegroupUpdater{
		Regions: regions,
		Clusters: clusters,
		Nodegroups: nodegroups,
		Config: make(map[string]aws.Config),
	}
	updater.initConfigs()
}

func (n *NodegroupUpdater) initConfigs() {
	for _, region := range n.Regions {
		cfg, err := config.LoadDefaultConfig(context.TODO(),
			config.WithRegion(region),
		)
		if err != nil {
			log.Fatalf("Unable to load AWS config for region %s, %v", region, err)
		}
		n.Configs[region] = cfg
	}
}

func (n *NodegroupUpdater) DiscoverClusters() error {
	for _, region := range n.Regions {
		cfg := n.Configs[region]
		client := eks.NewFromConfig(cfg)
		input := &eks.ListClustersInput{}
		req := client.ListClustersRequest(input)
		resp, err := req.Send(context.TODO())
		if err != nil {
			return err
		}
		for _, cluster := range resp.Clusters {
			n.Clusters[cluster] = region
		}
	}
}

func (n *NodegroupUpdater) DiscoverNodegroups(cluster string, region string) error {
	cfg := n.Configs[region]
	client := eks.NewFromConfig(cfg)
	input := &eks.ListNodegroupsInput{
		ClusterName: aws.String(cluster),
	}
	req := client.ListNodegroupsRequest(input)
	resp, err := req.Send(context.TODO())
	if err != nil {
		return err
	}
	for _, nodegroup := range resp.Nodegroups {
		n.Nodegroups = append(n.Nodegroups, nodegroup)
	}

}

func (n *NodegroupUpdater) ListUpdatesPending(cluster string, region string) error {
	cfg := n.Configs[region]
	client := eks.NewFromConfig(cfg)
	input := &eks.DescribeNodegroupInput{
		ClusterName: aws.String(cluster),
		NodegroupName: aws.String(nodegroup),
	}
	req := client.DescribeNodegroupRequest(input)
	resp, err := req.Send(context.TODO())
	if err != nil {
		return err
	}
	// Check if there are updates pending

	
}




func (n *NodegroupUpdater) UpdateNodegroup(nodegroup string, region string) error {
	for _, region := range n.Regions {
		for _, cluster := range n.Clusters {
			for _, nodegroup := range n.Nodegroups {
				// Update the nodegroup
			}
		}
	}
}

func main(){
	for _, region := range regions {
		for _, cluster := range clusters {
			for _, nodegroup := range nodegroups {
				// Update the nodegroup
			}
		}
	}




	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion("us-west-2"),

	client := eks.NewFromConfig(cfg)

}