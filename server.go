package main

import (
	"fmt"
	"github.com/kartikeya/sample_app/controller"
	"github.com/kartikeya/sample_app/repository"
	"github.com/kartikeya/sample_app/router"
	"github.com/kartikeya/sample_app/service"
	"gorm.io/gorm"
	"net/http"
)

var (
	httpRouter     router.Router             = router.NewMuxRouter()
	DB             *gorm.DB                  = repository.ConnectDatabase()
	PostRepository repository.PostRepository = repository.NewPostgresRepository(DB)
	postService    service.PostService       = service.NewPostService(PostRepository)
	postController controller.PostController = controller.NewPostController(postService)
)

func main() {
	repository.ConnectDatabase()

	const port string = ":8080"
	httpRouter.GET("/", func(resp http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(resp, "UP and Running.....")
	})
	httpRouter.GET("/posts", postController.GetPosts)
	httpRouter.POST("/posts", postController.AddPost)
	httpRouter.SERVE(port)
}
