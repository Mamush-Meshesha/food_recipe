alter table "public"."catagories"
  add constraint "catagories_id_fkey"
  foreign key ("id")
  references "public"."recipe"
  ("id") on update restrict on delete restrict;
