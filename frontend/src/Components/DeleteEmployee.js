import React, { useState } from "react";
import { useMutation } from "@apollo/client";
import { DELETE_EMPLOYEE } from "../GraphQL/Mutations";

function DeleteEmployees({employee}) {
    const [deleteEmployee] = useMutation(DELETE_EMPLOYEE)

    const handleDeleteEmployee = () => {
        deleteEmployee({
            variables: {
                id: employee.id
            }
        })
        window.location.reload()
    }
    
    return (
        <div>
            <button type="button" className="btn btn-danger"
                    data-bs-toggle="modal"
                    data-bs-target={`#emp-${employee.id}DeleteModal`}>
                    Delete
            </button>

            <div className="modal fade" id={`emp-${employee.id}DeleteModal`} tabindex="-1" aria-labelledby={`emp-${employee.id}DeleteModalLabel`} aria-hidden="true">
                <div className="modal-dialog">
                    <div className="modal-content">
                    <div className="modal-header">
                        <h1 className="modal-title fs-5" id={`emp-${employee.id}DeleteModalLabel`}>{employee.id}</h1>
                        <button type="button" className="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
                    </div>
                    <div className="modal-body">
                        <div className="container mt-5">
                            <h2>Delete <b>{employee.name}</b>?</h2>
                            <table className="table">
                                <tbody>
                                <tr>
                                    <td>Position: {employee.position}</td>
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
                            <button type="button" className="btn btn-danger" onClick={handleDeleteEmployee}>Confirm</button>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    )
}

export default DeleteEmployees