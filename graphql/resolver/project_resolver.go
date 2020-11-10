package resolver

import (
	"context"
	//"fmt"
	"github.com/golang_project_01_server/datasources"
	"github.com/golang_project_01_server/graphql/models"
)

type projectResolver struct {
	Project      *models.Project
	rootResolver *Resolver
}

func (r *projectResolver) Id() *string {
	return &r.Project.Id
}

func (r *projectResolver) Name() *string {
	return &r.Project.Name
}

func (r *projectResolver) ParentId() *string {
	return &r.Project.ParentId
}

func (r *projectResolver) StartDate() *string {
	return &r.Project.StartDate
}

func (r *projectResolver) EndDate() *string {
	return &r.Project.EndDate
}

func (r *projectResolver) Description() *string {
	return &r.Project.Description
}

func (r *projectResolver) ClientId() *string {
	return &r.Project.ClientId
}

func (r *Resolver) GetAllProjects(ctx context.Context) (*[]*projectResolver, error) {

	authorization := ctx.Value("X-Authorization")

	t := authorization.(string)

	array := datasources.GetAllProjects(t)

	projectResolvers := make([]*projectResolver, 0)

	for _, pro := range array {
		p := models.NewProject(&pro)
		projectResolvers = append(projectResolvers, &projectResolver{
			Project:      p,
			rootResolver: r,
		})
	}

	return &projectResolvers, nil

}
