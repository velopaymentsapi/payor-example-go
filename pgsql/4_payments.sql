CREATE TABLE payments (
  id varchar(100) PRIMARY KEY DEFAULT gen_random_uuid(),
  batch_id varchar(100) DEFAULT gen_random_uuid(),
  payee_id varchar(100) DEFAULT gen_random_uuid(),
  memo varchar(255) NULL,
  amount INTEGER NOT NULL,
  currency varchar(10) NOT NULL,
  velo_source_account varchar(255) NULL,
  velo_payment_id varchar(255) NULL,
  velo_status varchar(255) NULL,
  velo_reason varchar(255) NULL,
  created_at timestamp without time zone default (now() at time zone 'utc'),
  updated_at timestamp without time zone default (now() at time zone 'utc'),
  FOREIGN KEY (batch_id) REFERENCES batches(id),
  FOREIGN KEY (payee_id) REFERENCES payees(id)
);

CREATE INDEX batch_ind ON payments(batch_id);
CREATE INDEX payee_ind ON payments(payee_id);