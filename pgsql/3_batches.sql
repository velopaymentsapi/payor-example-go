CREATE TABLE batches (
  id varchar(100) PRIMARY KEY DEFAULT gen_random_uuid(),
  name varchar(255) NOT NULL,
  velo_payout_id varchar(255) NULL,
  velo_quote_id varchar(255) NULL,
  velo_payments_submitted INTEGER NOT NULL DEFAULT 0,
  velo_payments_rejected INTEGER NOT NULL DEFAULT 0,
  velo_payments_accepted INTEGER NOT NULL DEFAULT 0,
  velo_payments_incomplete INTEGER NOT NULL DEFAULT 0,
  velo_payments_released INTEGER NOT NULL DEFAULT 0,
  velo_payments_confirmed INTEGER NOT NULL DEFAULT 0,
  velo_payments_returned INTEGER NOT NULL DEFAULT 0,
  velo_status varchar(255) NOT NULL,
  created_at timestamp without time zone default (now() at time zone 'utc'),
  updated_at timestamp without time zone default (now() at time zone 'utc')
);