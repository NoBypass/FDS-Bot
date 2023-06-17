import { BotEvent } from "../types/discord"
import { log } from "console"
import chalk from "chalk"

const event : BotEvent = {
    name: "ready",
    once: true,
    execute: () => {
        log(chalk.green('\nBot is online\n'))
    }
}

export default event