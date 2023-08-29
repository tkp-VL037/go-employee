import { gql } from "@apollo/client";

export const ADD_EMPLOYEE = gql`
    mutation AddEmployee($input: NewEmployee!) {
        addEmployee(input: $input) {
            id
            age
            position
            viewCount
        }
    }
`

export const UPDATE_EMPLOYEE = gql`
    mutation UpdateEmployee($id: ID!, $input: UpdateEmployee!) {
        updateEmployeeDetail(id: $id, input: $input) {
            id
            age
            name
            position
            viewCount
        }
    }
`