package config

import (
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

func init() {
	if os.Getenv("MODE") != "product" {
		if err := godotenv.Load("config/.env"); err != nil {
			panic(err)
		}
	}

	app = new(application).init()
	jwtCfg = new(jwtConfig).init()
}

var (
	app    *application
	jwtCfg *jwtConfig
)

func Application() *application { //nolint
	return app
}

type application struct {
	Address          string
	DevMode          bool
	EventWorkerCount int
}

func (a *application) init() *application {
	a.Address = os.Getenv("ADDRESS")
	a.DevMode = os.Getenv("MODE") == "dev"

	c, err := strconv.Atoi(os.Getenv("EVENT_WORKER_COUNT"))
	if err != nil {
		panic(err)
	}

	a.EventWorkerCount = c

	return a
}

func JWT() *jwtConfig { //nolint
	return jwtCfg
}

type jwtConfig struct {
	Secret     string
	Expiration time.Duration
	Algorithm  *jwt.SigningMethodHMAC

	RefreshSecret     string
	RefreshExpiration time.Duration
	RefreshAlgorithm  *jwt.SigningMethodHMAC
}

func (j *jwtConfig) init() *jwtConfig {
	exp := os.Getenv("JWT_EXPIRATION")

	expiration, err := strconv.Atoi(exp)
	if err != nil {
		panic(err)
	}

	refreshExp := os.Getenv("JWT_REFRESH_EXPIRATION")

	refreshExpiration, err := strconv.Atoi(refreshExp)
	if err != nil {
		panic(err)
	}

	algorithm, ok := algorithms[os.Getenv("JWT_ALGORITHM")]
	if !ok {
		panic("invalid jwt algorithm")
	}

	refreshAlgorithm, ok := algorithms[os.Getenv("JWT_REFRESH_ALGORITHM")]
	if !ok {
		panic("invalid jwt algorithm")
	}

	j.Secret = os.Getenv("JWT_SECRET")
	j.Expiration = time.Second * time.Duration(expiration)
	j.Algorithm = algorithm
	j.RefreshSecret = os.Getenv("JWT_REFRESH_SECRET")
	j.RefreshExpiration = time.Second * time.Duration(refreshExpiration)
	j.RefreshAlgorithm = refreshAlgorithm

	return j
}

var algorithms = map[string]*jwt.SigningMethodHMAC{
	"HS256": jwt.SigningMethodHS256,
	"HS384": jwt.SigningMethodHS384,
	"HS512": jwt.SigningMethodHS512,
}
