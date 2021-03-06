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
  currency
)
VALUES (
?, ?,?, ?, ?, ?, ?, ?, ?
);


-- name: InsertTransaction :execresult
INSERT INTO transactions (
  id,
  transaction_ref,
  transaction_type,
  amount,
  secretkey,
  transaction_status,
  transaction_description,
  wallet_id
)
VALUES (
  ?,?,?,?,?,?,?,?
);


-- name: ChangeWalletStatus :exec
UPDATE wallets SET
  is_active = ?
WHERE id = ?;
