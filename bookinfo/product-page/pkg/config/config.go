package config

import (
	"fmt"

	spinConfig "github.com/fermyon/spin/sdk/go/config"
)

type ServicesConfig struct {
	ProductPage Endpoint
	Details     Endpoint
	Reviews     Endpoint
	Ratings     Endpoint
}

type Endpoint struct {
	Name     string
	Endpoint string
	Children []Endpoint
}

const (
	defaultServicesDomain  = ""
	defaultDetailsHostname = "details"
	defaultDetailsPort     = "9080"
	defaultRatingsHostname = "ratings"
	defaultRatingsPort     = "9080"
	defaultReviewsHostname = "reviews"
	defaultReviewsPort     = "9080"

	servicesDomainConfigKey  = "services_domain"
	detailsHostnameConfigKey = "details_hostname"
	detailsPortConfigKey     = "details_service_port"
	ratingsHostnameConfigKey = "ratings_hostname"
	ratingsPortConfigKey     = "ratings_service_port"
	reviewsHostnameConfigKey = "reviews_hostname"
	reviewsPortConfigKey     = "reviews_service_port"
)

func NewServicesConfig() *ServicesConfig {
	servicesDomain := getConfigFromKey(servicesDomainConfigKey, defaultServicesDomain)
	detailsHostname := getConfigFromKey(detailsHostnameConfigKey, defaultDetailsHostname)
	detailsPort := getConfigFromKey(detailsPortConfigKey, defaultDetailsPort)
	ratingsHostname := getConfigFromKey(ratingsHostnameConfigKey, defaultRatingsHostname)
	ratingsPort := getConfigFromKey(ratingsPortConfigKey, defaultRatingsPort)
	reviewsHostname := getConfigFromKey(reviewsHostnameConfigKey, defaultReviewsHostname)
	reviewsPort := getConfigFromKey(reviewsPortConfigKey, defaultReviewsPort)

	details := Endpoint{
		Name:     fmt.Sprintf("http://%s%s:%s", detailsHostname, servicesDomain, detailsPort),
		Endpoint: "details",
	}

	ratings := Endpoint{
		Name:     fmt.Sprintf("http://%s%s:%s", ratingsHostname, servicesDomain, ratingsPort),
		Endpoint: "ratings",
	}

	reviews := Endpoint{
		Name:     fmt.Sprintf("http://%s%s:%s", reviewsHostname, servicesDomain, reviewsPort),
		Endpoint: "reviews",
		Children: []Endpoint{
			ratings,
		},
	}

	productPage := Endpoint{
		Name:     fmt.Sprintf("http://%s%s:%s", detailsHostname, servicesDomain, detailsPort),
		Endpoint: "productpage",
		Children: []Endpoint{
			details,
			reviews,
		},
	}
	fmt.Println(details)
	return &ServicesConfig{
		ProductPage: productPage,
		Details:     details,
		Reviews:     reviews,
		Ratings:     ratings,
	}
}

func getConfigFromKey(key string, defaultVal string) string {
	val, err := spinConfig.Get(key)
	if err != nil {
		fmt.Printf("error fetching config: %s, %s\n", key, err.Error())
		return val
	}
	if val == "" {
		fmt.Printf("config %s not set\n", key)
	}
	fmt.Printf("config %s set with value %s\n", key, val)
	return val
}
