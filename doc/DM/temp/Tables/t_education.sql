-- 
-- TABLE: t_education 
--

CREATE TABLE t_education(
    education_id      BIGINT          AUTO_INCREMENT,
    education_name    VARCHAR(100),
    seq               INT,
    status            INT             DEFAULT 0,
    created_at        DATETIME,
    created_by        VARCHAR(40),
    updated_at        DATETIME,
    updated_by        VARCHAR(40),
    deleted_at        DATETIME,
    deleted_by        VARCHAR(40),
    PRIMARY KEY (education_id)
)ENGINE=INNODB
;



