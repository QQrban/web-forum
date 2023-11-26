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
  /*Content TEXT NOT NULL,*/
  User_Id INTEGER NOT NULL,
  Category_Id INTEGER NOT NULL,
  Created TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  Updated TIMESTAMP NOT NULL DEFAULT "0000-00-00 00:00:00",
  Likes INTEGER NOT NULL DEFAULT 0,
  Dislikes INTEGER NOT NULL DEFAULT 0,
  Tags TEXT NOT NULL DEFAULT "",
  FOREIGN KEY (User_Id) REFERENCES User(Id)
  FOREIGN KEY (Category_Id) REFERENCES Category(Id)
);

CREATE TABLE IF NOT EXISTS Comment (
  Id INTEGER PRIMARY KEY,
  User_Id INTEGER NOT NULL,
  Post_Id INTEGER NOT NULL,
  Content TEXT NOT NULL,
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

/*CREATE TABLE IF NOT EXISTS Post_Tag (
  Id INTEGER PRIMARY KEY,
  Post_Id INTEGER NOT NULL,
  Tag_Id INTEGER NOT NULL,
  FOREIGN KEY (Post_Id) REFERENCES Post(Id),
  FOREIGN KEY (Tag_Id) REFERENCES Tag(Id)
);*/

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

/* Temporary example data */
INSERT INTO User (Username, Password, Email, Role_Id, Avatar, Created) VALUES 
('tvooglai','$2a$04$66TioZ12kSBSpjCShW6UYOCo81o5DAAcejdvkBCSCghUPWinAXZLG','toomas.vooglaid@gmail.com',3,'/static/images/avatars/santa.svg','2023-10-09T08:18:51Z'),
('kramazan','$2a$04$N9if82YuPoa7E56J1Bae/eq3d8aYapOAFQDZrMmULrgp03Z1yV61i','kramazan@kood.ee',3,'/static/images/avatars/arab.svg','2023-10-09T08:18:51Z'),
('CodeCool','$2a$04$t5tXiYQGuNbczilsac1OruZZYQXq/Rd/LaWw5p1UnKsSQBJyXmzke','codecool@kood.ee',2,'/static/images/avatars/batman.svg','2023-11-02T08:18:51Z'),
('AvoCoder','$2a$04$ewHvhWJjp2GVb.jhbeo6CuB7/nZXcSLnUQPqVjpl5TlighWu6PrGK','avocoder@kood.ee',2,'/static/images/avatars/avocado.svg','2023-11-04T08:18:51Z'),
('koodMember','$2a$04$WJi72QoqmCh5R0NbaWhKoeC.Jh.LgmxBIFvV.a.UJuyaa5OF0FzzG','koodMember@kood.ee',1,'/static/images/avatars/manglasses.svg','2023-11-04T08:18:51Z'),
('albert','$2a$04$rlXE1KygSgjlzY9od9Y/ueuCC69XHKOmAwdVVKzCddZURk8QKnEiG','albert@kood.ee',1,'/static/images/avatars/einstein.svg','2023-11-05T08:18:51Z'),
('nipitiri','$2a$04$Yq4uuwyqU5ienSq6dGWbEuOwa9NTWvB4M49oDp0ZsyCGBYL.ripAy','nipitiri@kood.ee',1,'/static/images/avatars/cactus.svg','2023-11-06T08:18:51Z');

INSERT INTO Post (Title, User_Id, Category_Id, Created, Likes, Dislikes, Tags) VALUES 
('Rules of Conduct', 1, 4, "2023-10-10 09:25:47", 1, 2,"rules,conduct"),
('Forum Etiquette: The Basics', 1, 4, "2023-10-11 09:05:07", 0, 0,"rules,conduct,etiquette"),
('Prohibited Content: What Not to Share', 1, 4, "2023-10-12 10:05:07", 0, 3,"rules,prohibited"),
('Stay on Topic: Guidelines for Relevant Posting', 1, 4,"2023-10-15 10:05:07", 3, 1,"rules,guidelines"),
('Reporting Violations: Keeping the Forum Safe', 1, 4, "2023-10-18 20:15:23", 1, 1,"rules,reporting");

INSERT INTO Comment (User_Id, Post_Id, Content, Created) VALUES 
(1,1,'1. Be nice to each other', "2023-10-10 09:25:47"),
(1,2,'Before you post, remember to always treat fellow forum members with respect and courtesy.', "2023-10-11 09:05:07"),
(1,3,'Avoid posting any material that is knowingly false, defamatory, inaccurate, abusive, vulgar, hateful, harassing, obscene, profane, sexually oriented, threatening, invasive of a person''s privacy, or otherwise violative of any law.', "2023-10-12 10:05:07"),
(1,4,'Ensure your contributions are relevant to the thread’s subject and add value to the discussion.', "2023-10-15 10:05:07"),
(1,5,'If you encounter a post that violates our community guidelines, please report it to the moderation team immediately.', "2023-10-18 20:15:23");

INSERT INTO Post (Title, User_Id, Category_Id, Created, Likes, Dislikes, Tags) VALUES 
('Board Update: Fresh Features', 2, 5, "2023-10-13 20:15:23", 2, 0,"updates,features"),
('Monthly Event Schedule', 2, 5, "2023-10-14 02:10:21", 5, 1,"updates,events,schedule"),
('Maintenance Downtime', 2, 5, "2023-10-16 04:11:01", 0, 4,"updates,maintenance"),
('Community Highlights', 2, 5, "2023-10-16 09:11:01", 5, 2,"updates,highlights");

INSERT INTO Comment (User_Id, Post_Id, Content, Created) VALUES 
(2,6,'Check out the new forum features just released!', "2023-10-13 20:15:23"),
(2,7,'Stay updated with our forum event calendar for this month.', "2023-10-14 02:10:21"),
(2,8,'The forum will be undergoing maintenance on Sunday at 3 PM.', "2023-10-16 04:11:01"),
(2,9,'Our community spotlight shines on outstanding members this week.', "2023-10-16 09:11:01");

INSERT INTO Post (Title, User_Id, Category_Id, Created, Likes, Dislikes, Tags) VALUES 
('Feature Request: User Profiles', 3, 6, "2023-11-06 09:11:01", 2, 0, "features,users,profiles"),
('Site Speed Issues', 3, 6, "2023-11-06 10:11:01", 0, 1, "issues,speed");

INSERT INTO Comment (User_Id, Post_Id, Content, Created) VALUES 
(3,10,'Would love the ability to customize our profiles more. What does everyone think?', "2023-11-06 09:11:01"),
(3,11,'Has anyone else experienced slower forum load times lately?', "2023-11-06 10:11:01");

INSERT INTO Post (Title, User_Id, Category_Id, Created, Likes, Dislikes, Tags) VALUES 
('Hello from a Newbie!', 4, 7, "2023-11-06 10:17:00", 0, 0, "introduction,newbie");

INSERT INTO Comment (User_Id, Post_Id, Content, Created) VALUES 
(4,12,"Hi everyone, I'm Alex, excited to join this vibrant community!", "2023-11-06 10:17:00");

INSERT INTO Post (Title, User_Id, Category_Id, Created, Likes, Dislikes, Tags) VALUES 
('Looking for a Math Study Group', 3, 8, "2023-11-06 11:27:04", 0, 0, "study,math"),
('Chemistry Exam Prep Partner Needed', 4, 8, "2023-11-06 13:27:04", 0, 0, "study,chemistry");

INSERT INTO Comment (User_Id, Post_Id, Content, Created) VALUES 
(3,13,'I’m on the hunt for fellow math enthusiasts to dive into algebra and calculus. Anyone interested?', "2023-11-06 11:27:04"),
(4,14,'In search of a study partner for upcoming chemistry finals. We can share notes and quiz each other!', "2023-11-06 13:27:04");

INSERT INTO Post (Title, User_Id, Category_Id, Created, Likes, Dislikes, Tags) VALUES 
('Weekend Plans?', 2, 9, "2023-11-06 15:02:42", 2, 0, "plans,weekend"),
('Book Recommendations', 1, 9, "2023-11-06 16:03:42", 0, 1, "books,reading"),
('Favorite Coffee Spots', 3, 9, "2023-11-06 16:05:44", 3, 0, "coffee,cafe");

INSERT INTO Comment (User_Id, Post_Id, Content, Created) VALUES 
(2,15,"What's everyone up to this weekend? Any fun plans or just relaxing at home?", "2023-11-06 15:02:42"),
(1,16,'Just finished a great novel and looking for something new to dive into. Suggestions?', "2023-11-06 16:03:42"),
(3,17,'On a quest to find the best coffee in town. Where do you get your caffeine fix?', "2023-11-06 16:05:44");

INSERT INTO Post (Title, User_Id, Category_Id, Created, Likes, Dislikes, Tags) VALUES 
('Maximize Your Study Sessions', 1, 10, "2023-11-06 16:15:04", 4, 1, "study,hacks"),
('Every time you order pizza...', 3,10, "2023-11-06 16:15:05", 2, 2, "pizza,food");

INSERT INTO Comment (User_Id, Post_Id, Content, Created) VALUES 
(1,18,'Share your favorite study hacks! Mine is using a timer to break work into focused intervals.', "2023-11-06 16:15:04"),
(3,19,"Every time you order pizza, you enter a world of cheesy goodness, where every bite dances in your mouth. 

It's a box full of happiness that turns any ordinary moment into a party. 

Seriously, who knew a circle made of dough, sauce, and toppings could make life feel so round and complete?", "2023-11-06 16:15:05"),
(4,19,"Speaking out loud about certain matters might not always be the best approach for various reasons. It can expose your vulnerabilities, making you feel unnecessarily exposed in situations that don't warrant such openness. It might also invite conflict, especially in scenarios where opinions are deeply divided, and consensus seems impossible. Publicly discussing sensitive topics can inadvertently offend others who have different perspectives, creating tension or discomfort around you. 

Openly talking about personal matters might not only concern you but also others involved in these matters, breaching their privacy and possibly affecting your relationships. Discussing plans or ideas prematurely can also lead to unnecessary scrutiny, criticism, or theft of ideas, especially in competitive environments. Furthermore, vocalizing issues without proper context or understanding can spread misinformation, contributing to confusion and misunderstanding. In some cases, speaking about something without complete information can lead to premature judgments, potentially leading to regret later. 

Public discussions about others’ issues or private matters, even indirectly, can lead to trust issues, as it shows a lack of respect for others' privacy.
", '2023-11-07T03:00:40Z'),
(3,19,"I agree with you. I think it's important to be mindful of what we say and how we say it. It's also important to be mindful of who we say it to. And, I'm Sorry!", '2023-11-07T04:00:40Z');

INSERT INTO Post (Title, User_Id, Category_Id, Created, Likes, Dislikes, Tags) VALUES 
('Summer Tech Internships Open Now', 2, 11, "2023-11-06 16:17:15", 2, 0, "internships,summer"),
('Part-Time Developer Roles for Students', 3, 11, "2023-11-06 18:07:24", 4, 2, "part-time,developer,student");

INSERT INTO Comment (User_Id, Post_Id, Content, Created) VALUES 
(2,20,'Just spotted some amazing summer internship opportunities for tech enthusiasts. Check them out before they close!', "2023-11-06 16:17:15"),
(2,21,'Companies are looking for part-time junior developers - a perfect way to earn while you learn!', "2023-11-06 18:07:24");

INSERT INTO Post (Title, User_Id, Category_Id, Created, Likes, Dislikes, Tags) VALUES 
('Crafting the Perfect Tech Resume', 4, 12, "2023-11-06 18:08:24", 2, 0, "resume,tech");

INSERT INTO Comment (User_Id, Post_Id, Content, Created) VALUES 
(4,22,'Let’s discuss the key ingredients for a tech resume that stands out!', "2023-11-06 18:08:24");

INSERT INTO Post (Title, User_Id, Category_Id, Created, Likes, Dislikes, Tags) VALUES 
('Freelancing Essentials', 1, 13, "2023-11-06 18:08:26", 3, 1, "freelancing,essentials");

INSERT INTO Comment (User_Id, Post_Id, Content, Created) VALUES 
(1,23,'Starting your freelance journey? Here are some essential tips to help you thrive.', "2023-11-06 18:08:26");

INSERT INTO Post (Title, User_Id, Category_Id, Created, Likes, Dislikes, Tags) VALUES 
('Alumni Success: From Graduation to CEO', 3, 14, "2023-11-08 16:01:56", 3, 0, "alumni,success");

INSERT INTO Comment (User_Id, Post_Id, Content, Created) VALUES 
(3,24,'Read about one of our grad’s journey from the classroom to the boardroom.', "2023-11-08 16:01:56");

/* Post reactions */

/*INSERT INTO Post_Reaction (Post_Id, User_Id, Reaction) VALUES
(1,2,1),(1,3,-1),(1,4,-1),
(3,3,-1),(3,5,-1),(3,4,-1),
(4,2,1),(4,3,1),(4,5,1),(4,4,-1),
5|1|1|1
6|2|2|0
7|2|5|1
8|2|0|4
9|2|5|2
10|3|2|0
11|3|0|1
15|2|2|0
16|1|0|1
17|3|3|0
18|1|4|1
19|3|2|2
20|2|2|0
21|3|4|2
22|4|2|0
23|1|3|1
24|3|3|0)*/