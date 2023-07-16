import { CacheType, Client, CommandInteraction } from 'discord.js'
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
        : client.on(event.name, (...args: unknown[]) => {
            switch (event.name) {
              case 'interactionCreate':
                console.log(
                  chalk.yellow.bold('Interaction fired: ') +
                    chalk.yellow.bold.underline(
                      (args[0] as CommandInteraction).user.username,
                    ) +
                    chalk.yellow.bold(' with args: ') +
                    chalk.yellow.bold.underline(args),
                )
                break
              default:
                console.log(
                  chalk.yellow.bold('Event fired: ') +
                    chalk.yellow.bold.underline(event.name) +
                    chalk.yellow.bold(' with args:') +
                    chalk.yellow.bold.underline(args),
                )
            }
            event.execute(...args)
          })

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
