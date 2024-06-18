-- ###################### INSERT_VALUES ######################

USE `transaction-processor`;

--
-- Dumping data for table `transactions`
INSERT INTO `transaction-processor`.`account` VALUES(
    '17d340fa-5bf5-4429-8167-bafe4c0af0a7', 'Checking', 'CASH', 'USD', '2024-01-07 07:17:17', '2024-01-07 07:17:17'
);

--
-- Dumping data for table `account`
INSERT INTO `transaction-processor`.`transactions` VALUES (
  'aa5551bb-a051-4e82-ac5b-7e59c43c8867', '17d340fa-5bf5-4429-8167-bafe4c0af0a7', 227.32, 'CREDIT', '2024-01-07 07:17:27'
);

-- ###################### INSERT_VALUES ######################