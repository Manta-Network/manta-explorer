package util

import (
	"os"

	"github.com/BurntSushi/toml"
	"github.com/go-kratos/kratos/pkg/log"
)

type EnvironmentConfig struct {
	EventStorageKey    string
	AddressType        string
	BalanceAccuracy    string
	CommissionAccuracy string
	WSEndPoint         string
	NetworkNode        string
	IsProduction       bool
}

var config EnvironmentConfig

var (
	CurrentRuntimeSpecVersion int
	EventStorageKey           = GetEnv("SUBSTRATE_EVENT_KEY", "0x26aa394eea5630e07c48ae0c9558cef780d41e5e16056765bc8461851072c9d7")
	AddressType               = GetEnv("SUBSTRATE_ADDRESS_TYPE", "42")
	BalanceAccuracy           = GetEnv("SUBSTRATE_ACCURACY", "9")
	CommissionAccuracy        = GetEnv("COMMISSION_ACCURACY", "9")
	WSEndPoint                = GetEnv("CHAIN_WS_ENDPOINT", "wss://ws.f1.testnet.manta.network")
	NetworkNode               = GetEnv("NETWORK_NODE", "manta")
	IsProduction              = false
)

func GetEnvConfig() {
	filePath := "../configs/env.toml"
	if _, err := toml.DecodeFile(filePath, &config); err != nil {
		log.Error("Toml ERROR", err)
		panic(err)
	}
}

func SetEnv() {
	WSEndPoint = GetEnv("CHAIN_WS_ENDPOINT", config.WSEndPoint)
	NetworkNode = GetEnv("NETWORK_NODE", config.NetworkNode)
}

func GetEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		value = defaultValue
	}
	return value
}
