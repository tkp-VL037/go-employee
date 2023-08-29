import React, { useEffect, useState } from 'react';
import {useLazyQuery} from '@apollo/client'
import {GET_EMPLOYEE_DETAIL} from '../GraphQL/Queries'


function GetEmployeeDetail({employeeId}) {
    const [employee, setEmployee] = useState({})
    var [getEmployee, {error, loading, data}] = useLazyQuery(GET_EMPLOYEE_DETAIL)

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

    useEffect(() => {
        if (data) {
          setEmployee(data.getEmployeeDetail);
        }
      }, [data]);
    

    const handleRefresh = () => {
        window.location.reload()
    }

    return (
        <div>
            <button type="button" class="btn btn-primary" data-bs-toggle="modal" data-bs-target={`#emp-${employeeId}-view`}
                onClick={() => handleGetEmployeeDetail(employeeId)}>
            View
            </button>
            <div class="modal fade" id={`emp-${employeeId}-view`}>
                <div class="modal-dialog">
                    <div class="modal-content">
                        <div className="modal-header">
                            <h1 className="modal-title fs-5" id={`emp-${employeeId}ViewModalLabel`}>{employeeId}</h1>
                            <button type="button" className="btn-close" data-bs-dismiss="modal" aria-label="Close" onClick={handleRefresh}></button>
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
        </div>
    )
}

export default GetEmployeeDetail