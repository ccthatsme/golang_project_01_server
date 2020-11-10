# golang_project_01_server

Hello!

To run the graphql server, run the main.go file in the root of this project
Authentication mutation must be ran first before trying to use the other queries

GraphQl related files are in the graphql directory(schema, resolvers, models, middleware) 

-practice directory has files that I used to test things out

-routes is older code where we used json instead of Nexient APIs, and old server

Schema:

schema {
    query: Query
    mutation: Mutation
}

type Query {
    getAllEmployees: [Employee]
    getAllProjects: [Project]
    getEmployee(id: String!): Employee
    getProject(id: String!): Project
}

type Mutation {
    authenticate(input: User): AuthResponse
}

type Employee {
    Id: String
    EmployeeNetworkId: String
    DisplayName: String
    Email: String
    Department: String
    HireDate: String
    JobTitle: String
    Manager: String
    Practice: String
}

input User {
    username: String!
    password: String!

}

type AuthResponse {
    AccessToken:String
    DisplayName: String
    RefreshToken: String
    Username: String
    Groups:[String!]
}



type Project{
    Id: String
    Name: String
    ParentId: String
    StartDate: String
    EndDate:String
    Description: String
    ClientId: String
}
