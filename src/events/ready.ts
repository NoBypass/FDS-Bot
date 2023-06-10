import { BotEvent } from '../types/discord'
import { log } from 'console'
import black from 'chalk'
import blueBg from 'chalk'

const event: BotEvent = {
  name: 'ready',
  once: true,
  execute: () => {
    log(black(blueBg(' Bot started up ')))
  },
}

export default event
