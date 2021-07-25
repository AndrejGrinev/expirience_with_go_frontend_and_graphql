package utils

import (
	"errors"
	"github.com/jackc/pgx/v4/pgxpool"
	"regexp"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

var (
	// ErrGotHelp returned after showing requested help
	ErrGotHelp = errors.New("help printed")
	// ErrBadArgs returned after showing command args error message
	ErrBadArgs = errors.New("option error printed")
)

var Cfg *Config

var Db *pgxpool.Pool
// правило PlaceCode
var ValidPlaceCode = regexp.MustCompile(`^[PARC]\d+$`)

