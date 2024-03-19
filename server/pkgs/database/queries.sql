select u.user_id, u.username, u.user_id, f.follower_id from chatterbox.user u left join follower f on u.user_id=f.user_id;

select * from chatterbox.user;

select u.user_id, username, email, created_at, COALESCE(f.user_id, 0) as following from chatterbox.user u left join (select user_id, follower_id from follower where follower_id=0) f on u.user_id=f.user_id;