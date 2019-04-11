-- migrate:up
CREATE EXTENSION pgcrypto;

CREATE TYPE payment_scheme_type_enum AS ENUM (
	'FPS',
    'Bacs',
    'SEPAINSTANT'
);

CREATE TYPE resource_enum AS ENUM (
	'Payment'
);

CREATE TABLE scheme_payment_type
(
    scheme_payment_type_id SERIAL,
    value VARCHAR(100) NOT NULL,

    payment_scheme payment_scheme_type_enum NOT NULL,
    resource resource_enum NOT NULL,

    created_at int8 NOT NULL DEFAULT extract(epoch from now()),
    updated_at int8 NOT NULL DEFAULT extract(epoch from now()),

    PRIMARY KEY(scheme_payment_type_id),
    UNIQUE (value)
);

create table currency
(
    currency_id uuid NOT NULL DEFAULT gen_random_uuid(),
    name VARCHAR(50) NOT NULL,
    abbrv VARCHAR(20) NOT NULL, 

    created_at int8 NOT NULL DEFAULT extract(epoch from now()),
    updated_at int8 NOT NULL DEFAULT extract(epoch from now()),

    PRIMARY KEY(currency_id),
    UNIQUE (name, abbrv)
);

create table organisation
(
  organisation_id uuid NOT NULL DEFAULT gen_random_uuid(),
  
  created_at int8 NOT NULL DEFAULT extract(epoch from now()),
  updated_at int8 NOT NULL DEFAULT extract(epoch from now()),

  PRIMARY KEY(organisation_id)
);

create table customer
(
    customer_id uuid NOT NULL DEFAULT gen_random_uuid(),
    customer_name VARCHAR(100),
    customer_address VARCHAR(255),

    created_at int8 NOT NULL DEFAULT extract(epoch from now()),
    updated_at int8 NOT NULL DEFAULT extract(epoch from now()),

    PRIMARY KEY(customer_id)
);

create table bank
(
    bank_id integer NOT NULL,
    bank_code VARCHAR(20),
    bank_name VARCHAR(100),
    bank_address VARCHAR(255),

    created_at int8 NOT NULL DEFAULT extract(epoch from now()),
    updated_at int8 NOT NULL DEFAULT extract(epoch from now()),

    PRIMARY KEY(bank_id)
);

create table account
(
    account_id uuid NOT NULL DEFAULT gen_random_uuid(),
    bank_id integer REFERENCES bank(bank_id),
    account_number bigint NOT NULL,
    account_number_code VARCHAR(20),
    account_type int8 NOT NULL DEFAULT 0,

    account_name VARCHAR(50),

    created_at int8 NOT NULL DEFAULT extract(epoch from now()),
    updated_at int8 NOT NULL DEFAULT extract(epoch from now()),

    PRIMARY KEY(account_id),
    UNIQUE (bank_id, account_number)
);

create table customer_account
(
    customer_account_id uuid NOT NULL DEFAULT gen_random_uuid(),
    
    customer_id uuid REFERENCES customer(customer_id),
    account_id uuid REFERENCES account(account_id),
    
    created_at int8 NOT NULL DEFAULT extract(epoch from now()),
    updated_at int8 NOT NULL DEFAULT extract(epoch from now()),

    PRIMARY KEY(customer_account_id),
    UNIQUE (customer_id, account_id)
);


create table payment
(
    payment_id uuid NOT NULL DEFAULT gen_random_uuid(),
    version int8 NOT NULL DEFAULT 0,
    organisation_id uuid NOT NULL REFERENCES organisation(organisation_id),
    type resource_enum NOT NULL DEFAULT 'Payment',

    currency_id uuid NOT NULL REFERENCES currency(currency_id) ON DELETE CASCADE,

    amount decimal NOT NULL,

    beneficary_id uuid NOT NULL REFERENCES customer_account(customer_account_id) ON DELETE CASCADE,
    debtor_id uuid NOT NULL REFERENCES customer_account(customer_account_id) ON DELETE CASCADE,

    end_to_end_reference VARCHAR(255),
    numeric_reference bigint NOT NULL DEFAULT 0,
    payment_purpose VARCHAR(255),

    scheme_payment_type_id integer NOT NULL REFERENCES scheme_payment_type(scheme_payment_type_id) ON DELETE CASCADE,

    reference VARCHAR(255),
    processing_date date,

    created_at int8 NOT NULL DEFAULT extract(epoch from now()),
    updated_at int8 NOT NULL DEFAULT extract(epoch from now()),

    PRIMARY KEY(payment_id)
);

-- migrate:down
-- DROP TABLE attribute;
DROP TABLE payment;
DROP TABLE currency;
DROP TABLE organisation;
DROP TABLE customer_account;
DROP TABLE scheme_payment_type;
DROP TABLE customer;
DROP TABLE account;
DROP TABLE bank;
DROP TYPE resource_enum;
DROP TYPE payment_scheme_type_enum;

DROP EXTENSION pgcrypto;