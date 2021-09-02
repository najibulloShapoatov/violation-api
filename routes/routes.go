package routes

import (
	"mobile-rest-api/app"
	"mobile-rest-api/controllers"
	"mobile-rest-api/utils"
	"net/http"
	"github.com/gorilla/mux"
)

//Init routes
func Init() *mux.Router {

	var router = mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/api/tarrifs", controllers.GetListTarris).Methods("GET")
	router.HandleFunc("/api/customer/checkAuth", controllers.CheckAuthCustomer)
	router.HandleFunc("/api/customer", controllers.GetCustomer)
	router.HandleFunc("/api/customer/sign-up", controllers.CreateCustomer)
	router.HandleFunc("/api/customer/validate-sms", controllers.CustomerCheckSMS)
	router.HandleFunc("/api/customer/subscribtions", controllers.CustomerSubscribtionsAll)
	router.HandleFunc("/api/customer/payments", controllers.CustomerPayments)
	router.HandleFunc("/api/customer/subscribe", controllers.CustomerSubscribe)
	router.HandleFunc("/api/customer/violation", controllers.CustomerGetAllViolations)
	router.HandleFunc("/api/customer/violation/image", controllers.CustomerImages)
	
	router.HandleFunc("/api/customer/update", controllers.CustomerUpdate).Methods("POST")

	router.HandleFunc("/api/events", controllers.GetEvents)

	router.HandleFunc("/api/histories", controllers.GetHistories)

	router.HandleFunc("/api/customer/subscribtion/auto-prolonged", controllers.SubscribeChangeAutoprolonged)

	router.HandleFunc("/api/customer/insurance/create", controllers.CreateInsurance)
	router.HandleFunc("/api/customer/insurances", controllers.GetInsurance)

	router.HandleFunc("/api/customer/pay-services", controllers.GetListPayServices)

	router.HandleFunc("/api/page", controllers.GetPageBySlug)
	router.HandleFunc("/api/pages", controllers.GetListPages)

	router.HandleFunc("/api/faqs", controllers.GetListActiveFaq)

	router.HandleFunc("/api/banners", controllers.GetListBanners)
	router.HandleFunc("/api/banner-by-id", controllers.GetBannerByID)
	router.HandleFunc("/api/banner-by-type", controllers.GetBannerByType)


	router.HandleFunc("/api/articles", controllers.GetListArticles)
	router.HandleFunc("/api/article", controllers.GetArticleByID)

	router.HandleFunc("/api/feedback/create", controllers.CreateFeedBack)
	router.HandleFunc("/api/feedbacks", controllers.GetFeedBack)

	

	// Choose the folder to serve
	staticDir := "/public/"

	// Create the route
	router.
		PathPrefix(staticDir).
		Handler(http.StripPrefix(staticDir, http.FileServer(http.Dir("."+staticDir))))

	app.NotAuthURls = append(app.NotAuthURls,

		//"/api/user/login",
		"/upload",
		//real
		"/api/customer/sign-up",
		"/api/customer/validate-sms",
	)

	router.Use(app.JWTAutentication) //attach JWT auth middleware

	router.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		var res = utils.Message(707, "Эти ресурсы не были найдены на нашем сервере")

		utils.Respond(w, res)
	})

	return router
}
