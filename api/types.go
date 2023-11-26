package api

import (
	"database/sql"
	"log"
	"time"
)

type Response struct {
	Message string
	Role    string
	Ok      bool
}

type Session struct {
	UUID    string `json:"uuid"`
	Expires string `json:"expires"`
	User_Id int64  `json:"user_id"`
}

func (s *Session) Create() error {
	_, err := DB.Exec("INSERT INTO Session(UUID, Expires, User_Id) VALUES(?,?,?)",
		s.UUID, s.Expires, s.User_Id)
	return err
}

func (s *Session) Delete() error {
	_, err := DB.Exec("DELETE FROM Session WHERE UUID=?", s.UUID)
	return err
}

func DeleteSession(uuid string) error {
	_, err := DB.Exec("DELETE FROM Session WHERE UUID=?", uuid)
	return err
}

func (s *Session) Get() error {
	err := DB.QueryRow("SELECT UUID, Expires, User_Id FROM Session WHERE UUID=?", s.UUID).Scan(&s.UUID, &s.Expires, &s.User_Id)
	return err
}

func GetSession(uuid string) (Session, error) {
	s := Session{UUID: uuid}
	err := DB.QueryRow("SELECT UUID, Expires, User_Id FROM Session WHERE UUID=?", uuid).Scan(&s.UUID, &s.Expires, &s.User_Id)
	return s, err
}

func (s *Session) Update() error {
	_, err := DB.Exec("UPDATE Session SET Expires=? WHERE UUID=?", s.Expires, s.UUID)
	return err
}

func flushExpiredSessions() error {
	res, err := DB.Exec("DELETE FROM Session WHERE Expires < ?", time.Now().Local().Format("2006-01-02 15:04:05"))
	if err != nil {
		return err
	}
	n, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if n > 0 {
		log.Printf("Flushed %d expired sessions\n", n)
	}
	return nil
}

type SessionWithRole struct {
	UUID    string `json:"uuid"`
	Expires string `json:"expires"`
	User_Id int64  `json:"user_id"`
	Role    string `json:"role"`
}

func getSessionWithRole(uuid string) (SessionWithRole, error) {
	s := SessionWithRole{UUID: uuid}
	err := DB.QueryRow("SELECT Expires, User_Id, Role.Name FROM Session, User, Role WHERE UUID=? AND User_Id = User.Id AND Role_Id = Role.Id", uuid).Scan(&s.Expires, &s.User_Id, &s.Role)
	return s, err
}

func totalOnline() int {
	var count int
	err := DB.QueryRow("SELECT count(DISTINCT User_Id) FROM Session").Scan(&count)
	if err != nil {
		log.Fatal(err)
	}
	return count
}

func DeleteAllSessionsFor(userID int64) error {
	_, err := DB.Exec("DELETE FROM Session WHERE User_Id=?", userID)
	return err
}

// --------------------------------------------
type User struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Avatar   string `json:"avatar"`
	Role_Id  int64  `json:"role_id"`
	Created  string `json:"created"`
}

func (u *User) Create() (int64, error) {
	res, err := DB.Exec("INSERT INTO User(Username, Password, Email, Avatar, Role_Id) VALUES(?,?,?,?,?)",
		u.Username, u.Password, u.Email, u.Avatar, u.Role_Id)
	if err != nil {
		return 0, err
	}
	u.Id, err = res.LastInsertId()
	return u.Id, err
}

func (u *User) Update() error {
	_, err := DB.Exec("UPDATE User SET Username=?, Password=?, Email=?, Avatar=?, Role_Id=? WHERE Id=?",
		u.Username, u.Password, u.Email, u.Avatar, u.Role_Id, u.Id)
	return err
}

func (u *User) Delete() error {
	_, err := DB.Exec("DELETE FROM User WHERE Id=?", u.Id)
	return err
}

func (u *User) Get() error {
	err := DB.QueryRow("SELECT Id, Username, Password, Email, Avatar, Role_Id, Created FROM User WHERE Id=?",
		u.Id).Scan(&u.Id, &u.Username, &u.Password, &u.Email, &u.Avatar, &u.Role_Id, &u.Created)
	return err
}

// Additional methods
func (u *User) duplicateRegister() (bool, error) {
	var count int
	err := DB.QueryRow("SELECT count(*) FROM User WHERE Username=? or Email=?", u.Username, u.Email).Scan(&count)
	if err == nil {
		return count == 0, nil
	}
	return false, err
}

func (u *User) GetByUsername() error {
	err := DB.QueryRow("SELECT Id, Password, Email, Avatar, Role_Id, Created FROM User WHERE Username=?",
		u.Username).Scan(&u.Id, &u.Password, &u.Email, &u.Avatar, &u.Role_Id, &u.Created)
	return err
}

