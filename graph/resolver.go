package graph

import "github.com/brandaogabriel7/studio-sol-backend-test-2/src/football"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{
	GameScoreService *football.GameScoreService
}
