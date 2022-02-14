package accounts

import (
	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy"
	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy/resources"
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/non-standard/validators"
	uuid "github.com/google/uuid"
)

// AzureSubscriptionAccount represents an Azure subscription account.
type AzureSubscriptionAccount struct {
	AzureEnvironment      string `validate:"omitempty,oneof=AzureCloud AzureChinaCloud AzureGermanCloud AzureUSGovernment"`
	CertificateBytes      *resources.SensitiveValue
	CertificateThumbprint string
	ManagementEndpoint    string     `validate:"omitempty,uri"`
	StorageEndpointSuffix string     `validate:"omitempty,hostname"`
	SubscriptionID        *uuid.UUID `validate:"required"`

	Account
}

// NewAzureSubscriptionAccount creates and initializes an Azure subscription
// account with a name.
func NewAzureSubscriptionAccount(name string, subscriptionID uuid.UUID, options ...func(*AzureSubscriptionAccount)) (*AzureSubscriptionAccount, error) {
	if octopusdeploy.IsEmpty(name) {
		return nil, octopusdeploy.CreateRequiredParameterIsEmptyOrNilError(octopusdeploy.ParameterName)
	}

	account := AzureSubscriptionAccount{
		Account: *NewAccount(name, AccountType("AzureSubscription")),
	}

	// iterate through configuration options and set fields (without checks)
	for _, option := range options {
		option(&account)
	}

	// assign pre-determined values to "mandatory" fields
	account.AccountType = AccountType("AzureSubscription")
	account.Name = name
	account.SubscriptionID = &subscriptionID

	// validate to ensure that all expectations are met
	err := account.Validate()
	if err != nil {
		return nil, err
	}

	return &account, nil
}

// Validate checks the state of this account and returns an error if invalid.
func (a *AzureSubscriptionAccount) Validate() error {
	v := validator.New()
	err := v.RegisterValidation("notblank", validators.NotBlank)
	if err != nil {
		return err
	}
	err = v.RegisterValidation("notall", resources.NotAll)
	if err != nil {
		return err
	}
	return v.Struct(a)
}
