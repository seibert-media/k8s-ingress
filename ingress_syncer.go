package k8s_ingress

import "github.com/seibert-media/k8s-ingress/model"

type ingressFetcher interface {
	Fetch() ([]model.Domain, error)
}

type ingressSyncer struct {
	fetcher ingressFetcher
}

func (i *ingressSyncer) Sync() error {
	var _, err = i.fetcher.Fetch()
	return err
}
