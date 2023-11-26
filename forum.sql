CREATE TABLE IF NOT EXISTS User (
  Id INTEGER PRIMARY KEY,
  Username TEXT NOT NULL UNIQUE,
  Password TEXT NOT NULL,
  Email TEXT NOT NULL UNIQUE,
  Avatar TEXT NOT NULL DEFAULT "",
  Role_Id INTEGER NOT NULL DEFAULT (1),
  Created TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (Role_Id) REFERENCES Role(Id)
);

CREATE TABLE IF NOT EXISTS Role (
  Id INTEGER PRIMARY KEY,
  Name TEXT NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS Session (
  UUID TEXT NOT NULL UNIQUE,
  Expires TIMESTAMP NOT NULL,
  User_Id INTEGER NOT NULL,
  FOREIGN KEY (User_Id) REFERENCES User(Id)
);

CREATE TABLE IF NOT EXISTS Post (
  Id INTEGER PRIMARY KEY,
  Title TEXT NOT NULL,
  Content_Id INTEGER NOT NULL DEFAULT 0, /* Content is in first comment */
  User_Id INTEGER NOT NULL,
  Category_Id INTEGER NOT NULL,
  Created TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  Updated TIMESTAMP NOT NULL DEFAULT "0000-00-00 00:00:00",
  Likes INTEGER NOT NULL DEFAULT 0,
  Dislikes INTEGER NOT NULL DEFAULT 0,
  Tags TEXT NOT NULL DEFAULT "",
  FOREIGN KEY (User_Id) REFERENCES User(Id)
  FOREIGN KEY (Category_Id) REFERENCES Category(Id)
  FOREIGN KEY (Content_Id) REFERENCES Comment(Id)
);

CREATE TABLE IF NOT EXISTS Comment (
  Id INTEGER PRIMARY KEY,
  User_Id INTEGER NOT NULL,
  Post_Id INTEGER NOT NULL,
  Content TEXT NOT NULL,
  Is_Post INTEGER NOT NULL DEFAULT 0,
  Created TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  Updated TIMESTAMP NOT NULL DEFAULT "0000-00-00 00:00:00",
  Likes INTEGER NOT NULL DEFAULT 0,
  Dislikes INTEGER NOT NULL DEFAULT 0,
  FOREIGN KEY (User_Id) REFERENCES User(Id),
  FOREIGN KEY (Post_Id) REFERENCES Post(Id)
);

CREATE TABLE IF NOT EXISTS Category (
  Id INTEGER PRIMARY KEY,
  Parent_Id INTEGER NOT NULL,
  Short TEXT NOT NULL UNIQUE,
  Title TEXT NOT NULL,
  Intro TEXT NOT NULL,
  FOREIGN KEY (Parent_Id) REFERENCES Category(Id)
);

CREATE TABLE IF NOT EXISTS Tag (
  Id INTEGER PRIMARY KEY,
  Name TEXT NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS Post_Reaction (
  Id INTEGER PRIMARY KEY,
  User_Id INTEGER NOT NULL,
  Post_Id INTEGER NOT NULL,
  Reaction INTEGER NOT NULL,
  FOREIGN KEY (User_Id) REFERENCES User(Id),
  FOREIGN KEY (Post_Id) REFERENCES Post(Id),
  UNIQUE (User_Id, Post_Id)
);

CREATE TABLE IF NOT EXISTS Comment_Reaction (
  Id INTEGER PRIMARY KEY,
  User_Id INTEGER NOT NULL,
  Comment_Id INTEGER NOT NULL,
  Reaction INTEGER NOT NULL,
  FOREIGN KEY (User_Id) REFERENCES User(Id),
  FOREIGN KEY (Comment_Id) REFERENCES Comment(Id),
  UNIQUE (User_Id, Comment_Id)
);

/* Structural data */
INSERT INTO Role (Id, Name) VALUES (1,'User');
INSERT INTO Role (Id, Name) VALUES (2,'Moderator');
INSERT INTO Role (Id, Name) VALUES (3,'Admin');

INSERT INTO Category (Short, Title, Intro, Parent_Id) VALUES 
('main', 'Main Board', 'Essential information for everyone', 0),
('community', 'Community General', 'Everything you need to know and share', 0),
('career', 'Career Path Guidance', "Seek advice on job interviews, internships, and network with alumni in the industry", 0),
('rules', 'Rules & Information', 'Important information about the rules of conduct on the forum', 1),
('news', 'Board News & Events', 'Forum news and events', 1),
('feedback', 'Feedback and suggestions for improvement', 'Ideas? Criticism? Feedback? This way', 1),
('intro', 'Introductions', "New here? Feel free to introduce yourself to the community!", 2),
('study', 'Study Buddy Search', "Find peers interested in teaming up for study collaborations", 2),
('chat', 'General Chat', "This and that, no matter what's it about", 2),
('lifehacks', 'Student Life Hacks', "Discover and share life hacks for managing your studies", 2),
('internship', 'Internship Opportunities', "Discover cool intern gigs, part-time hustles, and summer adventures to gain that real-world tech experience", 3),
('interviews', 'Resume and Interviews Tips', "Get the scoop on building a killer resume and nailing your interviews—no sweat!", 3),
('freelancing', 'Freelancing Advice', "All about the freelance life: making bank and juggling clients like a pro", 3),
('alumni', 'Alumni Stories', "Unfiltered career adventures, facepalms, and victories of our very own grads", 3);
