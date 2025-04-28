-- Insert dummy users
INSERT INTO users (name, email, password, role) VALUES
    ('Alice Johnson', 'alice@example.com', '$2y$10$X31G/./0HmxpxOcUHEsYfOuNvDyhBYY7US6nxwFZpa1A1AZzGn4a.', 'student'),
    ('Bob Smith', 'bob@example.com', '$2y$10$X31G/./0HmxpxOcUHEsYfOuNvDyhBYY7US6nxwFZpa1A1AZzGn4a.', 'instructor'),
    ('Charlie Brown', 'charlie@example.com', '$2y$10$X31G/./0HmxpxOcUHEsYfOuNvDyhBYY7US6nxwFZpa1A1AZzGn4a.', 'admin');

-- Insert dummy courses
INSERT INTO courses (title, description, instructor_id, price, category) VALUES
    ('Introduction to SQL', 'Learn the basics of SQL.', 2, 49.99, 'Database'),
    ('Advanced Golang', 'Deep dive into Golang features.', 2, 79.99, 'Programming');

-- Insert dummy enrollments
INSERT INTO enrollments (student_id, course_id) VALUES
    (1, 1),
    (1, 2);

-- Insert dummy materials
INSERT INTO materials (course_id, title, content, file_url) VALUES
    (1, 'SQL Basics', 'Introduction to SQL syntax.', 'https://example.com/sql_basics.pdf'),
    (2, 'Concurrency in Golang', 'Understanding Goroutines.', 'https://example.com/goroutines.pdf');

-- Insert dummy payments
INSERT INTO payments (student_id, course_id, amount, status) VALUES
    (1, 1, 49.99, 'completed'),
    (1, 2, 79.99, 'pending');
