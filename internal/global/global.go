package global

import (
	"context"
	"os"
	"reflect"
	"sync"
	"time"

	"github.com/chibx/vendor-pulse/internal/utils"
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog"
)

var (
	Logger     = initLogger()
	Validator  = validator.New(validator.WithRequiredStructEnabled())
	DB         *pgxpool.Pool
	Redis      *redis.Client
	Cloudinary *cloudinary.Cloudinary
)

func TagNameFunc(fld reflect.StructField) string {
	name := fld.Tag.Get("name")
	if name == "" {
		return fld.Name
	}
	return name
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

func initLogger() *zerolog.Logger {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	enableDebug := getEnv("DEBUG", "0")
	if enableDebug != "1" {
		// Disable debug logging
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}
	return utils.AsPointer(zerolog.New(os.Stdout))
}

func newDB() *pgxpool.Pool {
	dbUrl := getEnv("DATABASE_URL")
	var db *pgxpool.Pool
	var err error

	for range 5 {
		db, err = pgxpool.New(context.Background(), dbUrl)

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
	var wg sync.WaitGroup
	wg.Go(func() {
		DB = newDB()
	})
	wg.Go(func() {
		Redis = newRedis()
	})

	wg.Go(func() {
		Cloudinary = newCloudinary()
	})

	wg.Wait()
}
