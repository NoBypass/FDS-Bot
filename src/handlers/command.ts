import { Client, Routes } from 'discord.js'
import { REST } from '@discordjs/rest'
import { join } from 'path'
import { SlashCommand } from '../types/discord'
import { log } from 'console'
import chalk from 'chalk'
import registerFiles from '../lib/register-files'

module.exports = (client: Client) => {
  log(chalk.underline.magenta('\nCommands:\n'))

  const commands: SlashCommand[] = registerFiles<SlashCommand>(
    join(__dirname, '../commands'),
    (command) => {
      client.slashCommands.set(command.command.name, command)
      log(
        chalk.green(
          `Command '${chalk.yellow.bold(
            '/' + command.command.name,
          )}' registered`,
        ),
      )
    },
  )

  const { TOKEN, CLIENT_ID } = process.env
  const rest = new REST({ version: '10' }).setToken(TOKEN)

  rest
    .put(Routes.applicationCommands(CLIENT_ID), {
      body: commands.map((command) => command.command.toJSON()),
    })
    .then(() => {
      log(
        chalk.green(
          `Successfully loaded ${chalk.bold.underline(
            ` ${commands.length} `,
          )} command${commands.length == 1 ? '' : 's'}`,
        ),
      )
    })
    .catch(() => {
      log(chalk.red('Encountered error during registration'))
    })
}
