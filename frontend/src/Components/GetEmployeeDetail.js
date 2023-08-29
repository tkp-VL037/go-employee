import React, { useEffect, useState } from 'react';
import {useLazyQuery} from '@apollo/client'
import {GET_EMPLOYEE_DETAIL} from '../GraphQL/Queries'


function GetEmployeeDetail({employeeId}) {
    const [employee, setEmployee] = useState({})
    const [getEmployee, {error, loading, data}] = useLazyQuery(GET_EMPLOYEE_DETAIL)

    const handleGetEmployeeDetail = () => {
        getEmployee({
            variables: {
                employeeId: employeeId
            }
        })
        if (data) {
            setEmployee(data.getEmployeeDetail)
            console.log(data.getEmployeeDetail)
        }
    }

    return (
        <div>
            <button type="button" className="btn btn-info"
                    data-bs-toggle="modal"
                    data-bs-target={`#emp-${employeeId}ViewModal`}
                    onClick={() => handleGetEmployeeDetail(employeeId)}>
                    View
            </button>

            {/* View Modal */}
            {loading && <p>Loading...</p>}
            {error && <p>Error: {error.message}</p>}
            {data && data.getEmployee && (
                <div className="modal fade" id={`emp-${employeeId}ViewModal`} tabindex="-1" aria-labelledby={`emp-${employeeId}ViewModalLabel`} aria-hidden="true">
                    <div className="modal-dialog">
                        <div className="modal-content">
                        <div className="modal-header">
                            <h1 className="modal-title fs-5" id={`emp-${employeeId}ViewModalLabel`}>{employeeId}</h1>
                            <button type="button" className="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
                        </div>
                        <div className="modal-body">
                            <div className="container mt-5">
                                <table className="table">
                                    <tbody>
                                    <tr>
                                        <td>Name: <b>{employee.name}</b></td>
                                    </tr>
                                    <tr>
                                        <td>Position: {employee.position}</td>
                                    </tr>
                                    <tr>
                                        <td>View Count: <b>{employee.viewCount}</b></td>
                                    </tr>
                                    </tbody>
                                </table>
                            </div>
                        </div>
                        </div>
                    </div>
                </div>
            )}
        </div>
    )
}

export default GetEmployeeDetail