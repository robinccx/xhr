-- 
-- TABLE: t_position 
--

CREATE TABLE t_position(
    position_id      BIGINT          AUTO_INCREMENT,
    position_code    VARCHAR(40),
    position_name    VARCHAR(100),
    seq              INT,
    status           INT             DEFAULT 0,
    created_at       DATETIME,
    created_by       VARCHAR(40),
    updated_at       DATETIME,
    updated_by       VARCHAR(40),
    deleted_at       DATETIME,
    deleted_by       VARCHAR(40),
    PRIMARY KEY (position_id)
)ENGINE=INNODB
;