func (u *User) UsernameExists() bool {
	var id int64
	err := DB.QueryRow("SELECT Id FROM User WHERE Username = ?", u.Username).Scan(&id)
	return err == nil
}

func (u *User) EmailExists() bool {
	var id int64
	err := DB.QueryRow("SELECT Id FROM User WHERE Email = ?", u.Email).Scan(&id)
	return err == nil
}

func totalMembers() int {
	var count int
	err := DB.QueryRow("SELECT count(*) FROM User").Scan(&count)
	if err != nil {
		log.Fatal(err)
	}
	return count
}

func latestMember() string {
	var name string
	err := DB.QueryRow("SELECT Username FROM User ORDER BY Created DESC LIMIT 1").Scan(&name)
	if err != nil {
		log.Fatal(err)
	}
	return name
}

func postsByUser(userID int64) int {
	var count int
	err := DB.QueryRow("SELECT count(*) FROM Post WHERE User_Id=?", userID).Scan(&count)
	if err != nil {
		log.Fatal(err)
	}
	return count
}

func commentsByUser(userID int64) int {
	var count int
	err := DB.QueryRow("SELECT count(*) FROM Comment WHERE User_Id=?", userID).Scan(&count)
	if err != nil {
		log.Fatal(err)
	}
	return count
}

type UserInfo struct {
	Id       int64  `json:"current_id"`
	Username string `json:"username"`
	Avatar   string `json:"avatar"`
	Days     int    `json:"days"`
	Posts    int    `json:"posts"`
	Comments int    `json:"comments"`
}

func GetUserInfo(uid int64) (UserInfo, error) {
	u := User{Id: uid}
	err := u.Get()
	if err != nil {
		return UserInfo{}, err
	}
	when, err := time.Parse(time.RFC3339, u.Created)
	if err != nil {
		return UserInfo{}, err
	}
	since := time.Since(when)

	nbPosts := postsByUser(u.Id)
	nbComments := commentsByUser(u.Id)
	c := UserInfo{
		Id:       u.Id,
		Username: u.Username,
		Avatar:   u.Avatar,
		Days:     int(since.Hours() / 24),
		Posts:    nbPosts,
		Comments: nbComments - nbPosts,
	}

	//, Since: u.Created, Posts: totalPostsByUser(u.Id), Comments: totalCommentsByUser(u.Id)}
	return c, err
}

func GetTopCoders() []UserInfo {
	rows, err := DB.Query(
		"SELECT User.Id, Username, Avatar, count(Post.Id) FROM User, Post WHERE User.Id = User_Id GROUP BY User.Id ORDER BY count(Post.Id) DESC LIMIT 4",
	)
	if err != nil && err != sql.ErrNoRows {
		log.Fatal(err)
	}
	defer rows.Close()
	var coders []UserInfo
	for rows.Next() {
		u := UserInfo{}
		err := rows.Scan(&u.Id, &u.Username, &u.Avatar, &u.Posts)
		if err != nil {
			log.Fatal(err)
		}
		coders = append(coders, u)
	}
	return coders
}

// --------------------------------------------
type Role struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

func (r *Role) Get() error {
	err := DB.QueryRow("SELECT Id, Name FROM Role WHERE Id=?",
		r.Id).Scan(&r.Id, &r.Name)
	return err
}

// --------------------------------------------
type Post struct {
	Id          int64  `json:"id"`
	Title       string `json:"title"`
	Content_Id  int64  `json:"content_id"`
	User_Id     int64  `json:"user_id"`
	Category_Id int64  `json:"category_id"`
	Created     string `json:"created"`
	Updated     string `json:"updated"`
	Likes       int    `json:"likes"`
	Dislikes    int    `json:"dislikes"`
	Tags        string `json:"tags"`
}

func (p *Post) Create() (int64, error) {
	res, err := DB.Exec("INSERT INTO Post(Title, User_Id, Category_Id, Tags) VALUES(?,?,?,?)",
		p.Title, p.User_Id, p.Category_Id, p.Tags)
	if err != nil {
		return 0, err
	}
	p.Id, err = res.LastInsertId()
	return p.Id, err
}

func (p *Post) Update() error {
	_, err := DB.Exec("UPDATE Post SET Title=?, Content_Id=?, Updated=? WHERE Id=?",
		p.Title, p.Content_Id, p.Updated, p.Id)
	return err
}

func (p *Post) Delete() error {
	res, err := DB.Exec("DELETE FROM Post WHERE Id=?", p.Id)
	if err != nil {
		return err
	}
	var n int64
	n, err = res.RowsAffected()
	if err != nil && n > 0 {
		_, err = DB.Exec("DELETE FROM Comment WHERE Post_Id=?", p.Id)
	}
	return err
}

