package models

type WorkOrder struct {
  CreatedBy string
  ClientName string
  ClientId string
  RequiredResources int
  ManagedBy string
}