-- 
-- TABLE: t_nurse_change_log 
--

CREATE TABLE t_nurse_change_log(
    id                  BIGINT          AUTO_INCREMENT,
    nurse_id            BIGINT,
    change_type         INT,
    from_area_id        INT,
    from_dept_id        BIGINT,
    from_entity_id      BIGINT,
    from_entity_name    VARCHAR(100),
    to_area_id          INT,
    to_dept_id          BIGINT,
    to_entity_id        BIGINT,
    to_entity_name      VARCHAR(100),
    trans_date          DATETIME,
    status              INT             DEFAULT 0,
    created_at          DATETIME,
    created_by          VARCHAR(40),
    updated_at          DATETIME,
    updated_by          VARCHAR(40),
    deleted_at          DATETIME,
    deleted_by          VARCHAR(40),
    PRIMARY KEY (id)
)ENGINE=INNODB
;



-- 
-- INDEX: idx_t_nurse_change_logtransdate 
--

CREATE INDEX idx_t_nurse_change_logtransdate ON t_nurse_change_log(trans_date)
;
-- 
-- INDEX: idx_t_nurse_change_lognurseid 
--

CREATE INDEX idx_t_nurse_change_lognurseid ON t_nurse_change_log(nurse_id)
;
-- 
-- INDEX: idx_t_nurse_change_logchangetype 
--

CREATE INDEX idx_t_nurse_change_logchangetype ON t_nurse_change_log(change_type)
;