func (p *Post) Get() error {
	err := DB.QueryRow("SELECT Id, Title, User_Id, Category_Id, Created, Updated, Likes, Dislikes, Tags FROM Post WHERE Id=?",
		p.Id).Scan(&p.Id, &p.Title, &p.User_Id, &p.Category_Id, &p.Created, &p.Updated, &p.Likes, &p.Dislikes, &p.Tags)
	return err
}

// Additional methods
type PostWithAuthor struct {
	Post        Post     `json:"post"`
	Created_Raw string   `json:"created_raw"`
	Author      UserInfo `json:"author"`
}

func GetPostWithAuthor(id int64) (PostWithAuthor, error) {
	post := Post{Id: id}
	err := post.Get()
	if err != nil {
		return PostWithAuthor{}, err
	}
	created_raw := post.Created
	post.Created = FormatDate(post.Created)
	author, err := GetUserInfo(post.User_Id)
	if err != nil {
		return PostWithAuthor{}, err
	}
	return PostWithAuthor{Post: post, Created_Raw: created_raw, Author: author}, nil
}

type PostWithComments struct {
	Post     Post                `json:"post"`
	Comments []CommentWithAuthor `json:"comments"`
}

func GetPostWithComments(id int64) (PostWithComments, error) {
	post := Post{Id: id}
	err := DB.QueryRow("SELECT Id, Title, User_Id, Category_Id, Created, Updated, Likes, Dislikes FROM Post WHERE Id=?",
		id).Scan(&post.Id, &post.Title, &post.User_Id, &post.Category_Id, &post.Created, &post.Updated, &post.Likes, &post.Dislikes)
	if err != nil {
		return PostWithComments{}, err
	}
	var comms []CommentWithAuthor
	comms, err = GetCommentsWithAuthors(id)
	data := PostWithComments{Post: post, Comments: comms}
	return data, err
}

type PostInfo struct {
	Post_Id     int64  `json:"post_id"`
	Title       string `json:"title"`
	Content_Id  int64  `json:"content_id"`
	User_Id     int64  `json:"user_id"`
	Category_Id int64  `json:"category_id"`
	Created     string `json:"created"`
	Created_Raw string `json:"created_raw"`
	Updated     string `json:"updated"`
	Likes       int    `json:"likes"`
	Dislikes    int    `json:"dislikes"`
	Author      string `json:"author"`
	Avatar      string `json:"avatar"`
	Tags        string `json:"tags"`

	CommentsNumber        int    `json:"comments_number"`
	LastCommentatorId     int64  `json:"last_commentator_id"`
	LastCommentatorName   string `json:"last_commentator_name"`
	LastCommentatorAvatar string `json:"last_commentator_avatar"`
	LastCommentCreated    string `json:"last_comment_created"`
	LastCommentCreatedRaw string `json:"last_comment_created_raw"`
	LastCommentId         int64  `json:"last_comment_id"`
}

func GetTopicPosts(topicID int64) ([]PostInfo, error) {
	rows, err := DB.Query(
		"SELECT Post.Id, Title, Content_Id, User_Id, Category_Id, Post.Created, Updated, Username AS Author, Avatar, Tags "+ //Likes, Dislikes,
			"FROM Post, User WHERE Category_Id=? AND User_Id = User.Id", topicID,
	)
	if err != nil { //&& err != sql.ErrNoRows { //.Error() != "sql: no rows in result set" {
		return nil, err
	}
	defer rows.Close()
	var posts []PostInfo
	for rows.Next() {
		p := PostInfo{}
		err := rows.Scan(&p.Post_Id, &p.Title, &p.Content_Id, &p.User_Id, &p.Category_Id, &p.Created, &p.Updated, &p.Author, &p.Avatar, &p.Tags) //&p.Likes, &p.Dislikes,
		if err != nil {
			return nil, err
		}
		p.Created_Raw = p.Created
		p.Created = FormatDate(p.Created)
		p.Likes, p.Dislikes, err = getPostReactions(p.Post_Id)
		if err != nil {
			return nil, err
		}
		err = DB.QueryRow("SELECT count(*) FROM Comment WHERE Post_Id=?", p.Post_Id).Scan(&p.CommentsNumber)
		if err != nil {
			return nil, err
		}
		if p.CommentsNumber > 0 {
			DB.QueryRow("SELECT Comment.Id, User_Id, Username, Avatar, Comment.Created FROM Comment, User WHERE Post_Id=? AND User_Id = User.Id ORDER BY Comment.Created DESC LIMIT 1",
				p.Post_Id).Scan(&p.LastCommentId, &p.LastCommentatorId, &p.LastCommentatorName, &p.LastCommentatorAvatar, &p.LastCommentCreated)
			p.LastCommentCreatedRaw = p.LastCommentCreated
			p.LastCommentCreated = FormatDate(p.LastCommentCreated)
		}
		posts = append(posts, p)
	}
	return posts, nil
}

