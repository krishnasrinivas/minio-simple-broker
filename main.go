package main

import (
	"context"
	"os"

	"net/http"

	"github.com/pivotal-cf/brokerapi"
)

type minioBroker struct{}

func (b *minioBroker) Services(context context.Context) []brokerapi.Service {
	return []brokerapi.Service{
		brokerapi.Service{
			ID:          "minio simple broker id",
			Name:        "minio simple broker name",
			Description: "minio simple broker description",
			Bindable:    false,
			Plans: []brokerapi.ServicePlan{{
				ID:          "aaa",
				Name:        "bbb",
				Description: "ccc",
			}},
		},
	}
}

func (b *minioBroker) Provision(context context.Context, instanceID string,
	details brokerapi.ProvisionDetails, asyncAllowed bool) (brokerapi.ProvisionedServiceSpec, error) {
	return brokerapi.ProvisionedServiceSpec{false, "aaa", "bbb"}, nil
}

func (b *minioBroker) Deprovision(context context.Context, instanceID string,
	details brokerapi.DeprovisionDetails, asyncAllowed bool) (brokerapi.DeprovisionServiceSpec, error) {
	return brokerapi.DeprovisionServiceSpec{false, "aaa"}, nil
}

func (b *minioBroker) Bind(context context.Context, instanceID, bindingID string, details brokerapi.BindDetails) (brokerapi.Binding, error) {
	return brokerapi.Binding{}, nil
}

func (b *minioBroker) Unbind(context context.Context, instanceID, bindingID string, details brokerapi.UnbindDetails) error {
	return nil
}

func (b *minioBroker) Update(context context.Context, instanceID string, details brokerapi.UpdateDetails, asyncAllowed bool) (brokerapi.UpdateServiceSpec, error) {
	return brokerapi.UpdateServiceSpec{}, nil
}

func (b *minioBroker) LastOperation(context context.Context, instanceID, operationData string) (brokerapi.LastOperation, error) {
	return brokerapi.LastOperation{}, nil
}

func main() {
	handler := brokerapi.New(&minioBroker{}, nil, brokerapi.BrokerCredentials{os.Getenv("SECURITY_USER_NAME"), os.Getenv("SECURITY_USER_PASSWORD")})
	http.ListenAndServe(":8080", handler)
}
