package utils

import (
	"fmt"
	"github.com/jessevdk/go-flags"
	"html/template"
	"math/rand"
	"strconv"
	"strings"
	"github.com/jackc/pgx/v4/pgxpool"
	"context"
)
// setupConfig loads flags from args (if given) or command flags and ENV otherwise
func SetupConfig(args ...string) (*Config, error) {
	var err error

	cfg := &Config{}
	p := flags.NewParser(cfg, flags.Default) //  HelpFlag | PrintErrors | PassDoubleDash

	if len(args) == 0 {
		_, err = p.Parse()
	} else {
		_, err = p.ParseArgs(args)
	}
	if err != nil {
		if e, ok := err.(*flags.Error); ok && e.Type == flags.ErrHelp {
			return nil, ErrGotHelp
		}
		return nil, ErrBadArgs
	}

	return cfg, nil
}
// Проверка сессии staff
func StaffAcl(s Graphql_Data_Adsession_result) (int, string) {
	// если нет данных, то показываем ошибку
        sid, _ := strconv.Atoi(s.Data.Adsession.Sid)
	if sid == 0 {
		return 1, "Сид просрочен или отсутствует."
	}
	// если тип пользователя не сотрудник ТПро, то ошибка
	if s.Data.Adsession.Face_type < 333 {
		return 2, "Доступ запрещён."
	}
	return 0, ""
}
// рандомная строка заданной длины
func RandSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
// STRING TO HTML
func ToHtml(value string) template.HTML {
	return template.HTML(fmt.Sprint(value))
}
// CONCATE STRINGS
func ConcatStr(strs ...string) string {
	return strings.Trim(strings.Join(strs, ""), " ")
}
// INT to STRING
func ToStr(value int) string {
	return strconv.Itoa(value)
}
// CONNECT TO DB
func SetupDb() (*pgxpool.Pool, error) {
	// получаем пул соединений
	config, _ := pgxpool.ParseConfig("")
	config.ConnConfig.Host = Cfg.PgHost
	n, err := strconv.ParseUint(Cfg.PgPort, 10, 16)
	if err != nil {
		fmt.Println("Ошибка в ParseUint(cfg.PgPort): %v\n", err)
		return nil, err
	}
	config.ConnConfig.Port = uint16(n)
	config.ConnConfig.User = Cfg.PgUser
	config.ConnConfig.Password = Cfg.PgPassword
	config.ConnConfig.Database = Cfg.PgDbName
	db, err := pgxpool.ConnectConfig(context.Background(), config)
	if err != nil {
		fmt.Println("Невозможно подсоединиться к базе: %v\n", err)
		return nil, err
	}
	return db, nil
}
