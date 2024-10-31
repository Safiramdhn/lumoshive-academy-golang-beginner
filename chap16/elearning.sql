
CREATE TYPE status_enum AS enum ('active', 'deleted');
create type attendance_status_enum as enum ('ontime', 'late');
create  type assignment_status_enum as enum ('not_started', 'submited','missing');

CREATE TABLE users (
    id SERIAL PRIMARY KEY NOT NULL UNIQUE,
    email VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR NOT NULL,
    status status_enum DEFAULT 'active',
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);



create table admins (
	id SERIAL primary key not null unique,
	first_name varchar(255) not null,
	last_name varchar(255) not null,
	user_id int references users(id) not null,
	status status_enum DEFAULT 'active',
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

create table student (
	id SERIAL primary key not null unique,
	first_name varchar (255) not null,
	last_name varchar (255) not null,
	user_id int references users(id) not null,
	status status_enum DEFAULT 'active',
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

create table mentor (
 	id SERIAL primary key not null unique,
	first_name varchar (255) not null,
	last_name varchar (255) not null,
	user_id int references users(id) not null,
	added_by int references admins(id) not null,
	status status_enum DEFAULT 'active',
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

create table material (
	id SERIAL primary key not null unique,
	title varchar(255),
	description varchar,
	media_url varchar,
	added_by int references admins(id) not null,
	status status_enum DEFAULT 'active',
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

create table schedule (
	id SERIAL primary key not null unique,
	date date not null,
	time time not null,
	added_by int references admins(id) not null,
	status status_enum DEFAULT 'active',
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

create table classes (
	id SERIAL primary key not null unique,
	title varchar(255),
	description varchar,
	added_by int references admins(id) not null,
	mentor_id int references mentor(id) not null,
	schedule_id int references schedule(id) not null,
	material_id int references material(id) not null,
	status status_enum DEFAULT 'active',
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

create table announcement (
	id SERIAL primary key not null unique,
	title varchar(255),
	description varchar,
	added_by int references admins(id) not null,
	class_id int references classes(id),
	status status_enum DEFAULT 'active',
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

create table student_class (
	id SERIAL primary key not null unique,
	student_id int references student(id) not null,
	class_id int references classes(id) not null,
	status status_enum DEFAULT 'active',
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

create table student_attendance (
	id SERIAL primary key not null unique,
	student_id int references student(id) not null,
	student_class_id int references student_class(id) not null,
	datetime timestamp,
	attendance_status attendance_status_enum,
	status status_enum DEFAULT 'active',
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

create table mentor_attendance (
	id SERIAL primary key not null unique,
	mentor_id int references mentor(id) not null,
	class_id int references classes(id) not null,
	datetime timestamp,
	attendance_status attendance_status_enum,
	status status_enum DEFAULT 'active',
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);


create table assignments (
	id SERIAL primary key not null unique,
	class_id int references classes(id) not null,
	deadline timestamp not null,
	description varchar,
	title varchar(255) not null,
	status status_enum DEFAULT 'active',
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

create table student_assignment (
	id SERIAL primary key not null unique,
	assignment_id int references assignments(id) not null,
	student_id int references student(id) not null,
	submit_date timestamp,
	assignment_status assignment_status_enum default 'not_started',
	score real,
	status status_enum DEFAULT 'active',
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

create table leaderboard (
	id SERIAL primary key not null unique,
	student_id int references student(id) not null,
	class_id int references classes(id) not null,
	total_score real,
	total_attendance int,
	status status_enum DEFAULT 'active',
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

INSERT INTO users (email, password, status) VALUES
('admin1@example.com', 'password123', 'active'),
('student1@example.com', 'password123', 'active'),
('student2@example.com', 'password123', 'active'),
('student3@example.com', 'password123', 'active'),
('student4@example.com', 'password123', 'active'),
('student5@example.com', 'password123', 'active'),
('mentor1@example.com', 'password123', 'active'),
('mentor2@example.com', 'password123', 'active'),
('mentor3@example.com', 'password123', 'active'),
('mentor4@example.com', 'password123', 'active'),
('mentor5@example.com', 'password123', 'active');
select * from users
-- select id from users where email = 'admin1@example.com' and password = 'password123'

INSERT INTO admins (first_name, last_name, user_id) VALUES
('Admin', 'One', 1),
-- SELECT * FROM admins
-- UPDATE Admins SET status = 'deleted' WHERE id = 2

INSERT INTO student (first_name, last_name, user_id) VALUES
('Student', 'One', 2),
('Student', 'Two', 3),
('Student', 'Three', 4),
('Student', 'Four', 5),
('Student', 'Five', 6);

INSERT INTO mentor (id, first_name, last_name, user_id, added_by) VALUES
('Mentor', 'One', 7, 1),
('Mentor', 'Two', 8, 1),
('Mentor', 'Three', 9, 1),
('Mentor', 'Four', 10, 1),
('Mentor', 'Five', 11, 1);


INSERT INTO material (title, description, media_url, added_by) VALUES
('Intro to Programming', 'Basic programming concepts', 'https://example.com/material1', 1),
('Data Structures', 'Introduction to data structures', 'https://example.com/material2', 1),
('Algorithms', 'Sorting and searching algorithms', 'https://example.com/material3', 1),
('Database Basics', 'Introduction to SQL and databases', 'https://example.com/material4', 1),
('Web Development', 'HTML, CSS, JavaScript basics', 'https://example.com/material5', 1);


INSERT INTO schedule (date, time, added_by) VALUES
('2024-11-01', '10:00:00', 1),
('2024-11-02', '11:00:00', 1),
('2024-11-03', '12:00:00', 1),
('2024-11-04', '13:00:00', 1),
('2024-11-05', '14:00:00', 1);

INSERT INTO classes (title, description, added_by, mentor_id, schedule_id, material_id) VALUES
('Programming 101', 'Introduction to programming', 1, 1, 1, 1),
('Data Structures 101', 'Basics of data structures', 1, 2, 2, 2),
('Algorithms 101', 'Sorting and searching', 1, 3, 3, 3),
('Database 101', 'Database basics', 1, 4, 4, 4),
('Web Development 101', 'Introduction to web development', 1, 5, 5, 5);

INSERT INTO announcement (title, description, added_by, class_id) VALUES
('Class Rescheduled', 'Programming 101 class rescheduled to next week', 1, 1),
('New Material Added', 'New material for Data Structures', 1, 2),
('Guest Lecture', 'Guest lecture on algorithms', 1, 3),
('Extra Session', 'Extra session for Database basics', 1, 4),
('Project Week', 'Web Development project week announced', 1, 5);
