-- 
-- TABLE: t_nurse 
--

CREATE TABLE t_nurse(
    nurse_id        BIGINT                   AUTO_INCREMENT,
    nurse_code      VARCHAR(40),
    nurse_name      VARCHAR(100),
    start_date      DATETIME,
    end_date        DATETIME,
    sex             INT,
    birthday        DATETIME,
    area_id         INT,
    dept_id         BIGINT,
    position_id     BIGINT,
    education_id    BIGINT,
    title_id        BIGINT,
    staff_type      INT,
    id_no           VARCHAR(40),
    cert_no         VARCHAR(40),
    cert_enddate    DATETIME,
    reg_date        DATETIME,
    note            NATIONAL VARCHAR(500),
    session_id      VARCHAR(40),
    status          INT                      DEFAULT 0,
    created_at      DATETIME,
    created_by      VARCHAR(40),
    updated_at      DATETIME,
    updated_by      VARCHAR(40),
    deleted_at      DATETIME,
    deleted_by      VARCHAR(40),
    PRIMARY KEY (nurse_id)
)ENGINE=INNODB
;



-- 
-- INDEX: idx_t_nursecode 
--

CREATE INDEX idx_t_nursecode ON t_nurse(nurse_code)
;
-- 
-- INDEX: idx_t_nursesessionid 
--

CREATE INDEX idx_t_nursesessionid ON t_nurse(session_id)
;
