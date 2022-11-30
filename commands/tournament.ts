// import { Message, MessageEmbed } from "discord.js";
// import descEmbed from "../counters/descEmbed";
// import { hypixel_api_key, desc, client } from "../index";
// 
// export default {
//     callback: async (message: Message, ...args: string[]) => {
//         function formatReminder(d) {return descEmbed(d + ' Please make sure to format your message like this: "-tournament <team size> <players sepperated by a comma>"', message)}
//         if (typeof args[0] == 'number') formatReminder('You didn\'t enter a number for the team size' )
//         let len = parseInt(args[0])
//         if (len >= args.length - 1) formatReminder('You have less players than you want per team.')
// 
//         let teams = []; let count = 0
//         for (let i = 1; i < args.length; i++) {
//             let tempTeam = []
//             for (let j = 0; j < len; j++) {
//                 tempTeam.push(args[count].slice(0, 12)); count++
//             }
//             teams.push(tempTeam)
//         }
// 
//         if (teams[teams.length-1].length < len / 2) teams[teams.length-2].append(teams[teams.length-1])
// 
//         
//     }
// }