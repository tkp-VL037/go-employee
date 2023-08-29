import React, { useEffect, useState } from 'react';
import {useQuery, gql} from '@apollo/client'
import {GET_ALL_EMPLOYEES} from '../GraphQL/Queries'
import GetEmployeeDetail from './GetEmployeeDetail';
import UpdateEmployeeDetail from './UpdateEmployeeDetail';
import DeleteEmployee from './DeleteEmployee';

function GetEmployees() {
    const {error, loading, data} = useQuery(GET_ALL_EMPLOYEES)
    const [employees, setEmployees] = useState([])

    useEffect(() => {
        if (data) {
            setEmployees(data.getEmployees)
        }
    }, [data])
    
    return (
        <div className="container mt-5">
            <h2>List of Employees</h2>
            <table className="table">
                <thead>
                    <tr>
                    <th scope="col">ID</th>
                    <th scope="col">Name</th>
                    <th scope="col">Position</th>
                    <th scope="col">View Count</th>
                    <th scope="col">Actions</th>
                    </tr>
                </thead>
                <tbody>
                    {employees.map((emp) => {
                        return (
                            <tr>
                            <td>{emp.id}</td>
                            <td>{emp.name}</td>
                            <td>{emp.position}</td>
                            <td>{emp.viewCount}</td>
                            <td key={emp.id}>
                                <GetEmployeeDetail employeeId={emp.id} />
                                <UpdateEmployeeDetail employee={emp} />
                                <DeleteEmployee employee={emp} />
                            </td>
                            </tr>
                        )
                    })}
                </tbody>
            </table>
        </div>
    )
}

export default GetEmployees