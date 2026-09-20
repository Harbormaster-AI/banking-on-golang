package models

import "github.com/google/uuid"

// Get request
type GetRequest struct {
    Id uuid.UUID `json:"id"`
}

// Delete request
type DeleteRequest struct {
    Id uuid.UUID `json:"id"`
}

// Assign request
type AssignRequest struct {
    ParentId uuid.UUID `json:"parentId"`
    ChildId  uuid.UUID `json:"childId"`
}

// Unassign request
type UnassignRequest struct {
    Id uuid.UUID `json:"Id"`
}

// AddTo request
type AddToRequest struct {
    ParentId uuid.UUID   `json:"parentId"`
    ChildIds []uuid.UUID `json:"childIds"`
}

// RemoveFrom request
type RemoveFromRequest struct {
    ParentId uuid.UUID   `json:"parentId"`
    ChildIds []uuid.UUID `json:"childIds"`
}