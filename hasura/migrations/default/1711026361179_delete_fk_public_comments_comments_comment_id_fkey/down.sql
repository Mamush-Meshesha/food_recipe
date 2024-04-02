alter table "public"."comments"
  add constraint "comments_comment_id_fkey"
  foreign key ("comment_id")
  references "public"."comments"
  ("id") on update restrict on delete restrict;
