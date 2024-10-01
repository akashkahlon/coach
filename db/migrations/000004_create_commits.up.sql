create table commits (
	id serial primary key,
	repository_id int not null,
	sha text not null,
	author text not null,
	author_email text not null,
	commit_message text not null,
	github_data_all jsonb not null,
	created_at timestamp not null default current_timestamp,
	updated_at timestamp not null default current_timestamp
);