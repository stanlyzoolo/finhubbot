package messages

// Для одной единицы валюты
const MessageTemplateScale = `
{{.CountryFlagUTF8}}{{.Name}}: {{.OfficialRate}}`

// Для N единиц валюты
const MessageTemplateScales = `
{{.CountryFlagUTF8}}{{.Name}} за {{.Scale}}: {{.OfficialRate}}`
