## ALGALON, a Telegram GPT service

### Instructions

To use Algalon, start with the following setup:

1. Get your OpenAI API key and add it to gpt/key.go
2. Create a Telegram bot using @BotFather on Telegram
3. Create a new bot entry in bot/modes.go in which you paste the Telegram bot API key
4. Start the service and message the bot. You won't be authenticated yet, but you can find your Chat ID in the log / console
5. Paste this Chat ID of your Account into bot/config.go
6. Restart the service, your Telegram Account is now authenticated for your bot

If you don't need authentication, remove the check in bot/algalon.go 32-39. Keep in mind that any user can access your bot and your API credit may run out.