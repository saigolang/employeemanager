package cache

import (
	"employeemanager/pkg/structs/request"
	"employeemanager/pkg/structs/response"
	"sync"
)

var EmployeesResponse []response.FinalResponse

type Store interface {
	SetEmployees(employee []response.FinalResponse)
	GetEmployeeByRequest(request request.Request) (employee response.FinalResponse)
}

var (
	once     sync.Once
	instance Store
)

func GetCache() Store {
	once.Do(func() {
		instance = newStore()
	})
	return instance
}

type cache struct {
	employees []response.FinalResponse
}

func newStore() Store {
	return &cache{
		employees: EmployeesResponse,
	}
}

func (c cache) SetEmployees(employee []response.FinalResponse) {
	EmployeesResponse = employee
}

func (c cache) GetEmployeeByRequest(request request.Request) (employee response.FinalResponse) {
	return
}
