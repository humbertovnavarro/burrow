Do these in order

# Music Bot Node (burrow-muse) 
- Downloads songs from youtube and caches them.
- Uses discordgo websocket to stream opus packets from cached songs
- Listens for messages using rabbitmq for jobs
- Handles Queue management internally
- Connects to voice channels
- Go?

# Burrow broker (burrow-bot)
- Processes Discord API events via websocket (discordgo)
- Sends jobs to burrow-muse via amqp
- Handles state persistance (song history, messages, etc) through postgresql
- TypeScript, Discord.js

# Burrow web (burrow-web)
- Frontend for burrow services
- uses next api to send events through amqp
- send persisted information on postgres through next api
- Sends jobs to burrow-muse via amqp
- Next.JS, TypeScript.
