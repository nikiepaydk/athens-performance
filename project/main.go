package main

import (
	_ "dario.cat/mergo"
	_ "github.com/PuerkitoBio/goquery"
	_ "github.com/alicebob/miniredis/v2"
	_ "github.com/creasty/defaults"
	_ "github.com/go-gormigrate/gormigrate/v2"
	_ "github.com/gofiber/fiber/v2"
	_ "github.com/gofiber/template/html/v2"
	_ "github.com/google/go-querystring/query"
	_ "github.com/google/uuid"
	_ "github.com/jaswdr/faker/v2"
	_ "github.com/microcosm-cc/bluemonday"
	_ "github.com/mileusna/useragent"
	_ "github.com/pkg/errors"
	_ "github.com/redis/go-redis/v9"
	_ "github.com/riverqueue/river"
	_ "github.com/sony/sonyflake"
	_ "github.com/stretchr/testify"
	_ "github.com/thomaspoignant/go-feature-flag"
	_ "go.mozilla.org/pkcs7"
	_ "golang.org/x/crypto/ssh"
	_ "golang.org/x/net/ipv4"
	_ "golang.org/x/text"
	_ "gorm.io/driver/sqlite"
	_ "gorm.io/gorm"
)

func main() {
	println("Hello World!")
}