func GetAllPosts() []PostInfo {
	rows, err := DB.Query(
		"SELECT Post.Id, Title, User_Id, Category_Id, Post.Created, Updated, Username AS Author, Avatar, Tags " + //Likes, Dislikes,
			"FROM Post, User WHERE User_Id = User.Id",
	)
	if err != nil && err != sql.ErrNoRows { //}"sql: no rows in result set" {
		log.Fatal(err)
	}
	defer rows.Close()
	var posts []PostInfo
	for rows.Next() {
		p := PostInfo{}
		err := rows.Scan(&p.Post_Id, &p.Title, &p.User_Id, &p.Category_Id, &p.Created, &p.Updated, &p.Author, &p.Avatar, &p.Tags) //&p.Likes, &p.Dislikes,
		if err != nil {
			log.Fatal(err)
		}
		p.Created_Raw = p.Created
		p.Created = FormatDate(p.Created)
		p.Likes, p.Dislikes, err = getPostReactions(p.Post_Id)
		if err != nil {
			log.Fatal(err)
		}
		DB.QueryRow("SELECT count(*) FROM Comment WHERE Post_Id=?", p.Post_Id).Scan(&p.CommentsNumber)
		if p.CommentsNumber > 0 {
			DB.QueryRow("SELECT Comment.Id, User_Id, Username, Avatar, Comment.Created FROM Comment, User WHERE Post_Id=? AND User_Id = User.Id ORDER BY Comment.Created DESC LIMIT 1",
				p.Post_Id).Scan(&p.LastCommentId, &p.LastCommentatorId, &p.LastCommentatorName, &p.LastCommentatorAvatar, &p.LastCommentCreated)
			p.LastCommentCreatedRaw = p.LastCommentCreated
			p.LastCommentCreated = FormatDate(p.LastCommentCreated)
		}
		posts = append(posts, p)
	}
	return posts
}

func totalPosts() int {
	var count int
	err := DB.QueryRow("SELECT count(*) FROM Post").Scan(&count)
	if err != nil {
		log.Fatal(err)
	}
	return count
}

func GetLatestPosts() []PostWithAuthor {
	rows, err := DB.Query(
		"SELECT Id FROM Post ORDER BY Created DESC LIMIT 5",
	)
	if err != nil && err != sql.ErrNoRows {
		log.Fatal(err)
	}
	defer rows.Close()
	var posts []PostWithAuthor
	for rows.Next() {
		var id int64
		err := rows.Scan(&id)
		if err != nil {
			log.Fatal(err)
		}
		p, err := GetPostWithAuthor(id)
		if err != nil {
			log.Fatal(err)
		}
		posts = append(posts, p)
	}
	return posts
}

func GetHotTopics() []PostWithAuthor {
	rows, err := DB.Query(
		"SELECT Id FROM Post ORDER BY Likes+Dislikes DESC LIMIT 5",
	)
	if err != nil && err != sql.ErrNoRows {
		log.Fatal(err)
	}
	defer rows.Close()
	var posts []PostWithAuthor
	for rows.Next() {
		var id int64
		err := rows.Scan(&id)
		if err != nil {
			log.Fatal(err)
		}
		p, err := GetPostWithAuthor(id)
		if err != nil {
			log.Fatal(err)
		}
		posts = append(posts, p)
	}
	return posts
}

func GetPostsByUser(id int64) ([]PostInfo, error) {
	rows, err := DB.Query(
		"SELECT Post.Id, Title, User_Id, Category_Id, Post.Created, Updated, Username AS Author, Avatar, Tags "+ //Likes, Dislikes,
			"FROM Post, User WHERE User_Id = User.Id AND User_Id=?", id,
	)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	defer rows.Close()
	var posts []PostInfo
	for rows.Next() {
		p := PostInfo{}
		err := rows.Scan(&p.Post_Id, &p.Title, &p.User_Id, &p.Category_Id, &p.Created, &p.Updated, &p.Author, &p.Avatar, &p.Tags) //&p.Likes, &p.Dislikes,
		if err != nil {
			return nil, err
		}
		p.Created_Raw = p.Created
		p.Created = FormatDate(p.Created)
		p.Likes, p.Dislikes, err = getPostReactions(p.Post_Id)
		if err != nil {
			return nil, err
		}
		err = DB.QueryRow("SELECT count(*) FROM Comment WHERE Post_Id=?", p.Post_Id).Scan(&p.CommentsNumber)
		if err != nil {
			return nil, err
		}
		if p.CommentsNumber > 0 {
			DB.QueryRow("SELECT Comment.Id, User_Id, Username, Avatar, Comment.Created FROM Comment, User WHERE Post_Id=? AND User_Id = User.Id ORDER BY Comment.Created DESC LIMIT 1",
				p.Post_Id).Scan(&p.LastCommentId, &p.LastCommentatorId, &p.LastCommentatorName, &p.LastCommentatorAvatar, &p.LastCommentCreated)
			p.LastCommentCreatedRaw = p.LastCommentCreated
			p.LastCommentCreated = FormatDate(p.LastCommentCreated)
		}
		posts = append(posts, p)
	}
	return posts, nil
}

