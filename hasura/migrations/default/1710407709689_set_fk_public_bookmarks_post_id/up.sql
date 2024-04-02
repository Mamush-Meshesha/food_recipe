alter table "public"."bookmarks"
  add constraint "bookmarks_post_id_fkey"
  foreign key ("post_id")
  references "public"."recipe"
  ("id") on update restrict on delete restrict;
