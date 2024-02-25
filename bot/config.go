package bot

// REPLACE WITH YOUR CHATID
// Start servic and check output when messaging bot to see your id
// You will appear as not authorized in the console, but can see your id
var AuthorizedUsers = map[int]struct{}{
	0000000: {},
}

// To get notified if unauthorized users try to access your bot
var MainUser int64 = 0000000
