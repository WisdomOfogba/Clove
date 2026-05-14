package global

import (
	"context"
	"encoding/base64"
	"os"
	"reflect"
	"sync"
	"time"

	"github.com/bwmarrin/snowflake"
	"github.com/chibx/vendor-pulse/internal/squadco"
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/go-playground/validator/v10"
	pgxdecimal "github.com/jackc/pgx-shopspring-decimal"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog"
	"github.com/shopspring/decimal"
	pgxUUID "github.com/vgarvardt/pgx-google-uuid/v5"
	"google.golang.org/genai"
)

var (
	Logger      = initLogger()
	Validator   = validator.New(validator.WithRequiredStructEnabled())
	DB          *pgxpool.Pool
	Redis       *redis.Client
	SecretKey   []byte
	SnowFlake   *snowflake.Node
	Cloudinary  *cloudinary.Cloudinary
	AIClient    *genai.Client
	SquadClient *squadco.Client
)

func decimalCustomTypeFunc(field reflect.Value) interface{} {
	if value, ok := field.Interface().(decimal.Decimal); ok {
		return value.InexactFloat64()
	}
	return nil
}

func TagNameFunc(fld reflect.StructField) string {
	name := fld.Tag.Get("name")
	if name == "" {
		return fld.Name
	}
	return name
}

func asPointer[T any](a T) *T {
	return &a
}

func getEnv(env string, sub ...string) string {
	val := os.Getenv(env)
	if val == "" {
		if len(sub) > 0 {
			return sub[0]
		}
		panic("Environment Variable " + env + " not set")
	}
	return val
}

func loadKey(keyEnv string) []byte {
	keyBase64 := getEnv(keyEnv)

	var err error
	key, err := base64.StdEncoding.DecodeString(keyBase64)
	if err != nil {
		Logger.Fatal().Err(err).Str("secret_key", keyEnv).Msg("Invalid Secret Key")
	}
	return key
}

func initLogger() *zerolog.Logger {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	enableDebug := getEnv("DEBUG", "0")
	if enableDebug != "1" {
		// Disable debug logging
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}
	return asPointer(zerolog.New(os.Stdout))
}

func initGenAI(ctx context.Context) {
	apiKey := getEnv("GEN_AI_KEY")
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  apiKey,
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		Logger.Fatal().Err(err).Msg("Could not create Gen AI client")
		return
	}

	AIClient = client
}

func initSquadClient() {
	apiKey := getEnv("SQUAD_API_KEY")
	isProduction := getEnv("IS_SQUAD_PROD", "0") == "1"
	client, err := squadco.NewSquadClient(squadco.SquadOption{
		ApiKey: apiKey,
	})
	if err != nil {
		Logger.Fatal().Err(err).Msg("Could not create squadco client")
		return
	}
	client.SetEnvironment(isProduction)
	SquadClient = client
}

func newDB() *pgxpool.Pool {
	dbUrl := getEnv("DATABASE_URL")
	var db *pgxpool.Pool
	var err error

	for range 5 {
		db, err = pgxpool.New(context.Background(), dbUrl)
		db.Config().AfterConnect = func(_ context.Context, conn *pgx.Conn) error {
			pgxdecimal.Register(conn.TypeMap())
			pgxUUID.Register(conn.TypeMap())
			return nil
		}

		if err != nil {
			Logger.Error().Err(err).Msg("failed to initialize db conn for catalog service")
		} else {
			return db
		}

		Logger.Debug().Msg("Backing off for 2 seconds...")
		time.Sleep(100 * time.Millisecond)
	}

	Logger.Fatal().Msg("Could not connect to database after multiple retries")
	return nil
}

func newCloudinary() *cloudinary.Cloudinary {
	cldKey := getEnv("CLOUDINARY_KEY")
	cldSecret := getEnv("CLOUDINARY_SECRET")
	cldName := getEnv("CLOUDINARY_CLOUD_NAME")
	cld, err := cloudinary.NewFromParams(cldName, cldKey, cldSecret)
	if err != nil {
		Logger.Error().Err(err).Msg("failed to initialize cloudinary for catalog service")
		Logger.Fatal().Msg("Error setting up Cloudinary!!!")
	}

	return cld
}

func newRedis() *redis.Client {
	redisUrl := getEnv("REDIS_URL")
	opts, err := redis.ParseURL(redisUrl)
	if err != nil {
		Logger.Error().Err(err).Msg("failed to parse redis url for catalog service")
		Logger.Fatal().Msg("REDIS_URL should be set!!!")
	}

	client := redis.NewClient(opts)

	for range 5 {
		cmd := client.Ping(context.Background())
		err = cmd.Err()
		if err != nil {
			Logger.Error().Err(err).Msg("failed to connect to redis for catalog service")
		} else {
			return client
		}

		Logger.Debug().Msg("Backing off for 2 seconds...")
		time.Sleep(100 * time.Millisecond)
	}

	Logger.Fatal().Msg("Could not connect to Redis after multiple retries")
	return nil
}

func InitGlobals() {
	Validator.RegisterTagNameFunc(TagNameFunc)
	Validator.RegisterCustomTypeFunc(decimalCustomTypeFunc, decimal.Decimal{})
	var err error
	var wg sync.WaitGroup
	SecretKey = loadKey("SECRET_KEY")
	SnowFlake, err = snowflake.NewNode(1)
	if err != nil {
		Logger.Fatal().Err(err).Send()
	}
	wg.Go(func() {
		DB = newDB()
	})
	wg.Go(func() {
		Redis = newRedis()
	})

	wg.Go(func() {
		Cloudinary = newCloudinary()
	})

	wg.Go(func() {
		initGenAI(context.Background())
	})

	wg.Wait()

	initSquadClient()
}
