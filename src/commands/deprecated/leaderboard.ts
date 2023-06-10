import { Message, MessageEmbed } from "discord.js";
import { hypixel_api_key, desc, client } from "../src/index";
const path = require('path');
var ownname = path.basename(__filename);
import verifiedUsers from '../schemas/verified-users'

export default {
    callback: async (message: Message, ...args: string[]) => {

        var topTen = []
        var allMembers = await verifiedUsers.where('level').gt(1).select('level').select('memberid').select('xp').sort({level: -1,})

        const getNeededXP = (level) => {
            if (level < 10) {
                return Math.pow(level, 2) * 100
            } else if (level >= 10) {
                return 10000
            }
        }

        async function getXP(level) {
            let base = 0
            for (var i = 0; i < level; i++) {
                base += await getNeededXP(level - i)
            }
            return base
        }

        for (var i = 0; i < (allMembers as any).length; i++) {
            let levelXP = await getXP(allMembers[i].level)
            let totalXP = await levelXP + allMembers[i].xp
            let lvlProgress = 100 / (Math.pow(allMembers[i].level + 1, 2) * 100) * allMembers[i].xp
            if (lvlProgress < 10) var zero = '0'; else zero = '';
            let level = allMembers[i].level + ',' + zero + lvlProgress
            level = (level as any).split('.')
            topTen.push('<@' + allMembers[i].memberid + '> is level **' + level[0].substring(0, 4) + '** _(' + await totalXP + ' total xp)_')
        }

        var output = topTen.sort( (a, b) => parseFloat(a.match(/(\d+(?:\.\d+)?) total xp\b/)[1]) - parseFloat(b.match(/(\d+(?:\.\d+)?) total xp\b/)[1])).reverse();
        
        const exampleEmbed = new MessageEmbed()
            .setColor('#2F3136')
            .setTitle('XP Leaderboard:')
            .setDescription('#1  ' + topTen[0] + '\n#2  ' + topTen[1] + '\n#3  ' + topTen[2] + '\n#4  ' + topTen[3] + '\n#5  ' + topTen[4] + '\n#6  ' + topTen[5] + '\n#7  ' + topTen[6] + '\n#8  ' + topTen[7] + '\n#9  ' + topTen[8] + '\n#10 ' + topTen[9])
            .setFooter({ text: desc });

        message.channel.send({ embeds: [exampleEmbed] });
        message.delete()
    }
}