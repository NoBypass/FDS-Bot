import { Client, Collection, GatewayIntentBits } from 'discord.js'
import { Button, SlashCommand } from './types/discord'
import { config } from 'dotenv'
import { readdirSync } from 'fs'
import { join } from 'path'
import { log } from 'console'
import chalk from 'chalk'
// import { login } from './src/lib/api'

const { Guilds, MessageContent, GuildMessages, GuildMembers } =
  GatewayIntentBits
export const client = new Client({
  intents: [Guilds, MessageContent, GuildMessages, GuildMembers],
})
config()

log(chalk.bold.magenta.underline('\nFDS Bot v2 is starting'))
client.slashCommands = new Collection<string, SlashCommand>()
client.cooldowns = new Collection<string, number>()
client.buttons = new Collection<string, Button>()
// client.auth = await login('FDS-Bot', process.env.TOKEN)

const handlersDir = join(__dirname, "./handlers")
readdirSync(handlersDir).forEach(handler => {
    require(`${handlersDir}/${handler}`)(client)
})

client.login(process.env.TOKEN)