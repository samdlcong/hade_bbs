// Package http API.
// @title hade_bbs
// @version 1.0
// @description hade论坛
// @termsOfService https://github.com/swaggo/swag

// @contact.name samdlcong
// @contact.email samnew@126.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath /
// @query.collection.format multi

// @securityDefinitions.basic BasicAuth

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @x-extension-openapi {"example": "value on a json format"}

package http

import (
	_ "hade_bbs/app/http/swagger"
)
