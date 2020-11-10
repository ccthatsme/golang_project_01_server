package models

type Project struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	ParentId    string `json:"parentId"`
	StartDate   string `json:"startDate"`
	EndDate     string `json:"endDate"`
	Description string `json:"description"`
	ClientId    string `json:"clientId"`
}

func NewProject(pro *Project) *Project {

	return &Project{
		Id:          pro.Id,
		Name:        pro.Name,
		ParentId:    pro.ParentId,
		StartDate:   pro.StartDate,
		EndDate:     pro.EndDate,
		Description: pro.Description,
		ClientId:    pro.ClientId,
	}
}
