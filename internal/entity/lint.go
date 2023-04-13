// Package entity defines main entities for business logic (services), data base mapping and
// HTTP response objects if suitable. Each logic group entities in own file.
package entity

// Lint TODO
type Lint struct {
	ProjectId      int64  `json:"project_id"      example:"1"`
	OrganizationId int64  `json:"organization_id" example:"1"`
	Rule           string `json:"rule"        example:"json string"`
}

// id serial PRIMARY KEY,
// project bigint,
// organization bigint,
// rule text
