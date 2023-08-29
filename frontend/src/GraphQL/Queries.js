import { gql } from "@apollo/client";

export const GET_ALL_EMPLOYEES = gql`
    query Employees {
        getEmployees {
            id
            age
            name
            position
            viewCount
        }
    }
`

export const GET_EMPLOYEE_DETAIL = gql`
    query GetEmployee($employeeId: ID!) {
        getEmployeeDetail(id: $employeeId) {
            id
            age
            name
            position
            viewCount
        }
    }
`