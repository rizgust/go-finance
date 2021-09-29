CREATE TABLE "m_chart_of_accounts" (
   "id" SERIAL NOT NULL,
   "owner_id" INT NOT NULL,
   "code" VARCHAR(20) NOT NULL,
   "name" varchar(50) NOT NULL,
   "description" varchar(100),
   "category" SMALLINT NOT NULL,
   "level" SMALLINT NOT NULL,
   "created_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
   "created_by" INT NOT NULL,
   "updated_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
   "updated_by" INT NOT NULL,
   "deleted_at" TIMESTAMPTZ,
   "deleted_by" INT,
   PRIMARY KEY ("id"),
	UNIQUE (owner_id, code, name)
);

CREATE TABLE "m_banks" (
   "id" SERIAL NOT NULL,
   "owner_id" INT NOT NULL,
   "code" varchar(50) NOT NULL,
   "name" varchar(100) NOT NULL,
   "alias" varchar(50),
   "created_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
   "created_by" INT NOT NULL,
   "updated_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
   "updated_by" INT NOT NULL,
   PRIMARY KEY ("id"),
	UNIQUE (owner_id, code, name)
);

CREATE TABLE "m_bank_accounts" (
   "id" SERIAL NOT NULL,
   "owner_id" INT NOT NULL,
   "bank_id" INT NOT NULL,
   "branch_name" varchar(100) NOT NULL,
   "name" varchar(100) NOT NULL,
   "number" varchar(50) NOT NULL,
   "created_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
   "created_by" INT NOT NULL,
   "updated_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
   "updated_by" INT NOT NULL,
   PRIMARY KEY ("id"),
	UNIQUE (owner_id, bank_id, branch_name)
);

CREATE TABLE "m_virtual_accounts" (
   "id" SERIAL NOT NULL,
   "owner_id" INT NOT NULL,
   "bank_id" INT NOT NULL,
   "number" varchar(50) NOT NULL,
   "created_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
   "created_by" INT NOT NULL,
   "updated_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
   "updated_by" INT NOT NULL,
   PRIMARY KEY ("id"),
	UNIQUE (owner_id, bank_id, number)
);

CREATE TABLE "m_discounts" (
   "id" SERIAL NOT NULL,
   "owner_id" INT NOT NULL,
   "code" VARCHAR(20),
   "name" VARCHAR(50) NOT NULL,
   "description" VARCHAR(100),
   "is_active" BOOLEAN NOT NULL,
   "value" MONEY NOT NULL,
   "is_percent" BOOLEAN NOT NULL,
   "start_date" DATE,
   "end_date" DATE,
   "created_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
   "created_by" INT NOT NULL,
   "updated_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
   "updated_by" INT NOT NULL,
   "deleted_at" TIMESTAMPTZ,
   "deleted_by" INT,
   PRIMARY KEY ("id"));

CREATE TABLE "m_book_periods" (
   "id" SERIAL NOT NULL,
   "owner_id" INT NOT NULL,
   "code" VARCHAR(20),
   "name" varchar(50) NOT NULL,
   "description" varchar(100),
   "status" SMALLINT,
   "start_date" DATE NOT NULL,
   "end_date" DATE NOT NULL,
   "created_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
   "created_by" INT NOT NULL,
   "updated_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
   "updated_by" INT NOT NULL,
   "deleted_at" TIMESTAMPTZ,
   "deleted_by" INT,
   PRIMARY KEY ("id"),
	UNIQUE (owner_id, name, start_date, end_date)
);

CREATE TABLE "account_payables" (
   "id" SERIAL NOT NULL,
   "owner_id" INT NOT NULL,
   "code" VARCHAR(20) NOT NULL,
   "name" varchar(50) NOT NULL,
   "description" varchar(100),
   "src_account_id" INT NOT NULL,
   "dst_account_id" INT NOT NULL,
   "recurring_type" SMALLINT,
   "recurring_periode" SMALLINT,
   "created_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
   "created_by" INT NOT NULL,
   "updated_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
   "updated_by" INT NOT NULL,
   "deleted_at" TIMESTAMPTZ,
   "deleted_by" INT,
   PRIMARY KEY ("id"),
	UNIQUE (owner_id, code, src_account_id, dst_account_id)
);

