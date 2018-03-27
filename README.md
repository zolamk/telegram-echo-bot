# A Telegram Echo Bot Using Netlify Functions

This is a simple telegram echo bot built using golang and Netlify Functions.

## Steps

1. Clone This Repo

2. Set TELEGRAM_BOT_TOKEN environment variable in netlify

3. set your web hook like this
    `curl --request POST --url https://api.telegram.org/bot<TELEGRAM_BOT_TOKEN>/setWebhook --header 'content-type: application/json' --data '{"url": "<netlify site url>/.netlify/functions/echo"}'`

4. and you're good to go.