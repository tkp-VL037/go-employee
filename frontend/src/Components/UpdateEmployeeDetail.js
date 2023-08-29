import React, { useState } from "react";
import { useMutation } from "@apollo/client";
import { UPDATE_EMPLOYEE } from "../GraphQL/Mutations";

function UpdateEmployeeDetail({employee}) {
    const [name, setName] = useState(employee.name);
    const [position, setPosition] = useState(employee.position);

    const [updateEmployee] = useMutation(UPDATE_EMPLOYEE)

    const handleUpdateEmployee = () => {
        updateEmployee({
            variables: {
                id: employee.id,
                input: {
                    name: name,
                    position, position
                }
            }
        })

        window.location.reload()
    }

    return (
        <>
            <button type="button" className="btn btn-warning"
                    data-bs-toggle="modal"
                    data-bs-target={`#employee-${employee.id}UpdateModal`}>
                    Update
            </button>

            {/* Update Modal */}
            <div className="modal fade" id={`employee-${employee.id}UpdateModal`} tabindex="-1" aria-labelledby={`employee-${employee.id}UpdateModalLabel`} aria-hidden="true">
                <div className="modal-dialog">
                    <div className="modal-content">
                    <div className="modal-header">
                        <h1 className="modal-title fs-5" id={`employee-${employee.id}UpdateModalLabel`}>{employee.name}</h1>
                        <button type="button" className="btn-close" data-bs-dismiss="modal" aria-label="Close" onClick={handleUpdateEmployee}></button>
                    </div>
                    <div className="modal-body">
                        <div className="container mt-5">
                            <table className="table">
                                <tbody>
                                <tr>
                                    <td>Id: <b>{employee.id}</b></td>
                                </tr>
                                <tr>
                                    <td>
                                        <div className="mb-3 row">
                                            <label for="inputName" className="col-sm-2 col-form-label">Name</label>
                                            <div className="col-sm-10">
                                                <input type="text" className="form-control" id="inputName" value={name} onChange={(e) => setName(e.target.value)}/>
                                            </div>
                                        </div>
                                    </td>
                                </tr>
                                <tr>
                                    <td>
                                        <div className="mb-3 row">
                                            <label for="inputPosition" className="col-sm-2 col-form-label">Position</label>
                                            <div className="col-sm-10">
                                                <input type="text" className="form-control" id="inputPosition" value={position} onChange={(e) => setPosition(e.target.value)}/>
                                            </div>
                                        </div>
                                    </td>
                                </tr>
                                <tr>
                                    <td>View Count: {employee.viewCount}</td>
                                </tr>
                                </tbody>
                            </table>
                        </div>
                    </div>
                    <div className="modal-footer">
                        <button type="button" className="btn btn-secondary" data-bs-dismiss="modal">Cancel</button>
                        <button type="button" className="btn btn-success" onClick={handleUpdateEmployee}>Confirm</button>
                    </div>
                    </div>
                </div>
            </div>
        </>
    )
}

export default UpdateEmployeeDetail