func GetLikedPostsByUser(id int64) ([]PostInfo, error) {
	rows, err := DB.Query(
		"SELECT Post.Id, Title, Content_Id, Post.User_Id, Category_Id, Post.Created, Post.Updated, Username AS Author, Avatar, Tags "+ //Likes, Dislikes,
			"FROM Comment_Reaction, Post, User WHERE Comment_Reaction.User_Id=? AND Comment_Id=Content_Id "+
			"AND Reaction=1 AND Post.User_Id = User.Id", id,
	)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	defer rows.Close()
	var posts []PostInfo
	for rows.Next() {
		p := PostInfo{}
		err := rows.Scan(&p.Post_Id, &p.Title, &p.Content_Id, &p.User_Id, &p.Category_Id, &p.Created, &p.Updated, &p.Author, &p.Avatar, &p.Tags) //&p.Likes, &p.Dislikes,
		if err != nil {
			return nil, err
		}
		p.Created_Raw = p.Created
		p.Created = FormatDate(p.Created)
		p.Likes, p.Dislikes, err = getPostReactions(p.Post_Id)
		if err != nil {
			return nil, err
		}
		err = DB.QueryRow("SELECT count(*) FROM Comment WHERE Post_Id=?", p.Post_Id).Scan(&p.CommentsNumber)
		if err != nil {
			return nil, err
		}
		if p.CommentsNumber > 0 {
			DB.QueryRow("SELECT Comment.Id, User_Id, Username, Avatar, Comment.Created FROM Comment, User WHERE Post_Id=? AND User_Id = User.Id ORDER BY Comment.Created DESC LIMIT 1",
				p.Post_Id).Scan(&p.LastCommentId, &p.LastCommentatorId, &p.LastCommentatorName, &p.LastCommentatorAvatar, &p.LastCommentCreated)
			p.LastCommentCreatedRaw = p.LastCommentCreated
			p.LastCommentCreated = FormatDate(p.LastCommentCreated)
		}
		posts = append(posts, p)
	}
	return posts, nil
}

// --------------------------------------------
type Comment struct {
	Id       int64  `json:"id"`
	User_Id  int64  `json:"user_id"`
	Post_Id  int64  `json:"post_id"`
	Content  string `json:"content"`
	Is_Post  bool   `json:"is_post"`
	Created  string `json:"created"`
	Updated  string `json:"updated"`
	Likes    int    `json:"likes"`
	Dislikes int    `json:"dislikes"`
}

func (c *Comment) Create() (int64, error) {
	res, err := DB.Exec("INSERT INTO Comment(User_Id, Post_Id, Content, Is_Post) VALUES(?,?,?,?)",
		c.User_Id, c.Post_Id, c.Content, c.Is_Post)
	if err != nil {
		return 0, err
	}
	c.Id, err = res.LastInsertId()
	return c.Id, err
}

func (c *Comment) Update() error {
	_, err := DB.Exec("UPDATE Comment SET Content=?, Updated=? WHERE Id=?",
		c.Content, c.Updated, c.Id)
	return err
}

func (c *Comment) Delete() error {
	_, err := DB.Exec("DELETE FROM Comment WHERE Id=?", c.Id)
	return err
}

func (c *Comment) Get() error {
	err := DB.QueryRow("SELECT Id, User_Id, Post_Id, Content, Created, Updated, Likes, Dislikes FROM Comment WHERE Id=?",
		c.Id).Scan(&c.Id, &c.User_Id, &c.Post_Id, &c.Content, &c.Created, &c.Updated, &c.Likes, &c.Dislikes)
	return err
}

