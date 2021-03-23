package provider

import (
	"context"
)

func (client *ProviderClient) Watcher(ctx context.Context, outCh chan<- *DeploymentItem, errCh chan error) error {
	return nil
}

func (client *ProviderClient) Bidder(ctx context.Context, inCh <-chan *DeploymentItem, outCh chan<- *DeploymentItem, errCh chan error) error {
	return nil
}

func (client *ProviderClient) Publisher(ctx context.Context, inCh <-chan *DeploymentItem, outCh chan<- *DeploymentItem, errCh chan error) error {
	return nil
}

func (client *ProviderClient) Fulfiller(ctx context.Context, inCh <-chan *DeploymentItem, outCh chan<- *DeploymentItem, errCh chan error) error {
	return nil
}

func (client *ProviderClient) Finisher(ctx context.Context, inCh <-chan *DeploymentItem, errCh chan error) error {
	return nil
}
