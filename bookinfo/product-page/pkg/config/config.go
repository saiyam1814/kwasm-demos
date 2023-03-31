package config

import (
	"fmt"
	"os"
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

	servicesDomainEnvVar  = "SERVICES_DOMAIN"
	detailsHostnameEnvVar = "DETAILS_HOSTNAME"
	detailsPortEnvVar     = "DETAILS_SERVICE_PORT"
	ratingsHostnameEnvVar = "RATINGS_HOSTNAME"
	ratingsPortEnvVar     = "RATINGS_SERVICE_PORT"
	reviewsHostnameEnvVar = "REVIEWS_HOSTNAME"
	reviewsPortEnvVar     = "REVIEWS_SERVICE_PORT"
)

func NewServicesConfig() *ServicesConfig {
	servicesDomain := getEnvVar(servicesDomainEnvVar, defaultServicesDomain)
	detailsHostname := getEnvVar(detailsHostnameEnvVar, defaultDetailsHostname)
	detailsPort := getEnvVar(detailsPortEnvVar, defaultDetailsPort)
	ratingsHostname := getEnvVar(ratingsHostnameEnvVar, defaultRatingsHostname)
	ratingsPort := getEnvVar(ratingsPortEnvVar, defaultRatingsPort)
	reviewsHostname := getEnvVar(reviewsHostnameEnvVar, defaultReviewsHostname)
	reviewsPort := getEnvVar(reviewsPortEnvVar, defaultReviewsPort)

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

	return &ServicesConfig{
		ProductPage: productPage,
		Details:     details,
		Reviews:     reviews,
		Ratings:     ratings,
	}
}

func getEnvVar(key string, defaultVal string) string {
	val, ok := os.LookupEnv(key)
	if ok {
		return val
	}
	return defaultVal
}
