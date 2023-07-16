import { APIButtonComponentWithCustomId, Client } from 'discord.js'
import { join } from 'path'
import registerFiles from '../lib/register-files'
import { Button } from '../types/discord'
import { log } from 'console'
import chalk from 'chalk'

module.exports = (client: Client) => {
  log(chalk.underline.magenta('\nButtons:\n'))

  const buttons = registerFiles<Button>(
    join(__dirname, '../components/buttons'),
    (button) => {
      const buttonName = (
        button.button.data as Partial<APIButtonComponentWithCustomId>
      ).custom_id
      if (!buttonName) throw new Error('Button custom_id not found')
      client.buttons.set(buttonName, button)
      log(chalk.green(`Button '${chalk.yellow.bold(buttonName)}' registered`))
    },
  )

  log(
    chalk.green(
      `Successfully loaded ${chalk.bold.underline(
        ` ${buttons.length} `,
      )} button${buttons.length == 1 ? '' : 's'}`,
    ),
  )
}
