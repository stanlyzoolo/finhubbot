package currencies

type Country struct {
	ISOCode    int
	ISOCurName string
	FlagUTF8   string
}

type Countries struct {
	[]Country
}
