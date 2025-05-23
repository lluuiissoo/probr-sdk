package azure

import (
	"log"
	"os"

	"github.com/probr/probr-sdk/config"
	"github.com/probr/probr-sdk/utils"
)

var prefix string
var rgName string

//TenantID returns the azure Tenant in which the tests should be executed, configured by the user and may be set by the environment variable AZURE_TENANT_ID.
func TenantID() string {
	if config.GlobalConfig.CloudProviders.Azure.TenantID == "" {
		log.Printf("[ERROR] Azure connection config var not set: config.GlobalConfig.CloudProviders.Azure.TenantID")
	}
	return config.GlobalConfig.CloudProviders.Azure.TenantID
}

//ClientID returns the client (typically a service principal) that must be authorized for performing operations within the azure tenant, configured by the user and may be set by the environment variable AZURE_CLIENT_ID.
func ClientID() string {
	if config.GlobalConfig.CloudProviders.Azure.ClientID == "" {
		log.Printf("[ERROR] Azure connection config var not set: config.GlobalConfig.CloudProviders.Azure.ClientID")
	}
	return config.GlobalConfig.CloudProviders.Azure.ClientID
}

//ClientSecret returns the client secret to allow client authetication and authorization, configured by the user and may be set by the environment variable AZURE_CLIENT_SECRET.
func ClientSecret() string {
	if config.GlobalConfig.CloudProviders.Azure.ClientSecret == "" {
		log.Printf("[ERROR] Azure connection config var not set: config.GlobalConfig.CloudProviders.Azure.ClientSecret")
	}
	return config.GlobalConfig.CloudProviders.Azure.ClientSecret
}

//SubscriptionID returns the azure Subscription in which the tests should be executed, configured by the user and may be set by the environment variable AZURE_SUBSCRIPTION_ID.
func SubscriptionID() string {
	if config.GlobalConfig.CloudProviders.Azure.SubscriptionID == "" {
		log.Printf("[ERROR] Azure connection config var not set: config.GlobalConfig.CloudProviders.Azure.SubscriptionID")
	}
	return config.GlobalConfig.CloudProviders.Azure.SubscriptionID
}

//ResourceGroup returns the Probr user's azure resource group in which resurces should be created fpr testing, configured by the user and may be set by the environment variable AZURE_RESOURCE_GROUP.
func ResourceGroup() string {
	if config.GlobalConfig.CloudProviders.Azure.ResourceGroup == "" {
		log.Printf("[ERROR] Azure connection config var not set: config.GlobalConfig.CloudProviders.Azure.ResourceGroup")
	}
	return config.GlobalConfig.CloudProviders.Azure.ResourceGroup
}

//ResourceLocation returns the default location in which azure resources should be created, configured by the user and may be set by the environment variable AZURE_LOCATION.
func ResourceLocation() string {
	if config.GlobalConfig.CloudProviders.Azure.ResourceLocation == "" {
		log.Printf("[ERROR] Azure connection config var not set: config.GlobalConfig.CloudProviders.Azure.ResourceLocation")
	}
	return config.GlobalConfig.CloudProviders.Azure.ResourceLocation
}

//ManagementGroup returns an Azure Management Group which may be used for policy assignment, configured by the user and may be set by the environment variable AZURE_MANAGEMENT_GROUP.
func ManagementGroup() string {
	return config.GlobalConfig.CloudProviders.Azure.ManagementGroup
}

func randomPrefix() string {
	if prefix == "" {
		prefix = "test" + utils.RandomString(6) + ""
	}
	return "test" + utils.RandomString(6) + ""
}

func getFromEnvVar(varName string) string {
	v, b := os.LookupEnv(varName)
	if !b {
		log.Printf("[ERROR] Environment variable \"%v\" is not defined", varName)
	}
	return v
}
