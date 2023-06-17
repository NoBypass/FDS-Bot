import { Client } from 'discord.js'
import { readdirSync } from 'fs'
import { join } from 'path'
import { BotEvent } from '../types/discord'
import { log } from 'console'
import chalk from 'chalk'

module.exports = (client: Client) => {
  const eventsDir = join(__dirname, '../events')
  let eventCount = 0
  log(chalk.magenta.underline('\nEvents:\n'))
  readdirSync(eventsDir).forEach((file) => {
    eventCount++
    if (!file.endsWith('.ts')) return //log(red(`Found impostor file '${file}'`))
    // eslint-disable-next-line @typescript-eslint/no-var-requires
    const event: BotEvent = require(`${eventsDir}/${file}`).default
    event.once
      ? client.once(event.name, (...args: unknown[]) => event.execute(...args))
      : client.on(event.name, (...args: unknown[]) => event.execute(...args))

    log(chalk.green(`Event '${chalk.yellow.bold(event.name)}' registered`))
  })
  log(
    chalk.green(
      `Successfully loaded ${chalk.bold.underline(` ${eventCount} `)} event${
        eventCount == 1 ? '' : 's'
      }`,
    ),
  )
}
