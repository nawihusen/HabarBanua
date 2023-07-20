package main

// import (
// 	"be-service-saksi-management/config"
// 	"context"
// 	"net"
// 	"strconv"

// 	"time"

// 	_RepoGRPCAuthSaksi "be-service-saksi-management/saksi/delivery/grpc"
// 	_RepoGRPCSaksiAuth "be-service-saksi-management/saksi/delivery/grpc/auth"
// 	_DeliveryHTTPAccount "be-service-saksi-management/saksi/delivery/http"
// 	_RepoGRPCSaksi "be-service-saksi-management/saksi/repository/grpc"
// 	_MotionPayRepo "be-service-saksi-management/saksi/repository/http/motionpay"
// 	_RepoMySQLAccount "be-service-saksi-management/saksi/repository/mysql"
// 	_RepoRedisAccount "be-service-saksi-management/saksi/repository/redis"
// 	_UsecaseAccount "be-service-saksi-management/saksi/usecase"

// 	"database/sql"
// 	"flag"
// 	"fmt"
// 	"net/http"
// 	"net/url"

// 	"github.com/go-redis/redis/v8"
// 	_ "github.com/go-sql-driver/mysql"
// 	"github.com/golang-migrate/migrate/v4"
// 	"github.com/golang-migrate/migrate/v4/database/mysql"
// 	_ "github.com/golang-migrate/migrate/v4/source/file"
// 	grpcpool "github.com/processout/grpc-go-pool"
// 	"google.golang.org/grpc"
// 	"google.golang.org/grpc/credentials/insecure"

// 	log "github.com/sirupsen/logrus"
// 	"github.com/spf13/viper"

// 	"github.com/gofiber/fiber/v2"
// 	"github.com/gofiber/fiber/v2/middleware/cors"
// 	"github.com/gofiber/fiber/v2/middleware/logger"
// 	"github.com/gofiber/fiber/v2/middleware/recover"
// )

// func main() {
// 	// CLI options parse
// 	configFile := flag.String("c", "config.yaml", "Config file")
// 	flag.Parse()

// 	// Config file
// 	config.ReadConfig(*configFile)

// 	// Set log level
// 	switch viper.GetString("server.log_level") {
// 	case "error":
// 		log.SetLevel(log.ErrorLevel)
// 	case "warning":
// 		log.SetLevel(log.WarnLevel)
// 	case "info":
// 		log.SetLevel(log.InfoLevel)
// 	case "debug":
// 		log.SetLevel(log.DebugLevel)
// 	default:
// 		log.SetLevel(log.InfoLevel)
// 	}

// 	// Initialize database
// 	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", viper.GetString("mysql.user"), viper.GetString("mysql.password"), viper.GetString("mysql.host"), viper.GetString("mysql.port"), viper.GetString("mysql.database"))
// 	val := url.Values{}
// 	val.Add("multiStatements", "true")
// 	val.Add("parseTime", "1")
// 	val.Add("loc", "Asia/Jakarta")
// 	dsn := fmt.Sprintf("%s?%s", connection, val.Encode())
// 	dbConn, err := sql.Open(`mysql`, dsn)

// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	err = dbConn.Ping()
// 	if err != nil {
// 		fmt.Printf("%+v\n", err)
// 		log.Fatal(err)
// 	}

// 	defer func() {
// 		err := dbConn.Close()
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 	}()

// 	// Migrate database if any new schema
// 	driver, err := mysql.WithInstance(dbConn, &mysql.Config{})
// 	if err == nil {
// 		mig, err := migrate.NewWithDatabaseInstance(viper.GetString("mysql.path_migrate"), viper.GetString("mysql.database"), driver)
// 		log.Info(viper.GetString("mysql.path_migrate"))
// 		if err == nil {
// 			err = mig.Up()
// 			if err != nil {
// 				if err == migrate.ErrNoChange {
// 					log.Debug("No database migration")
// 				} else {
// 					log.Error(err)
// 				}
// 			} else {
// 				log.Info("Migrate database success")
// 			}
// 			version, dirty, err := mig.Version()
// 			if err != nil && err != migrate.ErrNilVersion {
// 				log.Error(err)
// 			}
// 			log.Debug("Current DB version: " + strconv.FormatUint(uint64(version), 10) + "; Dirty: " + strconv.FormatBool(dirty))
// 		} else {
// 			log.Warn(err)
// 		}
// 	} else {
// 		log.Warn(err)
// 	}

