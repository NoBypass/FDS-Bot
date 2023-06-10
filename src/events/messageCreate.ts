import { BotEvent } from '../types/discord'
import { Message } from 'discord.js'

const event: BotEvent = {
  name: 'messageCreate',
  once: true,
  // eslint-disable-next-line @typescript-eslint/no-empty-function
  execute: (message: Message) => {},
}

export default event
