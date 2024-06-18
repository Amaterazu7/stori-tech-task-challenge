-- ###################### INIT_DATABASE ######################

CREATE DATABASE IF NOT EXISTS `transaction-processor`;

USE `transaction-processor`;
GRANT ALL PRIVILEGES ON *.* TO 'goferProcessor'@'%';
FLUSH PRIVILEGES;

SHOW GRANTS FOR 'goferProcessor'@'%';

-- ###################### INIT_DATABASE ######################
