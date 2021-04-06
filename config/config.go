package config

import "github.com/spf13/viper"

func GetPrometheusEndpoint() string {
	return viper.GetString(prometheusEndpointKey)
}
