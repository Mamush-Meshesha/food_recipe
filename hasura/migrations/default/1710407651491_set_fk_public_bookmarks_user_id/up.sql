alter table "public"."bookmarks"
  add constraint "bookmarks_user_id_fkey"
  foreign key ("user_id")
  references "public"."users"
  ("id") on update restrict on delete restrict;
