CREATE TABLE payees (
  id varchar(100) PRIMARY KEY DEFAULT gen_random_uuid(),
  email varchar(255) NOT NULL,
  first_name varchar(255) NOT NULL,
  last_name varchar(255) NOT NULL,
  address1 varchar(255) NOT NULL,
  address2 varchar(255) DEFAULT NULL,
  city varchar(255) NOT NULL,
  state varchar(255) NOT NULL,
  postal_code varchar(255) NOT NULL,
  country_code varchar(255) NOT NULL,
  phone_number varchar(255) NOT NULL,
  date_of_birth date DEFAULT NULL, -- using caution .. month date only
  national_id varchar(255) DEFAULT NULL, -- using caution .. last 4
  bank_name varchar(255) DEFAULT NULL,
  routing_number varchar(255) DEFAULT NULL,
  account_number varchar(255) DEFAULT NULL, -- using caution .. last 4
  iban varchar(255) DEFAULT NULL, -- using caution .. last 4
  velo_id varchar(255) DEFAULT NULL,
  velo_creation_id varchar(255) DEFAULT NULL,
  velo_invite_status varchar(255) DEFAULT NULL,
  is_active BOOLEAN NOT NULL DEFAULT TRUE,
  created_at timestamp without time zone default (now() at time zone 'utc'),
  updated_at timestamp without time zone default (now() at time zone 'utc')
);