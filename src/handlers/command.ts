import { Client, Routes, SlashCommandBuilder } from 'discord.js'
import { REST } from '@discordjs/rest'
import { readdirSync } from 'fs'
import { join } from 'path'
import { SlashCommand } from '../types/discord'
import { log } from 'console'
import chalk from 'chalk'

module.exports = (client: Client) => {
  log(chalk.underline.magenta('\nCommands:\n'))

  const commands: SlashCommandBuilder[] = []
  const commandsDir = join(__dirname, '../commands')
  
  readdirSync(commandsDir).forEach((file) => {
    if (!file.endsWith('.ts')) return
    // eslint-disable-next-line @typescript-eslint/no-var-requires
    const command: SlashCommand = require(`${commandsDir}/${file}`).default
    try {
      client.slashCommands.set(command.command.name, command)
      commands.push(command.command)
    } catch (error) {
      log(chalk.red(`Failed to register ${chalk.bold(`'${commandsDir}/${file}'`)}`))
    }
  })

  const { TOKEN, CLIENT_ID } = process.env
  const rest = new REST({ version: '10' }).setToken(TOKEN)
  
  commands.forEach((command) => {
    log(chalk.green(`Command '${chalk.yellow.bold('/'+command.name)}' registered`))
  })
  rest
    .put(Routes.applicationCommands(CLIENT_ID), {
      body: commands.map((command) => command.toJSON()),
    })
    .then(() => {
      log(
    chalk.green(`Successfully loaded ${chalk.bold.underline(` ${commands.length} `)} command${
      commands.length == 1 ? '' : 's'
    }`),
      )
    })
    .catch(() => {
      log(chalk.red('Encountered error during registration'))
    })
}
