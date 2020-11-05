package resolver

import (
	"context"
	"github.com/golang_project_01_server/datasources"
	"github.com/golang_project_01_server/graphql/models"
)

type employeeResolver struct {
	Employee *models.Employee
}

func (r *employeeResolver) Id() *string {
	return &r.Employee.Id
}

func (r *employeeResolver) DisplayName() *string {
	return &r.Employee.DisplayName
}

func (r *employeeResolver) Email() *string {
	return &r.Employee.Email
}

func (r *Resolver) GetAllEmployees(ctx context.Context) (*[]*employeeResolver, error) {

	authorization := ctx.Value("X-Authorization")
	t := authorization.(string)

	array := datasources.GetAllEmployees(t)

	employeeResolvers := make([]*employeeResolver, 0)

	for _, emp := range array {
		employeeResolvers = append(employeeResolvers, &employeeResolver{
			Employee: emp,
		})
	}

	return &employeeResolvers, nil

}

// type Query {
//     user(uid: ID!): User
//     userByName(name: String!): User
//     users: [User]
//     cpu(id: ID!): CPU
//     cpus: [CPU]
//     memory: Memory
//     ifaces: [Iface]
//     iface(name: String!): Iface
//     service(name: String!): Service
// }

// type User struct {
// 	UID    string   `json:"uid"`
// 	GID    string   `json:"gid"`
// 	Name   string   `json:"name"`
// 	Home   string   `json:"home"`
// 	Groups []*Group `json:"groups"`
// }

// func ListUsers() ([]string, error) {
// 	file, err := os.Open(etcPasswdPath)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer file.Close()
//
// 	users := make([]string, 0)
//
// 	reader := bufio.NewReader(file)
// 	for {
// 		line, err := reader.ReadString('\n')
// 		if err == io.EOF {
// 			break
// 		}
// 		if err != nil {
// 			return nil, err
// 		}
//
// 		if strings.Index(line, "#") == 0 {
// 			continue
// 		}
//
// 		split := strings.Split(line, ":")
// 		if len(split) > 0 {
// 			users = append(users, split[0])
// 		}
// 	}
//
// 	return users, nil
// }
