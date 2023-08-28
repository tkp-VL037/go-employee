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