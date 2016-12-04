use db_im;

select * from t_token;

select count(*) from t_token where t_token_user_id = "1";

select max(t_token_id) from t_token;