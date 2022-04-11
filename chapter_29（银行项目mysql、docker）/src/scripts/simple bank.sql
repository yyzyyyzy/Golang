CREATE TABLE `account` (
  `id` bigint PRIMARY KEY AUTO_INCREMENT,
  `owenr` varchar(255) NOT NULL,
  `balance` bigint NOT NULL,
  `currency` varchar(255) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE `entries` (
  `id` bigint PRIMARY KEY AUTO_INCREMENT,
  `account_id` bigint,
  `amount` bigint NOT NULL COMMENT 'can be negative or postive',
  `created_at` timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE `transfers` (
  `id` bigint PRIMARY KEY AUTO_INCREMENT,
  `from_account_id` bigint,
  `to_account_id` bigint,
  `amount` bigint NOT NULL COMMENT 'must be postive',
  `created_at` timestamp NOT NULL DEFAULT (now())
);

ALTER TABLE `entries` ADD FOREIGN KEY (`account_id`) REFERENCES `account` (`id`);

ALTER TABLE `transfers` ADD FOREIGN KEY (`from_account_id`) REFERENCES `account` (`id`);

ALTER TABLE `transfers` ADD FOREIGN KEY (`to_account_id`) REFERENCES `account` (`id`);

CREATE INDEX `account_index_0` ON `account` (`owenr`);

CREATE INDEX `entries_index_1` ON `entries` (`account_id`);

CREATE INDEX `transfers_index_2` ON `transfers` (`from_account_id`);

CREATE INDEX `transfers_index_3` ON `transfers` (`to_account_id`);

CREATE INDEX `transfers_index_4` ON `transfers` (`from_account_id`, `to_account_id`);
