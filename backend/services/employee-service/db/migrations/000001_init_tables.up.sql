CREATE TABLE employees (
    id VARCHAR(255) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    age INT NOT NULL,
    position VARCHAR(255) NOT NULL
);

CREATE TABLE statistics (
    id VARCHAR(255) PRIMARY KEY,
    employee_id VARCHAR(255) NOT NULL,
    view_count BIGINT NOT NULL,
    FOREIGN KEY (employee_id) REFERENCES employees(id)
);
