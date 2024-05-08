package gateway

import "smart-pay-service/internal/domain/entity"

type CostCenter interface {
	InsertCoastCenter(obj []*entity.CostCenterInfo) (int, error)
	InsertCoastEmployee(obj []*entity.EmployeeMonthlyCosts, id int) error
	InsertCoastVariable(obj []*entity.MonthlyCosts, id int) error
}
