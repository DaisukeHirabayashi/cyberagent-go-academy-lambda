CREATE TABLE `city_outpatients` (
    `id` INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
    `pref_name` VARCHAR(255),
    `city_name` VARCHAR(255),
    `number` INT,
    `submit_date` TIMESTAMP NOT NULL,
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    index index_city_outpatients_on_pref_name_and_city_name_and_su_date (pref_name, city_name, submit_date),
    index index_city_outpatients_on_city_name_and_submit_date (city_name, submit_date)
);