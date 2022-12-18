CREATE TABLE `outpatient_histories` (
    `id` INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
    `facility_id` INT,
    `facility_name` VARCHAR(255),
    `zip_cod` VARCHAR(255),
    `pref_name` VARCHAR(255),
    `facility_addr` VARCHAR(255),
    `facility_tel` VARCHAR(255),
    `facility_type` VARCHAR(255),
    `ans_type` VARCHAR(255),
    `local_gov_code` VARCHAR(255),
    `city_name` VARCHAR(255),
    `facility_code` VARCHAR(255),
    `submit_date` TIMESTAMP NOT NULL,
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);