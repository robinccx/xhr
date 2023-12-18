-- 
-- TABLE: t_nurse_import 
--

CREATE TABLE t_nurse_import(
    id                 BIGINT                   AUTO_INCREMENT,
    session_id         VARCHAR(40),
    nurse_id           BIGINT,
    nurse_code         VARCHAR(40),
    nurse_name         VARCHAR(100),
    start_date         VARCHAR(40),
    end_date           VARCHAR(40),
    sex                VARCHAR(10),
    birthday           VARCHAR(40),
    area_id            INT,
    area_code          VARCHAR(40),
    dept_id            BIGINT,
    dept_code          VARCHAR(100),
    education_id       BIGINT,
    education_code     VARCHAR(40),
    title_id           BIGINT,
    title_code         VARCHAR(40),
    staff_type         INT,
    staff_type_code    VARCHAR(40),
    id_no              VARCHAR(40),
    cert_no            VARCHAR(40),
    cert_enddate       DATETIME,
    reg_date           DATETIME,
    note               NATIONAL VARCHAR(500),
    seq                INT,
    status             INT                      DEFAULT 0,
    created_at         DATETIME,
    created_by         VARCHAR(40),
    PRIMARY KEY (id)
)ENGINE=INNODB
;



-- 
-- INDEX: idx_t_nurse_importcode 
--

CREATE INDEX idx_t_nurse_importcode ON t_nurse_import(nurse_code)
;
-- 
-- INDEX: idx_t_nurse_importsessionid 
--

CREATE INDEX idx_t_nurse_importsessionid ON t_nurse_import(session_id)
;
