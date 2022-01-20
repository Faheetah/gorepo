You create your own repo struct via function, this becomes a *Repo object from the method return. The values could come from code or environment vars or config handling or whatever.

```go
db, _ := sql.Open("sqlite3", "./example.db")
repo := gorepo.Start(db)

var User struct{
  Id int,
  Name string,
  Email string,
}

var Post struct{
  Id int,
  Title string,
  Body string,
  Author &User,
}

repo.Table(&User{}).Create()

bob := &User{Name: "Bob", Email: "bob@example.com"}
repo.Insert(&bob)

rob := &User{Name: "Rob", Email: "rob@example.com"}
repo.Insert(&rob)

repo.Table(&Post{}).Create()

initialPost := &Post{
  Title: "Hello",
  Body: "I am Bob, not Rob",
  Author: &bob,
}

repo.Insert(&initializePost)

// The `posts` var gets updated in place
var posts []Post{}
repo.All(&posts)
id := posts[0].id

// The `post` var gets updated in place
var post Post{}
repo.Get(&post, id)

// Update the post struct, only assigned fiels will get updated
// Fields that are the same are handled through the DB's idempotence
post.Body = "I'm the Bob now"
post.Author = rob
repo.Update(&post)

// Now Bob is angry
repo.Delete(&post)

var authors []Post{}
var query Post{
  Authors: authors,
}
// Prepare a query and execute it
repo.Query(&query).Select([]str{"id", "title"}).OrderBy(gorepo.ASC).Limit(5).Preload(&authors)
```
