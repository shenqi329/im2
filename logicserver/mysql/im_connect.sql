use db_im;

select * from t_token;

select count(*) from t_token where t_token_user_id = "1";

select max(t_token_id) from t_token;

SET SQL_SAFE_UPDATES = 0;

delete from t_session_map;
delete from t_session;

select * from t_session_map;

select * from t_session;

select * from t_token;

select * from t_token limit 0;

select max(t_message_index) from t_message where t_message_session_id = 32;

select * from t_message where t_message_session_id = 32;

select max(t_message_index) from t_message where t_message_session_id = 32;

select * from t_message;
select * from t_message order by t_message_index desc limit 1000;
select * from t_session order by t_session_id desc limit 1000;
select * from t_session_map order by t_session_map_id desc limit 1000;





