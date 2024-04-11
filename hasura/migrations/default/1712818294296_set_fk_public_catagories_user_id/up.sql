alter table "public"."catagories" drop constraint "catagories_user_id_fkey",
  add constraint "catagories_user_id_fkey"
  foreign key ("user_id")
  references "public"."users"
  ("id") on update restrict on delete cascade;