CREATE TABLE "account_receiveables" (
   "id" SERIAL NOT NULL,
   "owner_id" INT NOT NULL,
   "code" VARCHAR(20) NOT NULL,
   "name" varchar(50) NOT NULL,
   "description" varchar(100),
   "src_account_id" INT NOT NULL,
   "dst_account_id" INT NOT NULL,
   "recurring_type" SMALLINT,
   "recurring_periode" SMALLINT,
   "user_type" SMALLINT NOT NULL,
   "created_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
   "created_by" INT NOT NULL,
   "updated_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
   "updated_by" INT NOT NULL,
   "deleted_at" TIMESTAMPTZ,
   "deleted_by" INT,
   PRIMARY KEY ("id"),
	UNIQUE (owner_id, code, src_account_id, dst_account_id)
);

CREATE TABLE "invoices" (
   "id" SERIAL NOT NULL,
   "owner_id" INT NOT NULL,
   "user_id" INT NOT NULL,
   "number" VARCHAR(20) NOT NULL,
   "ar_id" INT NOT NULL,
   "status" SMALLINT,
   "amount" MONEY NOT NULL,
   "amount_paid" MONEY NOT NULL,
   "date" DATE NOT NULL,
   "due_date" DATE NOT NULL,
   "additional_info" JSON,
   "payment_id" INT,
   "discount_id" INT,
   "created_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
   "created_by" INT NOT NULL,
   "updated_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
   "updated_by" INT NOT NULL,
   "deleted_at" TIMESTAMPTZ,
   "deleted_by" INT,
   "period_id" INT,
   "class_program_id" INT,
   "class_level_id" INT,
   "class_specialization_id" INT,
   "male" BOOLEAN,
   "recurring_type" SMALLINT,
   "recurring_period" SMALLINT,
   "installment" SMALLINT,
   "mutation" BOOLEAN,
   "boarding" BOOLEAN,
   "admission_line_id" SMALLINT,
   "admission_batch_id" SMALLINT,
   PRIMARY KEY ("id"),
	UNIQUE (owner_id, user_id, number)
);

CREATE TABLE "trx_receives" (
   "id" SERIAL NOT NULL,
   "owner_id" INT NOT NULL,
   "user_id" INT NOT NULL,
   "number" VARCHAR(20),
   "amount" MONEY NOT NULL,
   "date" DATE NOT NULL,
   "status" SMALLINT,
   "description" VARCHAR(255),
   "payment_method" SMALLINT NOT NULL,
   "created_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
   "created_by" INT NOT NULL,
   "updated_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
   "updated_by" INT NOT NULL,
   "journal_id" INT,
   PRIMARY KEY ("id")
);

CREATE TABLE "wallets" (
   "id" SERIAL NOT NULL,
   "owner_id" INT NOT NULL,
   "user_id" INT NOT NULL,
   "code" varchar(50) NOT NULL,
   "name" varchar(100) NOT NULL,
   "alias" varchar(50),
   "balance" MONEY NOT NULL,
   "account_id" INT,
   "other_info" JSONB,
   "allow_minus" BOOLEAN,
   "created_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
   "created_by" INT NOT NULL,
   "updated_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
   "updated_by" INT NOT NULL,
   "deleted_at" TIMESTAMPTZ,
   "deleted_by" INT,
   PRIMARY KEY ("id"),
	UNIQUE (owner_id, user_id, code)
);

CREATE TABLE "trx_receive_details" (
   "id" SERIAL NOT NULL,
   "trx_receives_id" INT NOT NULL,
   "amount" MONEY NOT NULL,
   "is_source" BOOLEAN NOT NULL,
   "account_id" INT NOT NULL,
   "model" VARCHAR(50),
   "model_id" INT,
   "created_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
   "created_by" INT NOT NULL,
   "updated_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
   "updated_by" INT NOT NULL,
   PRIMARY KEY ("id")
);

