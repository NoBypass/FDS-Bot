import { Client, Routes, SlashCommandBuilder } from 'discord.js'
import { REST } from '@discordjs/rest'
import { readdirSync } from 'fs'
import { join } from 'path'
import { SlashCommand } from '../types/discord'
import { log } from 'console'
import blue from 'chalk'
import underline from 'chalk'
import green from 'chalk'
import greenBg from 'chalk'
import black from 'chalk'
import red from 'chalk'
import yellow from 'chalk'

module.exports = (client: Client) => {
  const commands: SlashCommandBuilder[] = []
  const commandsDir = join(__dirname, '../commands')
  const { TOKEN, CLIENT_ID } = process.env

  readdirSync(commandsDir).forEach((file) => {
    if (!file.endsWith('.ts')) return
    // eslint-disable-next-line @typescript-eslint/no-var-requires
    const command: SlashCommand = require(`${commandsDir}/${file}`).default
    commands.push(command.command)
    client.slashCommands.set(command.command.name, command)
  })
  const rest = new REST({ version: '10' }).setToken(TOKEN)

  log(underline(blue('Commands:')))
  commands.forEach((command) => {
    log(yellow(`Command '/${command.name}' registered`))
  })
  rest
    .put(Routes.applicationCommands(CLIENT_ID), {
      body: commands.map((command) => command.toJSON()),
    })
    .then(() => {
      log(
        green('Successfully loaded ') +
          greenBg(
            black(
              `${commands.length} &s:reset;&c:green; command${
                commands.length == 1 ? '' : 's'
              }`,
            ),
          ),
      )
    })
    .catch(() => {
      log(red('Encountered error during registration'))
    })
}