// Additional methods
type CommentWithAuthor struct {
	Comment_Id  int64  `json:"comment_id"`
	Post_Id     int64  `json:"post_id"`
	Post_Title  string `json:"post_title"`
	Content     string `json:"content"`
	Created     string `json:"created"`
	Created_Raw string `json:"created_raw"`
	Updated     string `json:"updated"`
	Likes       int    `json:"likes"`
	Dislikes    int    `json:"dislikes"`

	User_Id  int64  `json:"user_id"`
	Username string `json:"author"`
	Avatar   string `json:"avatar"`
	Days     int    `json:"member_days"`
	Posts    int    `json:"posts_number"`
	Comments int    `json:"comments_number"`
}

func GetCommentsWithAuthors(postID int64) ([]CommentWithAuthor, error) {
	rows, err := DB.Query("SELECT Comment.Id, Post_Id, Post.Title, Comment.Content, Comment.Created, Comment.Updated, Comment.User_Id FROM Comment, Post WHERE Post.Id=Post_Id AND Post_Id=?", postID) //Comment.Likes, Comment.Dislikes,
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	defer rows.Close()
	var comms []CommentWithAuthor
	for rows.Next() {
		c := CommentWithAuthor{}
		err := rows.Scan(&c.Comment_Id, &c.Post_Id, &c.Post_Title, &c.Content, &c.Created, &c.Updated, &c.User_Id) //&c.Likes, &c.Dislikes,
		if err != nil {
			return nil, err
		}
		c.Created_Raw = c.Created
		c.Created = FormatDate(c.Created)
		c.Likes, c.Dislikes, err = getCommentReactions(c.Comment_Id)
		if err != nil {
			return nil, err
		}
		var u UserInfo
		u, err = GetUserInfo(c.User_Id)
		if err != nil {
			return nil, err
		}
		c.Username = u.Username
		c.Avatar = u.Avatar
		c.Days = u.Days
		c.Posts = u.Posts
		c.Comments = u.Comments

		comms = append(comms, c)
	}
	return comms, nil
}

