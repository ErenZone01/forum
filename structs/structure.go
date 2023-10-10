package structs

type Users struct {
	Id       int
	Username string
	Mdp      string
	Email    string
	Role     string
	Error    string
}

type Error struct {
	Title string
	Body  string
}

type Posts struct {
	Id        int
	Title     string
	Body      string
	N_like    int
	N_dislike int
	N_com     int
	Users_id  int
	Filter    int

	CreatedPost string
}

type All_Forum struct {
	Users            []Users
	Posts            []Posts
	Coms             []Comments
	Mylike           []Mylike
	MyPost           []Posts
	FilterLiked      []Posts
	MyCom            []Comments
	Utilisateur      []Users
	FilterCategories []Posts
}

type Mylike struct {
	N_like int
}

type Comments struct {
	Id        int
	Body      string
	N_like    int
	N_dislike int
	Users_id  int
	Posts_id  int
}

type Appreciation_com struct {
	Id       int
	Like     string
	Dislike  string
	Users_id int
	Coms_id  int
}

type Appreciation_post struct {
	Id       int
	Like     string
	Dislike  string
	Users_id int
	Posts_id int
}

type Categories struct {
	Id       int
	types    string
	Posts_id int
}
