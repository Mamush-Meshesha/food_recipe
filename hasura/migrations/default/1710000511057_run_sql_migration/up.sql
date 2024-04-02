CREATE OR REPLACE FUNCTION public.recipe_check(recipe_row recipe, hasura_session json)
RETURNS boolean
LANGUAGE sql
STABLE
AS $$
SELECT EXISTS (
    SELECT 1
    FROM checked A
    WHERE A.user_id = CAST((hasura_session ->> 'x-hasura-user-id') as INTEGER) AND A.rec_id = recipe_row.id
);
$$;