func totalCommentsByCategory(categoryID int64) (int, error) {
	var count int
	err := DB.QueryRow("SELECT count(*) FROM Comment, Post WHERE Post_Id = Post.Id AND Category_Id=?", categoryID).Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func GetLastCommentByCategory(categoryID int64) (CommentWithAuthor, error) {
	c := CommentWithAuthor{}
	//err := DB.QueryRow("SELECT Comment.Id FROM Comment,Post WHERE Post_Id = Post.Id AND Category_Id=? ORDER BY Comment.Created DESC LIMIT 1",)
	err := DB.QueryRow("SELECT Comment.Id, Post_Id, Post.Title, Comment.Created, Comment.User_Id, Username, Avatar "+
		"FROM Comment, Post, User WHERE Post_Id = Post.Id AND Comment.User_Id = User.Id AND Category_Id=? ORDER BY Comment.Created DESC LIMIT 1",
		categoryID).Scan(&c.Comment_Id, &c.Post_Id, &c.Post_Title, &c.Created, &c.User_Id, &c.Username, &c.Avatar)
	if err != nil {
		return c, err
	}
	//if categoryID == 9 {
	//	fmt.Println(c)
	//}
	c.Created_Raw = c.Created
	c.Created = FormatDate(c.Created)
	return c, err
}

// --------------------------------------------
type Category struct {
	Id        int64  `json:"id"`
	Parent_Id int64  `json:"parent_id"`
	Short     string `json:"short"`
	Title     string `json:"title"`
	Intro     string `json:"intro"`
}

func (c *Category) Create() (int64, error) {
	res, err := DB.Exec("INSERT INTO Category(Parent_Id, Short, Title, Intro) VALUES(?,?,?,?)",
		c.Parent_Id, c.Short, c.Title, c.Intro)
	if err != nil {
		return 0, err
	}
	c.Id, err = res.LastInsertId()
	return c.Id, err
}

func (c *Category) Update() error {
	_, err := DB.Exec("UPDATE Category SET Parent_Id=?, Short=?, Title=?, Intro=? WHERE Id=?",
		c.Parent_Id, c.Short, c.Title, c.Intro, c.Id)
	return err
}

func (c *Category) Delete() error {
	_, err := DB.Exec("DELETE FROM Category WHERE Id=?", c.Id)
	return err
}

func (c *Category) Get() error {
	err := DB.QueryRow("SELECT Id, Parent_Id, Short, Title, Intro FROM Category WHERE Id=?",
		c.Id).Scan(&c.Id, &c.Parent_Id, &c.Short, &c.Title, &c.Intro)
	return err
}

// Additional methods
func GetCategories(id int64) ([]Category, error) {
	rows, err := DB.Query("SELECT Id, Parent_Id, Short, Title, Intro FROM Category WHERE Parent_Id=?", id)
	if err != nil && err.Error() != "sql: no rows in result set" {
		return nil, err
	}
	defer rows.Close()
	var cats []Category
	for rows.Next() {
		c := Category{}
		err := rows.Scan(&c.Id, &c.Parent_Id, &c.Short, &c.Title, &c.Intro)
		if err != nil {
			return nil, err
		}
		cats = append(cats, c)
	}
	return cats, err
}

type CategoryWithParent struct {
	Category Category `json:"category"`
	Parent   Category `json:"parent"`
}

func GetCategoryWithParent(id int64) (CategoryWithParent, error) {
	c := CategoryWithParent{}
	err := DB.QueryRow("SELECT Id, Parent_Id, Short, Title, Intro FROM Category WHERE Id=?",
		id).Scan(&c.Category.Id, &c.Category.Parent_Id, &c.Category.Short, &c.Category.Title, &c.Category.Intro)
	if err != nil {
		return c, err
	}
	err = DB.QueryRow("SELECT Id, Parent_Id, Short, Title, Intro FROM Category WHERE Id=?",
		c.Category.Parent_Id).Scan(&c.Parent.Id, &c.Parent.Parent_Id, &c.Parent.Short, &c.Parent.Title, &c.Parent.Intro)
	return c, err
}

type CategoryWithCommentsInfo struct {
	Category       Category          `json:"category"`
	CommentsNumber int               `json:"comments_number"`
	LastComment    CommentWithAuthor `json:"last_comment"`
}

func GetChildrenCategories(parentID int64) ([]CategoryWithCommentsInfo, error) {
	rows, err := DB.Query("SELECT Id, Parent_Id, Short, Title, Intro FROM Category WHERE Parent_Id=?", parentID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var subCategories []CategoryWithCommentsInfo
	for rows.Next() {
		cwci := CategoryWithCommentsInfo{}
		c := Category{}
		err := rows.Scan(&c.Id, &c.Parent_Id, &c.Short, &c.Title, &c.Intro)
		if err != nil {
			return nil, err
		}
		cwci.Category = c
		cwci.CommentsNumber, err = totalCommentsByCategory(c.Id)
		if err != nil {
			return nil, err
		}
		if cwci.CommentsNumber > 0 {
			cwci.LastComment, err = GetLastCommentByCategory(c.Id)
			if err != nil {
				return nil, err
			}
		}
		subCategories = append(subCategories, cwci)
	}
	return subCategories, nil
}

type CategoryWithChildren struct {
	Category Category                   `json:"category"`
	Topics   []CategoryWithCommentsInfo `json:"topics"`
}

func GetCategoryStructure() []CategoryWithChildren {
	var data []CategoryWithChildren
	cat0, err := GetCategories(0)
	if err != nil {
		log.Fatal(err)
	}
	for _, c0 := range cat0 {
		cat1, err := GetChildrenCategories(c0.Id)
		if err != nil {
			log.Fatal(err)
		}
		data = append(data, CategoryWithChildren{c0, cat1})
	}
	return data
}

func totalTopics() int {
	var count int
	err := DB.QueryRow("SELECT count(*) FROM Category WHERE Parent_Id>0").Scan(&count)
	if err != nil {
		log.Fatal(err)
	}
	return count
}

// --------------------------------------------
type Tag struct {
	Id   int64  `json:"id"`
	Name string `json:"tag"`
}

func (t *Tag) Create() (int64, error) {
	res, err := DB.Exec("INSERT INTO Tag(Name) VALUES(?)",
		t.Name)
	if err != nil {
		return 0, err
	}
	t.Id, err = res.LastInsertId()
	return t.Id, err
}

func (t *Tag) Update() error {
	_, err := DB.Exec("UPDATE Tag SET Name=? WHERE Id=?",
		t.Name, t.Id)
	return err
}

func (t *Tag) Delete() error {
	_, err := DB.Exec("DELETE FROM Tag WHERE Id=?", t.Id)
	return err
}

func (t *Tag) Get() error {
	err := DB.QueryRow("SELECT Id, Name FROM Tag WHERE Id=?",
		t.Id).Scan(&t.Id, &t.Name)
	return err
}

// --------------------------------------------
type Post_Tag struct {
	Id      int64 `json:"id"`
	Post_Id int64 `json:"post_id"`
	Tag_Id  int64 `json:"tag_id"`
}

// --------------------------------------------
type Stats struct {
	TotalOnline  int    `json:"total_online"`
	TotalMembers int    `json:"total_members"`
	TotalTopics  int    `json:"total_topics"`
	TotalPosts   int    `json:"total_posts"`
	LatestMember string `json:"latest_member"`
}

func GetStats() Stats {
	stats := Stats{
		TotalOnline:  totalOnline(),
		TotalMembers: totalMembers(),
		TotalTopics:  totalTopics(),
		TotalPosts:   totalPosts(),
		LatestMember: latestMember(),
	}
	return stats
}

// --------------------------------------------
type PostReaction struct {
	Id       int64 `json:"id"`
	User_Id  int64 `json:"user_id"`
	Post_Id  int64 `json:"post_id"`
	Reaction int   `json:"reaction"`
}

func (pr *PostReaction) Create() (int64, error) {
	res, err := DB.Exec("INSERT INTO Post_Reaction(User_Id, Post_Id, Reaction) VALUES(?,?,?)",
		pr.User_Id, pr.Post_Id, pr.Reaction)
	if err != nil {
		return 0, err
	}
	pr.Id, err = res.LastInsertId()
	return pr.Id, err
}

func (pr *PostReaction) Update() error {
	_, err := DB.Exec("UPDATE Post_Reaction SET Reaction=? WHERE User_Id=? AND Post_Id=?",
		pr.Reaction, pr.User_Id, pr.Post_Id)
	return err
}

func (pr *PostReaction) Delete() error {
	_, err := DB.Exec("DELETE FROM Post_Reaction WHERE User_Id=? AND Post_Id=?", pr.User_Id, pr.Post_Id)
	return err
}

func (pr *PostReaction) Get() error {
	err := DB.QueryRow("SELECT Id, User_Id, Post_Id, Reaction FROM Post_Reaction WHERE User_Id=? AND Post_Id=?",
		pr.User_Id, pr.Post_Id).Scan(&pr.Id, &pr.User_Id, &pr.Post_Id, &pr.Reaction)
	return err
}

func getPostReactions(postID int64) (int, int, error) {
	var likes, dislikes int
	err := DB.QueryRow("SELECT count(*) FROM Comment, Comment_Reaction WHERE Post_Id=? AND Comment.Id=Comment_Id AND Is_Post=1 AND Reaction=1", postID).Scan(&likes)
	if err != nil {
		return 0, 0, err
	}
	err = DB.QueryRow("SELECT count(*) FROM Comment, Comment_Reaction WHERE Post_Id=? AND Comment.Id=Comment_Id AND Is_Post=1 AND Reaction=-1", postID).Scan(&dislikes)
	if err != nil {
		return 0, 0, err
	}
	return likes, dislikes, nil
}

func deletePostReaction(rid int64) error {
	_, err := DB.Exec("DELETE FROM Comment_Reaction WHERE Id=?", rid)
	return err
}

// --------------------------------------------
type CommentReaction struct {
	Id         int64 `json:"id"`
	User_Id    int64 `json:"user_id"`
	Comment_Id int64 `json:"comment_id"`
	Reaction   int   `json:"reaction"`
}

func (cr *CommentReaction) Create() (int64, error) {
	res, err := DB.Exec("INSERT INTO Comment_Reaction(User_Id, Comment_Id, Reaction) VALUES(?,?,?)",
		cr.User_Id, cr.Comment_Id, cr.Reaction)
	if err != nil {
		return 0, err
	}
	cr.Id, err = res.LastInsertId()
	return cr.Id, err
}

func (cr *CommentReaction) Update() error {
	_, err := DB.Exec("UPDATE Comment_Reaction SET Reaction=? WHERE User_Id=? AND Comment_Id=?",
		cr.Reaction, cr.User_Id, cr.Comment_Id)
	return err
}

func (cr *CommentReaction) Delete() error {
	_, err := DB.Exec("DELETE FROM Comment_Reaction WHERE User_Id=? AND Comment_Id=?", cr.User_Id, cr.Comment_Id)
	return err
}

func (cr *CommentReaction) Get() error {
	err := DB.QueryRow("SELECT Id, User_Id, Comment_Id, Reaction FROM Comment_Reaction WHERE User_Id=? AND Comment_Id=?",
		cr.User_Id, cr.Comment_Id).Scan(&cr.Id, &cr.User_Id, &cr.Comment_Id, &cr.Reaction)
	return err
}

func getCommentReactions(commentID int64) (int, int, error) {
	var likes, dislikes int
	err := DB.QueryRow("SELECT count(*) FROM Comment_Reaction WHERE Comment_Id=? AND Reaction=1", commentID).Scan(&likes)
	if err != nil {
		return 0, 0, err
	}
	err = DB.QueryRow("SELECT count(*) FROM Comment_Reaction WHERE Comment_Id=? AND Reaction=-1", commentID).Scan(&dislikes)
	if err != nil {
		return 0, 0, err
	}
	return likes, dislikes, nil
}

func deleteCommentReaction(commID int64) error {
	_, err := DB.Exec("DELETE FROM Comment_Reaction WHERE Id=?", commID)
	return err
}
