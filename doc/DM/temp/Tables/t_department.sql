-- 
-- TABLE: t_department 
--

CREATE TABLE t_department(
    dept_id       BIGINT          AUTO_INCREMENT,
    parent_id     BIGINT,
    dept_code     VARCHAR(40),
    dept_name     VARCHAR(100),
    dept_type     VARCHAR(40),
    att1          VARCHAR(40),
    att2          VARCHAR(40),
    att3          VARCHAR(100),
    leader        VARCHAR(40),
    phone         VARCHAR(20),
    email         VARCHAR(60),
    status        INT             DEFAULT 0,
    created_at    DATETIME,
    created_by    VARCHAR(40),
    updated_at    DATETIME,
    updated_by    VARCHAR(40),
    deleted_at    DATETIME,
    deleted_by    VARCHAR(40),
    PRIMARY KEY (dept_id)
)ENGINE=INNODB
;



-- 
-- INDEX: idx_t_departmentcode 
--

CREATE INDEX idx_t_departmentcode ON t_department(dept_code)
;
-- 
-- INDEX: idx_t_departmentpid 
--

CREATE INDEX idx_t_departmentpid ON t_department(parent_id)
;
