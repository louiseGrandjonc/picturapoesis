package constants

const (
	DateRegexFR = `(?:(\d{1,2})\s*(?:er)?[\s+\/]*([jJ]anvier|[fF][eé]vrier|[mM]ars|[aA]vril|[mM]ai|[jJ]uin|[jJ]uillet|[aA]o[uû]t|[sS]eptembre|[oO]ctobre|[nN]ovembre|[dD][eé]cembre|\d{1,2})\s+(20\d{2})?\s*(?:-|\/|au|.)\s*(\d{1,2})\s*(?:er)?[\s+\/]*([jJ]anvier|[fF][eé]vrier|[mM]ars|[aA]vril|[mM]ai|[jJ]uin|[jJ]uillet|[aA]o[uû]t|[sS]eptembre|[oO]ctobre|[nN]ovembre|[dD][eé]cembre|\d{1,2})\s+(20\d{2})?)|(?:[aAàÀ] partir du (\d{1,2})[\s+\/]([jJ]anvier|[fF][eé]vrier|[mM]ars|[aA]vril|[mM]ai|[jJ]uin|[jJ]uillet|[aA]o[uû]t|[sS]eptembre|[oO]ctobre|[nN]ovembre|[dD][eé]cembre|\d{1,2})\s+(20\d{2})?)`

	DateRegexEN = `(\d{1,2})[\s+\/](january|february|march|april|may|june|jully|agust|september|october|november|december|\d{1,2})\s+(20\d{2})?\s*(?:-|\/|to|.)\s*(\d{1,2})[\s+\/](january|february|march|april|may|june|jully|agust|september|october|november|december|\d{1,2})\s+(20\d{2})?`

	MonthRegex = `([jJ]anvier|[fF][eé]vrier|[mM]ars|[aA]vril|[mM]ai|[jJ]uin|[jJ]uillet|[aA]o[uû]t|[sS]eptembre|[oO]ctobre|[nN]ovembre|[dD][eé]cembre)`
)

var (
	MonthFR = []string{
		"[jJ]anvier",
		"[fF][eé]vrier",
		"[mM]ars",
		"[aA]vril",
		"[mM]ai",
		"[jJ]uin",
		"[jJ]uillet",
		"[aA]o[uû]t",
		"[sS]eptembre",
		"[oO]ctobre", "[nN]ovembre", "[dD][eé]cembre"}
)