CREATE TABLE "journals" (
   "id" SERIAL NOT NULL,
   "owner_id" INT NOT NULL,
   "number" VARCHAR(50) NOT NULL,
   "date" DATE NOT NULL,
   "general_ledger_id" INT,
   "status" SMALLINT NOT NULL,
   "amount" MONEY NOT NULL,
   "description" VARCHAR(255),
   "reff_model" VARCHAR(50),
   "reff_model_id" INT,
   "created_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
   "created_by" INT NOT NULL,
   "updated_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
   "updated_by" INT NOT NULL,
   "deleted_at" TIMESTAMPTZ,
   "deleted_by" INT,
   PRIMARY KEY ("id")
);

CREATE TABLE "trx_wallets" (
   "id" SERIAL NOT NULL,
   "wallet_id" INT NOT NULL,
   "amount" MONEY NOT NULL,
   "balance_before" MONEY NOT NULL,
   "balance" MONEY NOT NULL,
   "dest_model" VARCHAR(50),
   "dest_model_id" INT,
   "trx_receives_id" INT NOT NULL,
   "created_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
   "created_by" BIGINT NOT NULL,
   "updated_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
   "updated_by" INT NOT NULL,
   "debit" BOOLEAN NOT NULL,
   PRIMARY KEY ("id")
);

CREATE TABLE "trx_bank_transfers" (
   "id" SERIAL NOT NULL,
   "owner_id" INT NOT NULL,
   "user_id" INT NOT NULL,
   "bank_id" INT NOT NULL,
   "number" VARCHAR(20),
   "date" TIMESTAMPTZ NOT NULL,
   "status" SMALLINT NOT NULL,
   "amount" MONEY NOT NULL,
   "description" VARCHAR(100),
   "dest_model" VARCHAR(50),
   "dest_model_id" INT,
   "account_owner" VARCHAR(100),
   "account_number" VARCHAR(50),
   "additional_info" JSON,
   "trx_receives_id" INT NOT NULL,
   "created_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
   "created_by" INT NOT NULL,
   "updated_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
   "updated_by" INT NOT NULL,
   PRIMARY KEY ("id")
);

CREATE TABLE "account_receiveable_rules" (
   "id" SERIAL NOT NULL,
   "owner_id" INT NOT NULL,
   "ar_id" INT NOT NULL,
   "period_id" INT NOT NULL,
   "rule" JSON NOT NULL,
   "created_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
   "created_by" INT NOT NULL,
   "updated_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
   "updated_by" INT NOT NULL,
   "deleted_at" TIMESTAMPTZ,
   "deleted_by" INT,
   PRIMARY KEY ("id"),
	UNIQUE (owner_id, ar_id, period_id)
);

CREATE TABLE "general_ledgers" (
   "id" SERIAL NOT NULL,
   "owner_id" INT NOT NULL,
   "number" VARCHAR(50) NOT NULL,
   "status" SMALLINT NOT NULL,
   "date" DATE NOT NULL,
   "created_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
   "created_by" INT NOT NULL,
   "updated_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
   "updated_by" INT NOT NULL,
   "deleted_at" TIMESTAMPTZ,
   "deleted_by" INT,
   PRIMARY KEY ("id")
);

CREATE TABLE "journal_details" (
   "id" SERIAL NOT NULL,
   "owner_id" INT NOT NULL,
   "account_id" BIGINT NOT NULL,
   "journal_id" INT NOT NULL,
   "amount" MONEY NOT NULL,
   "reff_model" VARCHAR(50),
   "reff_model_id" INT,
   "created_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
   "created_by" INT NOT NULL,
   "updated_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
   "updated_by" INT NOT NULL,
   "deleted_at" TIMESTAMPTZ,
   "deleted_by" INT,
   PRIMARY KEY ("id")
);

CREATE TABLE "account_payable_rules" (
   "id" SERIAL NOT NULL,
   "owner_id" INT NOT NULL,
   "ap_id" INT NOT NULL,
   "period_id" INT NOT NULL,
   "rule" JSONB NOT NULL,
   "created_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
   "created_by" INT NOT NULL,
   "updated_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
   "updated_by" INT NOT NULL,
   "deleted_at" TIMESTAMPTZ,
   "deleted_by" INT,
   PRIMARY KEY ("id"),
	UNIQUE (owner_id, ap_id, period_id)
);

