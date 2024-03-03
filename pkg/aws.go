package pkg

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/eks"
)

type NodegroupUpdater struct {
	Regions    []string
	Clusters   map[string][]string
	Nodegroups map[string][]string
	Configs    map[string]aws.Config
}

func NewNodegroupUpdater(regions []string, clusters []string, nodegroups []string) *NodegroupUpdater {
	updater := &NodegroupUpdater{
		Regions: regions,
		Configs: make(map[string]aws.Config),
	}
	updater.initConfigs()
	return updater
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
	// find all clusters in the regions configured and add those clusters to a map indexed by region
	for _, region := range n.Regions {
		cfg := n.Configs[region]
		client := eks.NewFromConfig(cfg)
		input := &eks.ListClustersInput{}
		resp, err := client.ListClusters(context.TODO(), input)
		if err != nil {
			return err
		}
		for _, cluster := range resp.Clusters {
			n.Clusters[region] = append(n.Clusters[region], cluster)
		}

	}
	return nil
}

func (n *NodegroupUpdater) DiscoverNodegroups() error {

	for _, region := range n.Regions {
		for _, cluster := range n.Clusters[region] {

			cfg := n.Configs[region]
			client := eks.NewFromConfig(cfg)
			input := &eks.ListNodegroupsInput{
				ClusterName: aws.String(cluster),
			}
			resp, err := client.ListNodegroups(context.TODO(), input)

			if err != nil {
				return err
			}

			for _, nodegroup := range resp.Nodegroups {
				n.Nodegroups[region] = append(n.Nodegroups[region], nodegroup)
			}
		}
	}
	return nil
}

func (n *NodegroupUpdater) UpgradeNodegroup(cluster string, region string, force bool) error {

	return nil
}
