select users.id, users.username, followers.user_id, followers.follower_id from users left join followers on users.id=followers.user_id;

select * from users;

select id, username, email, created_at, COALESCE(f.user_id, -1) as following from users u left join (select user_id, follower_id from followers where follower_id=3) f on u.id=f.user_id;