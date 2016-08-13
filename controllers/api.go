package controllers

import "github.com/kataras/iris"

func init() {
	//iris.API("/", UserAPI{})
}

type UserAPI struct {
	*iris.Context
}

// GET /users
func (u UserAPI) Get() {
	u.Write("Get from /users")
	// u.JSON(iris.StatusOK,myDb.AllUsers())
}

// GET /users/:param1 which its value passed to the id argument
func (u UserAPI) GetBy(id string) { // id equals to u.Param("param1")
	u.Write("Get from /users/%s", id)
	// u.JSON(iris.StatusOK, myDb.GetUserById(id))

}

// POST /users
func (u UserAPI) Post() {
	name := u.FormValue("name")
	// myDb.InsertUser(...)
	println(string(name))
	println("Post from /users")
}

// PUT /users/:param1
func (u UserAPI) PutBy(id string) {
	name := u.FormValue("name") // you can still use the whole Context's features!
	// myDb.UpdateUser(...)
	println(string(name))
	println("Put from /users/" + id)
}

// DELETE /users/:param1
func (u UserAPI) DeleteBy(id string) {
	// myDb.DeleteUser(id)
	println("Delete from /" + id)
}
