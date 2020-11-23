package app

import (
	"github.com/cspinetta/go-tracing/src/controllers"
	"github.com/cspinetta/go-tracing/src/repository"
	"github.com/cspinetta/go-tracing/src/services"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"go.opentelemetry.io/otel/exporters/trace/jaeger"
	"go.opentelemetry.io/otel/label"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"log"
)

type App struct {
	Router      *gin.Engine
	UserService services.IUserService
}

func NewApp(db *sqlx.DB) *App {

	flush := jaegerTracer()
	defer flush()

	userRepository := repository.NewUserRepository(db)
	userService := services.NewUserService(userRepository)

	app := App{
		Router:      gin.New(),
		UserService: userService,
	}

	app.RegisterMiddlewares()
	app.RegisterControllers()

	return &app
}

func (app *App) Start(port string) {

	log.Printf("Starting application on port %v", port)

	if err := app.Router.Run(":" + port); err != nil {
		panic("Error running server")
	}
}

func (app *App) RegisterControllers() {
	userController := controllers.NewUserController(app.UserService)

	app.Router.GET("/ping", controllers.NewHealthCheckHandler().HandlePing)

	dataFields := app.Router.RouterGroup.Group("/user")
	{
		dataFields.POST("/", userController.SaveUserInfo)
		dataFields.GET("/", userController.ListUser)
		dataFields.GET("/:user-id", userController.GetUserInfo)
	}
}

func (app *App) RegisterMiddlewares() {
	app.Router.Use(otelgin.Middleware("go-tracing-service"))
}

func jaegerTracer() func() {
	flush, err := jaeger.InstallNewPipeline(
		jaeger.WithCollectorEndpoint("http://jaeger:14268/api/traces"),
		jaeger.WithProcess(jaeger.Process{
			ServiceName: "go-tracing-service",
			Tags: []label.KeyValue{
				label.String("exporter", "jaeger"),
				label.Float64("float", 312.23),
			},
		}),
		jaeger.WithSDK(&sdktrace.Config{DefaultSampler: sdktrace.AlwaysSample()}),
	)
	if err != nil {
		log.Fatal(err)
	}
	return flush
}

//func initStdOutputTracer() {
//	exporter, err := stdout.NewExporter(stdout.WithPrettyPrint())
//	if err != nil {
//		log.Fatal(err)
//	}
//	cfg := sdktrace.Config{
//		DefaultSampler: sdktrace.AlwaysSample(),
//	}
//	tp := sdktrace.NewTracerProvider(
//		sdktrace.WithConfig(cfg),
//		sdktrace.WithSyncer(exporter),
//	)
//	if err != nil {
//		log.Fatal(err)
//	}
//	otel.SetTracerProvider(tp)
//	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))
//}
