import chalk from 'chalk'
import { log } from 'console'
import { readdirSync } from 'fs'

const registerFiles = <T>(dir: string, run?: (element: T) => void) => {
  const elements: T[] = []

  readdirSync(dir).forEach((file) => {
    if (!file.endsWith('.ts')) return
    // eslint-disable-next-line @typescript-eslint/no-var-requires
    const element: T = require(`${dir}/${file}`).default
    try {
      if (run) run(element)
      elements.push(element)
    } catch (error) {
      log(chalk.red(`Failed to register ${chalk.bold(`'${dir}/${file}'`)}`))
    }
  })

  return elements
}

export default registerFiles
