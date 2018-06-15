package method

type ForumPost struct {
	ThreadID     int    `json:"thread_id"`      //Unique identifier of the thread
	PostID       int    `json:"post_id"`        //Id of the individual post
	Post         string `json:"post"`           //The content of the post
	UserID       string `json:"user_id"`        //Id of the user who created the post
	LastEditTime int    `json:"last_edit_time"` //When the post was last edited
	CreateDate   int    `json:"create_date"`    //When the post was created
}
