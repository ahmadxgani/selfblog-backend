CREATE TABLE "accounts" (
  "id" smallint GENERATED ALWAYS AS IDENTITY UNIQUE, 
  "email" varchar(255) NOT NULL UNIQUE,
  "username" varchar(16) NOT NULL UNIQUE,
  "password" varchar NOT NULL,
  "full_name" varchar(24) NOT NULL,
  "image" varchar,
  "created_at" timestamptz NOT NULL DEFAULT now(),
  "is_admin" smallint NOT NULL DEFAULT 0
);

CREATE TABLE "posts" (
  "id" smallserial PRIMARY KEY,
  "title" varchar(25) NOT NULL,
  "content" varchar NOT NULL,
  "slug" varchar NOT NULL,
  "draft" boolean NOT NULL DEFAULT FALSE,
  "likes" smallint NOT NULL DEFAULT 0,
  "created_at" timestamptz NOT NULL DEFAULT now(),
  "updated_at" timestamptz NOT NULL DEFAULT now(),

  CONSTRAINT "unique_title_slug" UNIQUE ("title", "slug")
);

CREATE TABLE "tags" (
  "id" smallint GENERATED ALWAYS AS IDENTITY UNIQUE, 
  "name" varchar NOT NULL UNIQUE DEFAULT 'uncategorized',
  "description" varchar NOT NULL DEFAULT 'Anything, just my arbitary posts'
); 

CREATE TABLE "comments" (
  "id" smallserial PRIMARY KEY,
  "content" varchar NOT NULL,
  "username" varchar NOT NULL,
  "post_id" smallint NOT NULL,
  "account_id" smallint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT now(),
  "updated_at" timestamptz NOT NULL DEFAULT now(),

  CONSTRAINT "fk_post" FOREIGN KEY ("post_id") REFERENCES "posts" ("id"),
  CONSTRAINT "fk_account" FOREIGN KEY ("account_id") REFERENCES "accounts" ("id")
);

CREATE TABLE "tag_posts" (
  "id" smallserial PRIMARY KEY,
  "tag_id" smallint NOT NULL,
  "post_id" smallserial NOT NULL,

  CONSTRAINT "unique_tag_post" UNIQUE ("tag_id", "post_id"),
  CONSTRAINT "fk_tag" FOREIGN KEY ("tag_id") REFERENCES "tags" ("id"),
  CONSTRAINT "fk_post" FOREIGN KEY ("post_id") REFERENCES "posts" ("id")
);

CREATE EXTENSION IF NOT EXISTS pgcrypto;

INSERT INTO accounts (username, password, full_name, email) VALUES ('ahmadxgani', crypt('qwerty', gen_salt('bf')), 'Ahmad Gani', 'foo@bar.com');
INSERT INTO tags (name, description) VALUES ('diary', 'My personal diary book.');
INSERT INTO tags (name, description) VALUES ('lifehack', 'Just a pro-tip that might cut down your time in solving a problem.');
INSERT INTO tags (name, description) VALUES ('tutorial', 'Programming tutorial, hope it helpful!');

