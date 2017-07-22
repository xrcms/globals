package globals

var (
	AllowSSL = false
)

const (
	// ConfigurationPath Путь к файлу настроек
	ConfigurationPath = "/etc/xrcms.conf"
	// GroupAdminsID Группа администраторов
	GroupAdminsID = 1
	// GroupGuestsID Группа для неавторизованных пользователей
	GroupGuestsID = 2
	// GroupUsersID Группа для авторизованных пользователей
	GroupUsersID = 3
	// MenuInfoblocksItemID ID пункта меню инфоблоков
	MenuInfoblocksItemID = 9
	// MenuModulesItemID ID пункта меню модулей
	MenuModulesItemID = 4
	// EmailConfirmationID ID шаблона письма подтверждения E-Mail
	EmailConfirmationID = 1
	// EmailRestorePasswordID ID шаблона письма восстановления пароля
	EmailRestorePasswordID = 2
	// EmailFeedbackID ID шаблона письма из формы обратной связи
	EmailFeedbackID = 3
	// Версия программы
	VERSION = "1.01.000"
	// Дата сборки
	RELEASE_DATE = "#RELEASE_DATE#"
	// Урл сайта CMS
	XR_DOMAIN = "xr-cms.ru"


	// NEW_INFOBLOCK_TEMPLATE Шаблон таблицы инфоблока
	NEW_INFOBLOCK_TEMPLATE = `
		CREATE TABLE IF NOT EXISTS "infoblock_#CODE#" (
			"id" BIGSERIAL,
			"code" Character( 255 ) DEFAULT ''::character varying,
			"name" Character Varying( 255 ) DEFAULT ''::character varying,
			"preview_picture" Bigint DEFAULT 0,
			"preview_text" Text DEFAULT ''::text,
			"detail_picture" Bigint DEFAULT 0,
			"detail_text" Text DEFAULT ''::text,
			"active_from" Timestamp Without Time Zone,
			"active_to" Timestamp Without Time Zone,
			"active" Boolean DEFAULT true,
			"author_id" Bigint DEFAULT 0,
			"editor_id" Bigint DEFAULT 0,
			"categories" Bigint[] DEFAULT '{}'::bigint[],
			"seo_title" Character Varying( 255 ) DEFAULT ''::character varying,
			"seo_keywords" Character Varying( 255 ) DEFAULT ''::character varying,
			"seo_description" Text DEFAULT ''::text,
			"updated_at" Timestamp With Time Zone DEFAULT now(),
			"created_at" Timestamp With Time Zone DEFAULT now()
		);
		CREATE UNIQUE INDEX IF NOT EXISTS "unique_infoblock_#CODE#_id" ON "infoblock_#CODE#" USING btree( "id" Asc NULLS Last );
		CREATE UNIQUE INDEX IF NOT EXISTS "unique_infoblock_#CODE#_code" ON "infoblock_#CODE#" USING btree( "code" Asc NULLS Last );
		CREATE INDEX IF NOT EXISTS "#CODE#_code_only_active_index" ON "infoblock_#CODE#" (LOWER("code")) WHERE "active" = true;
	`
)
