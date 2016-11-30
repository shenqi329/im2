use sys;

CREATE DATABASE IF NOT EXISTS db_im DEFAULT CHARSET utf8 COLLATE utf8_general_ci;

grant select on db_im.* to im_select@`%` identified by 'im_select';

grant select,update on db_im.* to im_update@`%` identified by 'im_update';

grant select,update,delete,insert on db_im.* to im_connect@`%` identified by 'im_connect';

grant all privileges on db_im.* to im_dba@`%` identified by 'im_dba';

use db_im;


SHOW VARIABLES LIKE 'event_scheduler';

SET GLOBAL event_scheduler = ON;
SET GLOBAL event_scheduler = ON;


CREATE DATABASE IF NOT EXISTS db_easynote DEFAULT CHARSET utf8 COLLATE utf8_general_ci;

grant select on db_easynote.* to easynote_select@`%` identified by 'easynote_select';

grant select,update on db_easynote.* to easynote_update@`%` identified by 'easynote_update';

grant select,update,delete,insert on db_easynote.* to easynote_connect@`%` identified by 'easynote_connect';

grant all privileges on db_easynote.* to easynote_dba@`%` identified by 'easynote_dba';