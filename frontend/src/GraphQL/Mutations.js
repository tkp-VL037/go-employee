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