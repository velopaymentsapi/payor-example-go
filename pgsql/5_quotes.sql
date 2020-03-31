CREATE TABLE quotes (
  id varchar(100) PRIMARY KEY DEFAULT gen_random_uuid(),
  batch_id varchar(255) NULL,
  velo_id varchar(255) NULL,
  velo_funding_status varchar(255) NOT NULL,
  velo_status varchar(255) NOT NULL,
  created_at timestamp without time zone default (now() at time zone 'utc'),
  updated_at timestamp without time zone default (now() at time zone 'utc'),
  FOREIGN KEY (batch_id) REFERENCES batches(id)
);
CREATE INDEX quote_batch_ind ON quotes(batch_id);