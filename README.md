# Основная страница маркетинга на GO

оригинал - **/work/support/SUXX/moder_tender_company_invite.suxx**

## предварительно:

1. выполнить **make conf**

<details><summary>2. в файле **.env** отредактировать путь к статике и апи, настройки базы, хост</summary>

```
путь к статике, например, **test.tender.pro**, можно локальный, если важен, допустим в разработке, **js**

путь к апи зависит от того как развернут проект. Если через **dcape**, то и указывать надо его адрес, например, **tpro.dev.lan**

настройки базы примерно такие:
PGUSER=tpro2_dev PGHOST=db PGPORT=5432 PGDBNAME=tpro2_dev PGPASSWORD=WBRwdfDXgmujAA

sudo bash -c 'echo "127.0.0.1 marketing.tpro.dev.lan" >> /etc/hosts'

```
</details>

## компиляция и запуск:
* **make up**

## страница:

```
Страница доступна _marketing.tpro.dev.lan/staff/moder_tender_company_invite?sid={{SID сотрудника или SID орга(доп-но нужен параметр acc_marketing=1 с включенной настройкой)}}&tenderid={{ID конкурса}}_

Можно на неё зайти из **tpro.dev.lan** (залогинится под маркетингом, зайти в конкурс и рядом со старой ссылкой будет новая - для тестовых)
```

## структура кода:

```
prototype_go
├── lib
│   └── handlers
│       ├── common
│       ├── public
│       ├── staff
│           └── moder_tender_company_invite.go
│       └── user
│   └── utils
│       └── utils.go
│       └── utils_test.go
├── www
│   ├── includes
│       ├── header.gohtml
│       ├── footer.gohtml
│       └── page_path.gohtml
│   └── tmpl
│       ├── common
│       ├── public
│       ├── staff
│           └── moder_tender_company_invite.gohtml
│       └── user
├── dockerfile
├── docker-compose.yml
├── readme.md
├── main.go
└── makefile
```
