-- name: GenerateWallet :one
INSERT INTO wallet (
  wallet_number,
  is_active,
  first_name,
  last_name,
  email,
  secretkey,
  bvn,
  dob,
  currency
)
VALUES (
  @wallet_number,
  @is_active,
  @first_name,
  @last_name,
  @email,
  @secretkey,
  @bvn,
  @dob,
  @currency
)
RETURNING id;

-- name: InsertTransaction :one
INSERT INTO transactions (
  id,
  transaction_ref,
  transaction_type,
  transaction_timestamp,
  amount,
  secretkey,
  transaction_status,
  transaction_description,
  balance,
  wallet_id
)
VALUES (
  @id,
  @transaction_ref,
  @transaction_type,
  @transaction_timestamp,
  @amount,
  @secretkey,
  @transaction_status,
  @transaction_description,
  @balance,
  @wallet_id
)
RETURNING id;

-- name: UpdateWallet :exec
UPDATE wallet SET
  wallet_number = @wallet_number,
  is_active = @is_active,
  first_name = @first_name,
  last_name = @last_name,
  email = @email,
  secretkey = @secretkey,
  bvn= @bvn,
  dob = @dob,
  currency = @currency,
  modified_time = @modified_time
WHERE id = @id;
