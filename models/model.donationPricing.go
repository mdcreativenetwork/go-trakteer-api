package models

import "time"

type DonationPricing struct {
	Id             uint      `db:"id"`
	UserId         uint      `db:"user_id"`
	DonationTypeId uint      `db:"donation_type_id"`
	Amount         uint64    `db:"amount"`
	CreatedAt      time.Time `db:"created_at"`
	UpdatedAt      time.Time `db:"updated_at"`
	User           Users
	DonationTypes  DonationTypes
}

// CREATE TABLE IF NOT EXISTS "donation_pricing" (
//   "id" serial PRIMARY KEY,
//   "user_id" int NOT NULL UNSIGNED UNIQUE,
//   "donation_type_id" int NOT NULL UNSIGNED,
//   "amount" bigint NOT NULL UNSIGNED,
//   "created_at" NULL DEFAULT (now()),
//   "updated_at" NULL DEFAULT (now())
// );

// ALTER TABLE "donation_pricing" ADD FOREIGN KEY user_id REFERENCES "users" ON DELETE RESCRICT;
// ALTER TABLE "donation_pricing" ADD FOREIGN KEY donation_type_id REFERENCES "donation_type" ON DELETE RESCRICT;
