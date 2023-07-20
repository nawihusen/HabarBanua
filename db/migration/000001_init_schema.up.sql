CREATE TABLE `saksi_candidate` (
    `id`                    BIGINT          NOT NULL AUTO_INCREMENT,
    `member_no`             VARCHAR(16)     NOT NULL,
    `phone_number`          VARCHAR(20)     NOT NULL,
    `qr_code`               VARCHAR(255)    NOT NULL,
    `dtm_crt`               TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `dtm_upd`               TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
);

CREATE TABLE `tps` (
    `id`                    BIGINT          NOT NULL AUTO_INCREMENT,
    `prov_code`             BIGINT          NOT NULL,
    `kab_code`              BIGINT          NOT NULL,
    `kec_code`              BIGINT          NOT NULL,
    `kel_code`              BIGINT          NOT NULL,
    `tps_no`                BIGINT          NOT NULL,
    `dtm_crt`               TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `dtm_upd`               TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
);

CREATE TABLE `saksi_dprt` (
    `id`                    BIGINT          NOT NULL AUTO_INCREMENT,
    `member_no`             VARCHAR(16)     NOT NULL,
    `phone_number`          VARCHAR(20)     NOT NULL,
    `role`                  VARCHAR(6)      NOT NULL,
    `prov_code`             INT             NOT NULL,
    `kab_code`              INT             NOT NULL,
    `kec_code`              INT             NOT NULL,
    `kel_code`              INT             NOT NULL,
    `tps_id`                BIGINT,
    `dtm_crt`               TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `dtm_upd`               TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    CONSTRAINT `fk_saksi_dprt_tps` FOREIGN KEY (`tps_id`) REFERENCES `tps` (`id`)
);


