import { Client } from 'discord.js'
import { join } from 'path'
import { BotEvent } from '../types/discord'
import { log } from 'console'
import chalk from 'chalk'
import registerFiles from '../lib/register-files'

module.exports = (client: Client) => {
  log(chalk.magenta.underline('\nEvents:\n'))

  const events = registerFiles<BotEvent>(
    join(__dirname, '../events'),
    (event) => {
      event.once
        ? client.once(event.name, (...args: unknown[]) =>
            event.execute(...args),
          )
        : client.on(event.name, (...args: unknown[]) => event.execute(...args))

      log(chalk.green(`Event '${chalk.yellow.bold(event.name)}' registered`))
    },
  )

  log(
    chalk.green(
      `Successfully loaded ${chalk.bold.underline(` ${events.length} `)} event${
        events.length == 1 ? '' : 's'
      }`,
    ),
  )
}
