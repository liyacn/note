package apollo

import (
	"github.com/apolloconfig/agollo/v4"
	"github.com/apolloconfig/agollo/v4/env/config"
	"github.com/apolloconfig/agollo/v4/storage"
	"log"
)

type Config struct {
	AppID     string
	Cluster   string
	Namespace string
	IP        string
	Secret    string
}

func NewClient(cfg *Config) agollo.Client {
	client, err := agollo.StartWithConfig(func() (*config.AppConfig, error) {
		return &config.AppConfig{
			AppID:         cfg.AppID,
			Cluster:       cfg.Cluster,
			NamespaceName: cfg.Namespace,
			IP:            cfg.IP,
			Secret:        cfg.Secret,
		}, nil
	})
	if err != nil {
		log.Fatal(err)
	}
	client.AddChangeListener(listener{})
	return client
}

type listener struct{}

// OnChange 配置变更时会触发
func (listener) OnChange(event *storage.ChangeEvent) {
	log.Printf("OnChange: %+v\n", event.Changes)
}

// OnNewestChange 启动加载和配置变更时都会触发
func (listener) OnNewestChange(event *storage.FullChangeEvent) {
	log.Printf("OnNewestChange: %v\n", event.Changes)
}