CREATE TABLE "trx_expenses" (
   "id" SERIAL NOT NULL,
   "owner_id" INT NOT NULL,
   "user_id" INT NOT NULL,
   "number" VARCHAR(20),
   "amount" MONEY NOT NULL,
   "date" DATE NOT NULL,
   "status" SMALLINT,
   "description" VARCHAR(255),
   "payment_method" SMALLINT NOT NULL,
   "created_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
   "created_by" INT NOT NULL,
   "updated_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
   "updated_by" INT NOT NULL,
   "journal_id" INT,
   PRIMARY KEY ("id")
);

CREATE TABLE "trx_expense_details" (
   "id" SERIAL NOT NULL,
   "trx_expenses_id" INT NOT NULL,
   "amount" MONEY NOT NULL,
   "is_source" BOOLEAN NOT NULL,
   "account_id" INT NOT NULL,
   "model" VARCHAR(50),
   "model_id" INT,
   "created_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
   "created_by" INT NOT NULL,
   "updated_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
   "updated_by" INT NOT NULL,
   PRIMARY KEY ("id")
);

CREATE TABLE "purchases" (
   "id" SERIAL NOT NULL,
   "owner_id" INT NOT NULL,
   "user_id" INT NOT NULL,
   "number" VARCHAR(20) NOT NULL,
   "ap_id" INT NOT NULL,
   "status" SMALLINT,
   "amount" MONEY NOT NULL,
   "amount_paid" MONEY NOT NULL,
   "date" DATE NOT NULL,
   "due_date" DATE NOT NULL,
   "additional_info" JSON,
   "period_id" INT,
   "expense_id" INT,
   "discount_id" INT,
   "created_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
   "created_by" INT NOT NULL,
   "updated_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
   "updated_by" INT NOT NULL,
   "deleted_at" TIMESTAMPTZ,
   "deleted_by" INT,
   PRIMARY KEY ("id"),
	UNIQUE (owner_id, user_id, number)
);

CREATE TABLE "general_ledger_details" (
   "id" SERIAL NOT NULL,
   "owner_id" INT NOT NULL,
   "general_ledger_id" INT NOT NULL,
   "account_id" INT NOT NULL,
   "debit" MONEY NOT NULL,
   "credit" MONEY NOT NULL,
   "balance_before" MONEY NOT NULL,
   "balance" MONEY NOT NULL,
   "created_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
   "created_by" INT NOT NULL,
   "updated_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
   "updated_by" INT NOT NULL,
   "deleted_at" TIMESTAMPTZ,
   "deleted_by" INT,
   PRIMARY KEY ("id")
);


ALTER TABLE "m_bank_accounts" ADD CONSTRAINT "FK_m_bank_accounts_banks" FOREIGN KEY ("bank_id") REFERENCES "m_banks"("id");

ALTER TABLE "m_virtual_accounts" ADD CONSTRAINT "FK_va_banks" FOREIGN KEY ("bank_id") REFERENCES "m_banks"("id");

ALTER TABLE "account_payables" ADD CONSTRAINT "FK_ap_dst_coa" FOREIGN KEY ("dst_account_id") REFERENCES "m_chart_of_accounts"("id");

ALTER TABLE "account_payables" ADD CONSTRAINT "FK_ap_src_coa" FOREIGN KEY ("src_account_id") REFERENCES "m_chart_of_accounts"("id");

ALTER TABLE "account_receiveables" ADD CONSTRAINT "FK_ar_src_coa" FOREIGN KEY ("src_account_id") REFERENCES "m_chart_of_accounts"("id");

ALTER TABLE "account_receiveables" ADD CONSTRAINT "FK_ar_dst_coa" FOREIGN KEY ("dst_account_id") REFERENCES "m_chart_of_accounts"("id");

