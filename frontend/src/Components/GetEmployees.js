import React, { useEffect, useState } from 'react';
import {useQuery, gql} from '@apollo/client'
import {GET_ALL_EMPLOYEES} from '../GraphQL/Queries'
import GetEmployeeDetail from './GetEmployeeDetail';


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
                                <GetEmployeeDetail employeeId={emp.id}/>
                                <button type="button" className="btn btn-warning"
                                        data-bs-toggle="modal"
                                        data-bs-target={`#emp-${emp.id}UpdateModal`}>
                                        Update
                                </button>
                                <button type="button" className="btn btn-danger"
                                        data-bs-toggle="modal"
                                        data-bs-target={`#emp-${emp.id}DeleteModal`}>
                                        Delete
                                </button>

                                {/* Update Modal */}
                                <div className="modal fade" id={`emp-${emp.id}UpdateModal`} tabindex="-1" aria-labelledby={`emp-${emp.id}UpdateModalLabel`} aria-hidden="true">
                                    <div className="modal-dialog">
                                        <div className="modal-content">
                                        <div className="modal-header">
                                            <h1 className="modal-title fs-5" id={`emp-${emp.id}UpdateModalLabel`}>{emp.name}</h1>
                                            <button type="button" className="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
                                        </div>
                                        <div className="modal-body">
                                            <div className="container mt-5">
                                                <table className="table">
                                                    <tbody>
                                                    <tr>
                                                        <td>Id: <b>{emp.id}</b></td>
                                                    </tr>
                                                    <tr>
                                                        <td>
                                                            <div className="mb-3 row">
                                                                <label for="inputPosition" className="col-sm-2 col-form-label">Position</label>
                                                                <div className="col-sm-10">
                                                                    <input type="position" className="form-control" id="inputPosition" value={emp.position}/>
                                                                </div>
                                                            </div>
                                                        </td>
                                                    </tr>
                                                    <tr>
                                                        <td>View Count: {emp.viewCount}</td>
                                                    </tr>
                                                    </tbody>
                                                </table>
                                            </div>
                                        </div>
                                        <div className="modal-footer">
                                            <button type="button" className="btn btn-secondary" data-bs-dismiss="modal">Cancel</button>
                                            <button type="button" className="btn btn-success">Confirm</button>
                                        </div>
                                        </div>
                                    </div>
                                </div>

                                {/* Delete Modal */}
                                <div className="modal fade" id={`emp-${emp.id}DeleteModal`} tabindex="-1" aria-labelledby={`emp-${emp.id}DeleteModalLabel`} aria-hidden="true">
                                    <div className="modal-dialog">
                                        <div className="modal-content">
                                        <div className="modal-header">
                                            <h1 className="modal-title fs-5" id={`emp-${emp.id}DeleteModalLabel`}>Modal title</h1>
                                            <button type="button" className="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
                                        </div>
                                        <div className="modal-body">
                                            <div className="container mt-5">
                                                <h2>Delete Employee?</h2>
                                                <table className="table">
                                                    <tbody>
                                                    <tr>
                                                        <td>Position: {emp.position}</td>
                                                    </tr>
                                                    <tr>
                                                        <td>View Count: {emp.viewCount}</td>
                                                    </tr>
                                                    </tbody>
                                                </table>
                                                </div>
                                            </div>
                                            <div className="modal-footer">
                                                <button type="button" className="btn btn-secondary" data-bs-dismiss="modal">Cancel</button>
                                                <button type="button" className="btn btn-danger">Confirm</button>
                                            </div>
                                        </div>
                                    </div>
                                </div>
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