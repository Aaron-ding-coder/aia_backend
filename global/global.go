package global

import (
	"aia_backend/handler"
	"aia_backend/models"
	"context"
	"fmt"
)

func MustSetup(ctx context.Context) {
	fmt.Println("start mongo init")
	models.Init(ctx)

	fmt.Println("start object init")
	handler.InitObjectStorage()

	fmt.Println("init done")
}
