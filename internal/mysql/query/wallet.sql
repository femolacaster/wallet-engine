-- name: GenerateWallet :execresult
INSERT INTO wallets (
  id,
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
?, ?,?, ?, ?, ?, ?, ?, ?, ?
);


-- name: InsertTransaction :execresult
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
  ?,?,?,?,?,?,?,?,?,?
);


-- name: ChangeWalletStatus :exec
UPDATE wallets SET
  is_active = ?
WHERE id = ?;