// 	// Initialize Redis
// 	dbRedis := redis.NewClient(&redis.Options{
// 		Addr:     viper.GetString("redis.host") + ":" + viper.GetString("redis.port"),
// 		Username: viper.GetString("redis.username"),
// 		Password: viper.GetString("redis.password"),
// 		DB:       viper.GetInt("redis.database"),
// 		PoolSize: viper.GetInt("redis.max_connection"),
// 	})

// 	_, err = dbRedis.Ping(context.Background()).Result()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	log.Info("Redis connection established")

// 	// Initialize gRPC Pool
// 	var grpcPoolAuth, grpcPoolMember, grpcPoolWilayah, grpcPoolRealCount *grpcpool.Pool

// 	authConn := func() (client *grpc.ClientConn, err error) {
// 		address := fmt.Sprintf("%s:%s", viper.GetString("grpc.auth_service.host"), viper.GetString("grpc.auth_service.port"))
// 		client, err = grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		return
// 	}
// 	grpcPoolAuth, err = grpcpool.New(authConn, viper.GetInt("grpc.init"), viper.GetInt("grpc.capacity"), time.Duration(viper.GetInt("grpc.idle_duration"))*time.Second, time.Duration(viper.GetInt("grpc.max_life_duration"))*time.Second)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	memberConn := func() (client *grpc.ClientConn, err error) {
// 		address := fmt.Sprintf("%s:%s", viper.GetString("grpc.member_service.host"), viper.GetString("grpc.member_service.port"))
// 		client, err = grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		return
// 	}
// 	grpcPoolMember, err = grpcpool.New(memberConn, viper.GetInt("grpc.init"), viper.GetInt("grpc.capacity"), time.Duration(viper.GetInt("grpc.idle_duration"))*time.Second, time.Duration(viper.GetInt("grpc.max_life_duration"))*time.Second)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	wilayahConn := func() (client *grpc.ClientConn, err error) {
// 		address := fmt.Sprintf("%s:%s", viper.GetString("grpc.wilayah_service.host"), viper.GetString("grpc.wilayah_service.port"))
// 		client, err = grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		return
// 	}
// 	grpcPoolWilayah, err = grpcpool.New(wilayahConn, viper.GetInt("grpc.init"), viper.GetInt("grpc.capacity"), time.Duration(viper.GetInt("grpc.idle_duration"))*time.Second, time.Duration(viper.GetInt("grpc.max_life_duration"))*time.Second)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	realCountConn := func() (client *grpc.ClientConn, err error) {
// 		address := fmt.Sprintf("%s:%s", viper.GetString("grpc.realcountmgt_service.host"), viper.GetString("grpc.realcountmgt_service.port"))
// 		client, err = grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		return
// 	}
// 	grpcPoolRealCount, err = grpcpool.New(realCountConn, viper.GetInt("grpc.init"), viper.GetInt("grpc.capacity"), time.Duration(viper.GetInt("grpc.idle_duration"))*time.Second, time.Duration(viper.GetInt("grpc.max_life_duration"))*time.Second)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// Register repository & usecase
// 	Client := http.Client{}
// 	repoMySQLAccount := _RepoMySQLAccount.NewMySQLAccountRepository(dbConn)
// 	repoRedisAccount := _RepoRedisAccount.NewRedisAccountRepository(dbRedis)
// 	repoGRPCAuthorization := _RepoGRPCSaksi.NewGRPCAuthRepository(grpcPoolAuth)
// 	repoGRPCUser := _RepoGRPCSaksi.NewGRPCUserRepository(grpcPoolAuth)
// 	repoGRPCMember := _RepoGRPCSaksi.NewGRPCMemberRepository(grpcPoolMember)
// 	repoGRPCWilayah := _RepoGRPCSaksi.NewGRPCAdministrative(grpcPoolWilayah)
// 	repoGRPCRealCount := _RepoGRPCSaksi.NewGRPCRealCountRepository(grpcPoolRealCount)
// 	motionpayRepo := _MotionPayRepo.NewMotionPayRepository(Client)

// 	log.Println(motionpayRepo)

