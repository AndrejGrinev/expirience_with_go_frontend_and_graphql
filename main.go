package main

import (
	"fmt"
	"log"
	"net/http"
	"tpro/stage1g/handlers/staff"
	"tpro/stage1g/utils"
)

func main() {
	log.Println("Start service...")
	var err error
	// получим данные из конфига
	utils.Cfg, err = utils.SetupConfig()
	if err != nil {
		fmt.Println("Ошибка загрузки конфига:", err)
		return
	}
	http.HandleFunc("/staff/moder_tender_company_invite", handlers.ModerInviteHandler)
	log.Fatal(http.ListenAndServe(utils.Cfg.Listen, nil))
}
