-- Create user table
CREATE TABLE IF NOT EXISTS `users` (
    `id` VARCHAR(36) NOT NULL COMMENT 'User GUID',
    `account` VARCHAR(32) NOT NULL COMMENT 'User account',
    `password` VARCHAR(60) NOT NULL COMMENT 'User password (bcrypt)',
    `nickname` VARCHAR(8) NOT NULL COMMENT 'User nickname',
    `email` VARCHAR(32) NOT NULL COMMENT 'User email',
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Datetime of creation',
    `created_by` VARCHAR(36) NOT NULL DEFAULT 'system' COMMENT 'Record the creator',
    `updated_at` DATETIME NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'Datetime of modification',
    `updated_by` VARCHAR(36) NULL COMMENT 'Record the updater',
    `deleted_at` DATETIME NULL COMMENT 'Datetime of deletion',
    `deleted_by` VARCHAR(36) NULL COMMENT 'Record the deletor',
    `rec_status` TINYINT(1) NOT NULL DEFAULT 1 COMMENT 'Soft deletion status',
    PRIMARY KEY (`id`),
    UNIQUE `uk-account` (`account`),
    UNIQUE `uk-email` (`email`)
) COMMENT='Users info';

-- Insert default account (e.g. admin)
INSERT IGNORE INTO `users` (id, account, password, nickname, email)
    VALUES ('88428498-035c-11ec-bce8-00ff20c8a33f', 'test', '$2a$08$IYjGRLt6.n6p7Jjw4Z8e/.haqAr0iJEM2ZZ/kjTGCvDrzjHmmc422', 'Tester', 'test@gmail.com');

-- Create chatroom table
CREATE TABLE IF NOT EXISTS `chatrooms` (
    `id` VARCHAR(36) NOT NULL COMMENT 'Room GUID',
    `name` VARCHAR(32) NOT NULL COMMENT 'Room name',
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Datetime of creation',
    `created_by` VARCHAR(36) NOT NULL DEFAULT 'system' COMMENT 'Record the creator',
    `updated_at` DATETIME NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'Datetime of modification',
    `updated_by` VARCHAR(36) NULL COMMENT 'Record the updater',
    `deleted_at` DATETIME NULL COMMENT 'Datetime of deletion',
    `deleted_by` VARCHAR(36) NULL COMMENT 'Record the deletor',
    `rec_status` TINYINT(1) NOT NULL DEFAULT 1 COMMENT 'Soft deletion status',
    PRIMARY KEY (`id`)
) COMMENT='Chatroom info';

-- Create participants table
CREATE TABLE IF NOT EXISTS `participants` (
    `id` VARCHAR(36) NOT NULL COMMENT 'Participant GUID',
    `user_id` VARCHAR(36) NOT NULL COMMENT 'Participant user ID',
    `room_id` VARCHAR(36) NOT NULL COMMENT 'Participant room ID',
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Datetime of creation',
    `created_by` VARCHAR(36) NOT NULL DEFAULT 'system' COMMENT 'Record the creator',
    `updated_at` DATETIME NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'Datetime of modification',
    `updated_by` VARCHAR(36) NULL COMMENT 'Record the updater',
    `deleted_at` DATETIME NULL COMMENT 'Datetime of deletion',
    `deleted_by` VARCHAR(36) NULL COMMENT 'Record the deletor',
    `rec_status` TINYINT(1) NOT NULL DEFAULT 1 COMMENT 'Soft deletion status',
    PRIMARY KEY (`id`),
    UNIQUE `uk-room-users` (`room_id`, `user_id`),
    FOREIGN KEY `fk-userID` (`user_id`) REFERENCES `users`(`id`),
    FOREIGN KEY `fk-roomID` (`room_id`) REFERENCES `chatrooms`(`id`)
) COMMENT='Participants info';

-- Create messages table
CREATE TABLE IF NOT EXISTS `messages` (
    `id` VARCHAR(36) NOT NULL COMMENT 'messages GUID',
    `content` TEXT NOT NULL COMMENT 'messages text',
    `user_id` VARCHAR(36) NOT NULL COMMENT 'Participant user ID',
    `room_id` VARCHAR(36) NOT NULL COMMENT 'Participant room ID',
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Datetime of creation',
    `created_by` VARCHAR(36) NOT NULL DEFAULT 'system' COMMENT 'Record the creator',
    `updated_at` DATETIME NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'Datetime of modification',
    `updated_by` VARCHAR(36) NULL COMMENT 'Record the updater',
    `deleted_at` DATETIME NULL COMMENT 'Datetime of deletion',
    `deleted_by` VARCHAR(36) NULL COMMENT 'Record the deletor',
    `rec_status` TINYINT(1) NOT NULL DEFAULT 1 COMMENT 'Soft deletion status',
    PRIMARY KEY (`id`),
    FOREIGN KEY `fk-userID` (`user_id`) REFERENCES `users`(`id`),
    FOREIGN KEY `fk-roomID` (`room_id`) REFERENCES `chatrooms`(`id`)
) COMMENT='Messages';