// 	// Register repository & usecase tps
// 	repoMySQLTps := _RepoMySQLAccount.NewMySQLTpsRepository(dbConn)
// 	usecaseTps := _UsecaseAccount.NewTpsUsecase(repoMySQLTps, repoGRPCMember, repoGRPCWilayah, repoGRPCAuthorization)

// 	// Register repository & usecase attendance
// 	repoMySQLAttendance := _RepoMySQLAccount.NewMySQLAttendanceRepository(dbConn)
// 	usecaseAttendance := _UsecaseAccount.NewAttendanceUsecase(repoMySQLAttendance, repoGRPCMember, repoMySQLAccount, repoGRPCWilayah, repoRedisAccount)

// 	// document evidence
// 	repoMySQLDocumentEvidence := _RepoMySQLAccount.NewMySQLDocumentEvidenceRepository(dbConn)
// 	usecaseDocumentEvidence := _UsecaseAccount.NewDocumentEvidenceUsecase(repoMySQLDocumentEvidence, repoGRPCMember, repoGRPCRealCount, repoGRPCWilayah)

// 	// recruiter
// 	repoMySQLRecruiter := _RepoMySQLAccount.NewMySQLRecruiterRepository(dbConn)

// 	usecaseRecruiter := _UsecaseAccount.NewRecruiterUsecase(repoMySQLRecruiter, repoGRPCMember, repoGRPCWilayah, repoRedisAccount, repoGRPCAuthorization, repoMySQLAccount, repoMySQLTps, repoMySQLAttendance)

// 	usecaseAccount := _UsecaseAccount.NewAccountUsecase(repoMySQLAccount, repoGRPCMember, repoGRPCAuthorization, repoRedisAccount, repoGRPCWilayah, repoMySQLTps, repoMySQLAttendance, repoMySQLRecruiter, repoGRPCUser, motionpayRepo)

// 	// Register Server Auth Saksi
// 	serverAuthSaksi := _RepoGRPCAuthSaksi.NewGRPCAuthorization(usecaseAccount, usecaseTps)

// 	// Initialize gRPC server
// 	go func() {
// 		listen, err := net.Listen("tcp", ":"+viper.GetString("server.grpc_port"))
// 		if err != nil {
// 			log.Fatalf("[ERROR] Failed to listen tcp: %v", err)
// 		}

// 		grpcServer := grpc.NewServer()
// 		_RepoGRPCSaksiAuth.RegisterDetailMemberServiceServer(grpcServer, serverAuthSaksi)
// 		_RepoGRPCSaksiAuth.RegisterSaksiManagementServiceServer(grpcServer, serverAuthSaksi)

// 		log.Println("gRPC server is running")
// 		if err := grpcServer.Serve(listen); err != nil {
// 			log.Fatalf("Failed to serve: %v", err)
// 		}
// 	}()

// 	// Initialize HTTP web framework
// 	app := fiber.New(fiber.Config{
// 		Prefork:       viper.GetBool("server.prefork"),
// 		StrictRouting: viper.GetBool("server.strict_routing"),
// 		CaseSensitive: viper.GetBool("server.case_sensitive"),
// 		BodyLimit:     viper.GetInt("server.body_limit"),
// 	})
// 	app.Use(logger.New())
// 	app.Use(recover.New())
// 	app.Use(cors.New(cors.Config{
// 		AllowOrigins: viper.GetString("middleware.allows_origin"),
// 	}))

// 	// HTTP routing
// 	app.Get(viper.GetString("server.base_path")+"/", func(c *fiber.Ctx) error {
// 		return c.SendString("Hello, World!")
// 	})
// 	if viper.GetBool("api_spec") {
// 		_DeliveryHTTPAccount.RouterOpenAPI(app)
// 	}
// 	_DeliveryHTTPAccount.RouterAPI(app, usecaseAccount, usecaseTps, usecaseAttendance, usecaseDocumentEvidence, usecaseRecruiter)

// 	// go func() {
// 	if err := app.Listen(":" + viper.GetString("server.port")); err != nil {
// 		log.Fatal(err)
// 	}
// 	// }()

// 	// Wait for interrupt signal to gracefully shutdown the server
// 	// quit := make(chan os.Signal, 1)
// 	// signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
// 	// <-quit
// 	// log.Info("Gracefully shutdown")
// 	// app.Shutdown()
// }
