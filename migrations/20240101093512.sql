-- Create "workspaces" table
CREATE TABLE "workspaces" (
  "id" bigserial NOT NULL,
  "created_at" timestamptz NOT NULL,
  "updated_at" timestamptz NOT NULL,
  "name" text NOT NULL,
  PRIMARY KEY ("id")
);
-- Create index "idx_workspaces_name" to table: "workspaces"
CREATE UNIQUE INDEX "idx_workspaces_name" ON "workspaces" ("name");
-- Create "channels" table
CREATE TABLE "channels" (
  "id" bigserial NOT NULL,
  "created_at" timestamptz NOT NULL,
  "updated_at" timestamptz NOT NULL,
  "workspace_id" bigint NOT NULL,
  "name" text NOT NULL,
  "type" text NOT NULL,
  "status" text NOT NULL,
  "description" text NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_channels_workspace" FOREIGN KEY ("workspace_id") REFERENCES "workspaces" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create index "idx_channels_workspace_name" to table: "channels"
CREATE UNIQUE INDEX "idx_channels_workspace_name" ON "channels" ("workspace_id", "name");
-- Create "users" table
CREATE TABLE "users" (
  "id" bigserial NOT NULL,
  "created_at" timestamptz NOT NULL,
  "updated_at" timestamptz NOT NULL,
  "email" text NULL,
  PRIMARY KEY ("id")
);
-- Create index "idx_users_email" to table: "users"
CREATE UNIQUE INDEX "idx_users_email" ON "users" ("email");
-- Create "workspace_members" table
CREATE TABLE "workspace_members" (
  "id" bigserial NOT NULL,
  "created_at" timestamptz NOT NULL,
  "updated_at" timestamptz NOT NULL,
  "user_id" bigint NOT NULL,
  "workspace_id" bigint NOT NULL,
  "tag_name" text NOT NULL,
  "display_name" text NOT NULL,
  "name_token" tsvector NULL GENERATED ALWAYS AS (to_tsvector('english'::regconfig, ((tag_name || ' '::text) || display_name))) STORED,
  "role" text NOT NULL,
  "status" text NOT NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_users_workspace_members" FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT "fk_workspaces_members" FOREIGN KEY ("workspace_id") REFERENCES "workspaces" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create index "idx_workspace_members_name_token" to table: "workspace_members"
CREATE INDEX "idx_workspace_members_name_token" ON "workspace_members" USING gin ("name_token");
-- Create index "idx_workspace_members_user_id" to table: "workspace_members"
CREATE INDEX "idx_workspace_members_user_id" ON "workspace_members" ("user_id");
-- Create index "idx_workspace_members_workspace_tag_name" to table: "workspace_members"
CREATE UNIQUE INDEX "idx_workspace_members_workspace_tag_name" ON "workspace_members" ("workspace_id", "tag_name");
-- Create "channel_members" table
CREATE TABLE "channel_members" (
  "channel_id" bigint NOT NULL,
  "workspace_member_id" bigint NOT NULL,
  PRIMARY KEY ("channel_id", "workspace_member_id"),
  CONSTRAINT "fk_channel_members_channel" FOREIGN KEY ("channel_id") REFERENCES "channels" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT "fk_channel_members_workspace_member" FOREIGN KEY ("workspace_member_id") REFERENCES "workspace_members" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create "messages" table
CREATE TABLE "messages" (
  "id" bigserial NOT NULL,
  "created_at" timestamptz NOT NULL,
  "updated_at" timestamptz NOT NULL,
  "channel_id" bigint NULL,
  "sender_id" bigint NULL,
  "content" text NOT NULL,
  "content_token" tsvector NULL GENERATED ALWAYS AS (to_tsvector('english'::regconfig, content)) STORED,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_messages_channel" FOREIGN KEY ("channel_id") REFERENCES "channels" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT "fk_workspace_members_messages" FOREIGN KEY ("sender_id") REFERENCES "workspace_members" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create index "idx_messages_channel_sender" to table: "messages"
CREATE INDEX "idx_messages_channel_sender" ON "messages" ("channel_id", "sender_id");
-- Create index "idx_messages_content_token" to table: "messages"
CREATE INDEX "idx_messages_content_token" ON "messages" USING gin ("content_token");
-- Create index "idx_messages_sender_id" to table: "messages"
CREATE INDEX "idx_messages_sender_id" ON "messages" ("sender_id");
-- Create "threads" table
CREATE TABLE "threads" (
  "id" bigserial NOT NULL,
  "created_at" timestamptz NOT NULL,
  "updated_at" timestamptz NOT NULL,
  "message_id" bigint NULL,
  "sender_id" bigint NULL,
  "content" text NOT NULL,
  "content_token" tsvector NULL GENERATED ALWAYS AS (to_tsvector('english'::regconfig, content)) STORED,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_threads_message" FOREIGN KEY ("message_id") REFERENCES "messages" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT "fk_workspace_members_threads" FOREIGN KEY ("sender_id") REFERENCES "workspace_members" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create index "idx_threads_content_token" to table: "threads"
CREATE INDEX "idx_threads_content_token" ON "threads" USING gin ("content_token");
-- Create index "idx_threads_message_id" to table: "threads"
CREATE INDEX "idx_threads_message_id" ON "threads" ("message_id");
-- Create index "idx_threads_sender_id" to table: "threads"
CREATE INDEX "idx_threads_sender_id" ON "threads" ("sender_id");
