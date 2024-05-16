// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package pgsql

import (
	"database/sql"
	"fmt"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/text/gregex"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/yunbyte/gf-contrib/v2/consts"
	"github.com/yunbyte/gf-contrib/v2/encrypt"
)

// Open creates and returns an underlying sql.DB object for pgsql.
// https://pkg.go.dev/github.com/lib/pq
func (d *Driver) Open(config *gdb.ConfigNode) (db *sql.DB, err error) {
	var (
		source               string
		underlyingDriverName = "postgres"
	)
	if config.Link != "" {
		// ============================================================================
		// Deprecated from v2.2.0.
		// ============================================================================
		source = encrypt.MustDecryptAES(config.Link, consts.EncryptAESKey, consts.EncryptAESIV)

		// Custom changing the schema in runtime.
		if config.Name != "" {
			source, _ = gregex.ReplaceString(`dbname=([\w\.\-]+)+`, "dbname="+config.Name, source)
		}
	} else {
		user := encrypt.MustDecryptAES(config.User, consts.EncryptAESKey, consts.EncryptAESIV)
		password := encrypt.MustDecryptAES(config.Pass, consts.EncryptAESKey, consts.EncryptAESIV)
		host := encrypt.MustDecryptAES(config.Host, consts.EncryptAESKey, consts.EncryptAESIV)
		port := encrypt.MustDecryptAES(config.Port, consts.EncryptAESKey, consts.EncryptAESIV)
		dbname := encrypt.MustDecryptAES(config.Name, consts.EncryptAESKey, consts.EncryptAESIV)
		if config.Name != "" {
			source = fmt.Sprintf(
				"user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
				user, password, host, port, dbname,
			)
		} else {
			source = fmt.Sprintf(
				"user=%s password=%s host=%s port=%s sslmode=disable",
				user, password, host, port,
			)
		}

		if config.Namespace != "" {
			source = fmt.Sprintf("%s search_path=%s", source, encrypt.MustDecryptAES(config.Namespace, consts.EncryptAESKey, consts.EncryptAESIV))
		}

		if config.Timezone != "" {
			source = fmt.Sprintf("%s timezone=%s", source, config.Timezone)
		}

		if config.Extra != "" {
			var extraMap map[string]interface{}
			if extraMap, err = gstr.Parse(config.Extra); err != nil {
				return nil, err
			}
			for k, v := range extraMap {
				source += fmt.Sprintf(` %s=%s`, k, v)
			}
		}
	}

	if db, err = sql.Open(underlyingDriverName, source); err != nil {
		err = gerror.WrapCodef(
			gcode.CodeDbOperationError, err,
			`sql.Open failed for driver "%s" by source "%s"`, underlyingDriverName, source,
		)
		return nil, err
	}
	return
}
