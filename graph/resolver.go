package graph

import (
	syncforcommerce "github.com/codatio/client-sdk-go/sync-for-commerce"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	CodatClient *syncforcommerce.CodatSyncCommerce
}
