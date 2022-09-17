  /*
  async function refreshAllRanks() {
    const users = await verifiedUsers.find({});

    const userMap = {};
    users.forEach((user) => {
        userMap[user.ign] = user;
    });
    for (let user in userMap) {
      let userid = userMap[user].memberid
      let membe = await guild.members.cache.get(await userid)
      console.log(membe)
      let index = userMap[user].customstats.index
      (await membe).roles.remove("964653007117627453");
      (await membe).roles.remove("964652871113142282");
      (await membe).roles.remove("964652446154633256");
      var role = ""
      if (index < 100) {
        //role = guild.roles.cache.find(r => r.id === "964653007117627453")
        role = "964653007117627453"
      }
      else if (index >= 100 && index < 200) {
          //role = guild.roles.cache.find(r => r.id === "964652871113142282")
          role = "964652871113142282"
      }
      else if (index >= 200 && index < 300) {
          //role = guild.roles.cache.find(r => r.id === "964652446154633256")
          role = "964652446154633256"
      }
      else if (index >= 400) {
          //role = guild.roles.cache.find(r => r.id === "964651864052334683")
          role = "964651864052334683"
      }
      (await membe).roles.add(role)
    }
  }
  refreshAllRanks()*/

import { Message, MessageEmbed } from "discord.js";
import { hypixel_api_key, desc, client } from "../index";
import verifiedUsers from "../schemas/verified-users";
import { getApiData } from '../counters/get-api-data'
import { getIndex } from '../counters/index-count'

export default {
    callback: async (message: Message, ...args: string[]) => {
            let user = message.member
            const member = await verifiedUsers.findOne({ memberid: user.id }) as any
            const uuid = member.uuid
            const data = await getApiData(uuid) as any
            const indexes = await getIndex(data) as any
            const index = indexes.index
            console.log(data)
            user.roles.remove("964653007117627453");
            user.roles.remove("964652871113142282");
            user.roles.remove("964652446154633256");
            var role = ""
            if (index < 100) {
              role = "964653007117627453"
            }
            else if (index >= 100 && index < 200) {
                role = "964652871113142282"
            }
            else if (index >= 200 && index < 300) {
                role = "964652446154633256"
            }
            else if (index >= 400) {
                role = "964651864052334683"
            }
            user.roles.add(role)

            const roles = new MessageEmbed()
            .setColor('#000000')
            .setTitle('Updated ' + member.ign + '\'s profile!')
            .setDescription('You got the role <@&' + role + '> because your index is **' + index + '**')
            .setFooter({ text: desc });

        message.channel.send({ embeds: [roles] });
    }
}