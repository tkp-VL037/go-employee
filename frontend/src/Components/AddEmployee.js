import React, { useState } from "react";
import { useMutation } from "@apollo/client";
import { ADD_EMPLOYEE } from "../GraphQL/Mutations";


function GetEmployeeDetail() {
    const [showModal, setShowModal] = useState(false);
    const [name, setName] = useState("");
    const [age, setAge] = useState("");
    const [position, setPosition] = useState("");
    const [addEmployee] = useMutation(ADD_EMPLOYEE);

    const handleAddEmployee = () => {
        setShowModal(true);
    };

    const handleModalClose = () => {
        setShowModal(false);
    };

    const handleFormSubmit = () => {
        addEmployee({
        variables: {
            input: {
                name: name,
                age: parseInt(age),
                position: position
            }
        }
        });

        setShowModal(false);
        setName("");
        setAge("");
        setPosition("");

        window.location.reload();
    };

    return (
        <div>
            <button className="btn btn-success btn-floating position-absolute bottom-0 end-0"
                onClick={handleAddEmployee}>
                +
            </button>

            {/* Modal */}
            {showModal && (
                <div className="modal fade show" tabIndex="-1" role="dialog" style={{ display: "block" }}>
                <div className="modal-dialog" role="document">
                    <div className="modal-content">
                    <div className="modal-header">
                        <h5 className="modal-title">Add New Employee</h5>
                        <button
                        type="button"
                        className="close"
                        onClick={handleModalClose}
                        >
                        <span aria-hidden="true">&times;</span>
                        </button>
                    </div>
                    <div className="modal-body">
                        <form>
                            <div className="form-group">
                                <label>Name</label>
                                <input
                                type="text"
                                className="form-control"
                                value={name}
                                onChange={(e) => setName(e.target.value)}
                                />
                            </div>
                            <div className="form-group">
                                <label>Age</label>
                                <input
                                type="number"
                                className="form-control"
                                value={age}
                                onChange={(e) => setAge(e.target.value)}
                                />
                            </div>
                            <div className="form-group">
                                <label>Position</label>
                                <input
                                type="text"
                                className="form-control"
                                value={position}
                                onChange={(e) => setPosition(e.target.value)}
                                />
                            </div>
                        </form>
                    </div>
                    <div className="modal-footer">
                        <button
                            type="button"
                            className="btn btn-primary"
                            onClick={handleFormSubmit}>
                            Add Employee
                        </button>
                    </div>
                    </div>
                </div>
                </div>
            )}
        </div>
    );
}

export default GetEmployeeDetail