alter table "public"."checked"
  add constraint "checked_rec_id_fkey"
  foreign key ("rec_id")
  references "public"."recipe"
  ("id") on update restrict on delete restrict;
