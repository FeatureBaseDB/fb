package fb

type WorkerServiceProviderID string

// type WorkerServiceProvider struct {
// 	ID          WorkerServiceProviderID `json:"id"`
// 	Roles       RoleTypes               `json:"roles"`
// 	Address     Address                 `json:"address"`
// 	Description string                  `json:"description"`
// }

// type WorkerServiceProviders []WorkerServiceProvider

type WorkerServiceID string

type WorkerServiceStatus string

const (
	WorkerServiceStatusActive   WorkerServiceStatus = "ACTIVE"
	WorkerServiceStatusDraining WorkerServiceStatus = "DRAINING"
	WorkerServiceStatusInactive WorkerServiceStatus = "INACTIVE"
)

// type WorkerService struct {
// 	ID                      WorkerServiceID         `json:"id"`
// 	Roles                   RoleTypes               `json:"roles"`
// 	WorkerServiceProviderID WorkerServiceProviderID `json:"worker-service-provider-id"`
// 	WorkerServiceProvider   WorkerServiceProvider   `json:"worker-service-provider"`
// 	DatabaseID              DatabaseID              `json:"database-id"`
// 	WorkersMin              int                     `json:"workers-min"`
// 	WorkersMax              int                     `json:"workers-max"`
// 	Status                  WorkerServiceStatus     `json:"status"`
// }

// type WorkerServices []WorkerService
