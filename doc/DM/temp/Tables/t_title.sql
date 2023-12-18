-- 
-- TABLE: t_title 
--

CREATE TABLE t_title(
    title_id      BIGINT          AUTO_INCREMENT,
    title_name    VARCHAR(100),
    seq           INT,
    status        INT             DEFAULT 0,
    created_at    DATETIME,
    created_by    VARCHAR(40),
    updated_at    DATETIME,
    updated_by    VARCHAR(40),
    deleted_at    DATETIME,
    deleted_by    VARCHAR(40),
    PRIMARY KEY (title_id)
)ENGINE=INNODB
;



