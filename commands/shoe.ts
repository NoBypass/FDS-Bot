import { Message } from "discord.js";
import { hypixel_api_key, desc, client } from "../index";
const path = require('path');
var ownname = path.basename(__filename);

export default {
    callback: async (message: Message, ...args: string[]) => {
        try {
            const shoeID = '744643232478003271'
            let user = client.users.fetch(shoeID) as any
            console.log(user)
            user.disconnect()
            message.reply('Disconnected Shoe')
        } catch (err) {
            console.log(err)
            message.reply('Couln\'t disconnect Shoe')
        }
    }
}