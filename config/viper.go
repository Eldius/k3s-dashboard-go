package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	prometheusEndpointKey = "prometheus.endpoint"
)

func Setup(cfgFile string) {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".k3s-dashboard-go" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".k3s-dashboard-go")
	}

	viper.SetDefault(prometheusEndpointKey, "http://prometheus-kube-prometheus-prometheus.prometheus.svc.cluster.local:9090")

	viper.SetEnvPrefix("dashboard")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