ALTER TABLE "invoices" ADD CONSTRAINT "FK_invoices_trx_rcv_d" FOREIGN KEY ("payment_id") REFERENCES "trx_receive_details"("id");

ALTER TABLE "invoices" ADD CONSTRAINT "FK_invoices_disc" FOREIGN KEY ("discount_id") REFERENCES "m_discounts"("id");

ALTER TABLE "invoices" ADD CONSTRAINT "FK_invoices_ar" FOREIGN KEY ("ar_id") REFERENCES "account_receiveables"("id");

ALTER TABLE "trx_receives" ADD CONSTRAINT "FK_trx_rcv_journals" FOREIGN KEY ("journal_id") REFERENCES "journals"("id");

ALTER TABLE "wallets" ADD CONSTRAINT "FK_wallets_coa" FOREIGN KEY ("account_id") REFERENCES "m_chart_of_accounts"("id");

ALTER TABLE "trx_receive_details" ADD CONSTRAINT "FK_trx_rcv_d_trx_rcv" FOREIGN KEY ("trx_receives_id") REFERENCES "trx_receives"("id");

ALTER TABLE "journals" ADD CONSTRAINT "FK_journals_gl" FOREIGN KEY ("general_ledger_id") REFERENCES "general_ledgers"("id");

ALTER TABLE "trx_wallets" ADD CONSTRAINT "FK_trx_wallets_wallets" FOREIGN KEY ("wallet_id") REFERENCES "wallets"("id");

ALTER TABLE "trx_wallets" ADD CONSTRAINT "FK_trx_wallets_trx_rcv" FOREIGN KEY ("trx_receives_id") REFERENCES "trx_receives"("id");

ALTER TABLE "trx_bank_transfers" ADD CONSTRAINT "FK_trx_bank_trf_banks" FOREIGN KEY ("bank_id") REFERENCES "m_banks"("id");

ALTER TABLE "trx_bank_transfers" ADD CONSTRAINT "FK_trx_bank_trf_trx_rcv" FOREIGN KEY ("trx_receives_id") REFERENCES "trx_receives"("id");

ALTER TABLE "account_receiveable_rules" ADD CONSTRAINT "FK_ar_rules_ar" FOREIGN KEY ("ar_id") REFERENCES "account_receiveables"("id");

ALTER TABLE "account_receiveable_rules" ADD CONSTRAINT "FK_ar_rules_book_periods" FOREIGN KEY ("period_id") REFERENCES "m_book_periods"("id");

ALTER TABLE "journal_details" ADD CONSTRAINT "FK_journal_d_journals" FOREIGN KEY ("journal_id") REFERENCES "journals"("id");

ALTER TABLE "journal_details" ADD CONSTRAINT "FK_journal_d_coa" FOREIGN KEY ("account_id") REFERENCES "m_chart_of_accounts"("id");

ALTER TABLE "account_payable_rules" ADD CONSTRAINT "FK_ap_rules_ap" FOREIGN KEY ("ap_id") REFERENCES "account_payables"("id");

ALTER TABLE "account_payable_rules" ADD CONSTRAINT "FK_ap_rules_book_periods" FOREIGN KEY ("period_id") REFERENCES "m_book_periods"("id");

ALTER TABLE "trx_expenses" ADD CONSTRAINT "FK_trx_expenses_journals" FOREIGN KEY ("journal_id") REFERENCES "journals"("id");

ALTER TABLE "trx_expense_details" ADD CONSTRAINT "FK_trx_expense_d_trx_expenses" FOREIGN KEY ("trx_expenses_id") REFERENCES "trx_expenses"("id");

ALTER TABLE "purchases" ADD CONSTRAINT "FK_purchases_ap" FOREIGN KEY ("ap_id") REFERENCES "account_payables"("id");

ALTER TABLE "purchases" ADD CONSTRAINT "FK_purchases_trx_expenses" FOREIGN KEY ("expense_id") REFERENCES "trx_expenses"("id");

ALTER TABLE "general_ledger_details" ADD CONSTRAINT "FK_gl_d_gl" FOREIGN KEY ("general_ledger_id") REFERENCES "general_ledgers"("id");


