alter table "public"."comments" drop constraint "comments_comment_id_fkey",
  add constraint "comments_comment_id_fkey"
  foreign key ("comment_id")
  references "public"."recipe"
  ("id") on update restrict on delete restrict;
