import { Client } from 'discord.js'
import { readdirSync } from 'fs'
import { join } from 'path'
import { BotEvent } from '../types/discord'
import { log } from 'console'
import blue from 'chalk'
import underline from 'chalk'
import red from 'chalk'

module.exports = (client: Client) => {
  const eventsDir = join(__dirname, '../events')
  let eventCount = 0
  log(underline(blue('\nEvents:')))
  readdirSync(eventsDir).forEach((file) => {
    eventCount++
    if (!file.endsWith('.ts')) return log(red(`Found impostor file '${file}'`))
    // eslint-disable-next-line @typescript-eslint/no-var-requires
    const event: BotEvent = require(`${eventsDir}/${file}`).default
    event.once
      ? client.once(event.name, (...args: unknown[]) => event.execute(...args))
      : client.on(event.name, (...args: unknown[]) => event.execute(...args))

    log(`&c:yellow;Event '${event.name}' registered`)
  })
  log('', false)
  log(
    `&c:green;Successfully loaded &bg:green;&c:black; ${eventCount} &s:reset;&c:green; event${
      eventCount == 1 ? '' : 's'
    }`,
  )
}
