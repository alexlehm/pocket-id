//testing
package appconfig

import (
)

type ValueReturn struct {
//	valueStore string
	Value string
}

func getValueReturn(value string) ValueReturn {
	return ValueReturn {
		Value: value,
	}
}

type DbConfigStruct struct {
	AppName ValueReturn
	SmtpFrom ValueReturn
	SmtpPort ValueReturn
	SmtpHost ValueReturn
	SmtpUser ValueReturn
	SmtpPassword ValueReturn
	SmtpTls ValueReturn
	SmtpSkipCertVerify ValueReturn
}

func NewDbConfig() *DbConfigStruct {
	return &DbConfigStruct{
		AppName:	getValueReturn("app-name"),
		SmtpFrom:	getValueReturn("alex@lehmann.cx"),
		SmtpPort:	getValueReturn("587"),
		SmtpHost:	getValueReturn("localhost"),
		SmtpUser:	getValueReturn(""),
		SmtpPassword:	getValueReturn(""),
		SmtpTls:	getValueReturn("none"),
		SmtpSkipCertVerify:	getValueReturn("false"),
	}
}

type AppConfigService struct {
	DbConfig	*DbConfigStruct
}

func NewAppConfigService() (*AppConfigService, error) {
	return &AppConfigService {
		DbConfig: getDbConfig(),
	}, nil
}

func getDbConfig() *DbConfigStruct {
  return NewDbConfig()
}

