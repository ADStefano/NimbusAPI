package routes

import (
	"net/http"

	"github.com/ADStefano/NimbusAPI/internal/api/controllers"
)

func CreateBucketRoutes(ctrl controllers.BucketController) []Route {
	return []Route{
	{
		Name:    "CreateBucket",
		Method:  http.MethodPost,
		Path:    "/buckets/:name",
		Handler: ctrl.CreateBucket,
		Auth:    true,
	},
	{
		Name:   "ListBuckets",
		Method: http.MethodGet,
		Path:   "/buckets",
		// Handler: ListBuckets,
		Auth: true,
	},
	{
		Name:   "DeleteBucket",
		Method: http.MethodDelete,
		Path:   "/buckets/:name",
		// Handler: DeleteBucket,
		Auth: true,
	},		
	}
}
