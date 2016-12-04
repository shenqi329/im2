use db_im;

select * from t_token;

select count(*) from t_token where t_token_user_id = "1";

select max(t_token_id) from t_token;

SET SQL_SAFE_UPDATES = 0;

delete from t_session_map;
delete from t_session;

select * from t_session_map;

select * from t_session;
select * from t_message;
select * from t_token;

select count(t_message_index) from t_message where t_message_session_id = 32;