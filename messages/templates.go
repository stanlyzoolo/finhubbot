package messages

// Для одной единицы валюты
const USDandEUR = `
{{.CountryFlagUTF8}}{{.Name}}: {{.OfficialRate}}`

// Для BYN за 100 RUB
const RUB = `
{{.CountryFlagUTF8}}{{.Name}} за {{.Scale}}: {{.OfficialRate}}`

// Для коммерческих банков
const Commercial = `{{.Bank.Name}}:
{{.USD.Flag}}{{.USD.Buying}}/{{.USD.Selling}} {{.EUR.Flag}}{{.EUR.Buying}}/{{.EUR.Selling}} {{.RUB.Flag}}{{.RUB.Buying}}/{{.RUB.Selling}}`  
