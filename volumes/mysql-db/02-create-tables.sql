-- ###################### CREATE_TABLES ######################

USE `transaction-processor`;

CREATE TABLE account (
    id         varchar(50)                              not null,
    name       varchar(50)                              not null,
    asset      enum ('CASH', 'CRYPTO')                  not null,
    Type       enum ('USD', 'GBP', 'EUR', 'ETH', 'BTC') not null,
    updated_at datetime                                 not null,
    created_at datetime                                 not null,

    CONSTRAINT account_pk PRIMARY KEY (id)
);

CREATE TABLE transactions (
    id         varchar(50)              not null,
    account_id varchar(50)              not null,
    amount     double                   not null,
    tx_type    enum ('DEBIT', 'CREDIT') not null,
    created_at datetime                 not null,

    CONSTRAINT transactions_pk PRIMARY KEY (id)
);

ALTER TABLE transactions ADD CONSTRAINT transactions_account_id_fk FOREIGN KEY (account_id) REFERENCES account (id);

-- ###################### CREATE_TABLES ######################