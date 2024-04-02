alter table "public"."rating" alter column "user_id" set not null;
alter table "public"."rating" add constraint "rating_user_id_key" unique ("user_id");
