import { hypixel_api_key, desc, client } from "../index";
import fetch from "node-fetch";
import verifiedUsers from '../schemas/verified-users'
import { getApiData } from '../counters/get-api-data'
import { getIndex } from '../counters/index-count'
import { Message, MessageEmbed } from "discord.js";

module.exports = async client => {
  const guild = client.guilds.cache.get('897642715368534027')

  async function getUuid() {
    return await verifiedUsers.distinct('uuid').exec()
  }

  var allUuids = await getUuid()

  async function updateAccounts() {
    for (let i = 0; i < allUuids.length; i++) {
      const data = await getApiData(allUuids[i]) as any
      const indexes = await getIndex(data) as any

      await verifiedUsers.findOneAndUpdate(
        {
          uuid: (await data).player?.uuid,
        },
        {
          ign: (await data).player?.displayname,
          stats: {
            duelswins: (await data).player?.stats?.Duels?.wins || 0,
            duelsdeaths: (await data).player?.stats?.Duels?.deaths || 0,
            duelskills: (await data).player?.stats?.Duels?.kills || 0,
            bridgewins: (await data).player?.achievements?.duels_bridge_wins || 0,
            bedwarsfinals: (await data).player?.stats?.Bedwars?.final_kills_bedwars || 0,
            bedwarsstars: (await data).player?.achievements?.bedwars_level || 0,
          },
        },
      )
    }
  }

  /*
  async function updateGuildXP() {
    var allMembers = await verifiedUsers.distinct('uuid')
    for (let i = 0; i < allMembers.length; i++) {
      let guildObject = await (getGuildApiData as any).guild.members
      let guildMember = await guildObject.find(x => x.uuid === await allMembers[i].uuid).foo;
      let recentXP = await guildMember.expHistory[0]
      console.log(await recentXP);
    }
  }
  updateGuildXP()*/

  var now = new Date() as any
  var millisTill10 = new (Date as any)(now.getFullYear(), now.getMonth(), now.getDate(), 0, 33, 0, 0) - now;
  if (millisTill10 < 0) {
    millisTill10 += 86400000;
  }
  console.log(millisTill10 + ' : ' + new Date().toLocaleString())
  setTimeout(function() {
    updateAccounts()
    refreshDuelsWins()
    setTimeout(() => {
      refreshDuelsDeaths()
    }, 10000)
    resetLookingToPlay()
    //updateGuildXP()
    console.log('New cycle completed: ' + now)
  }, millisTill10);

  async function resetLookingToPlay() {
    let ltpChannel = client.channels.cache.get('998259961102598205')
    await ltpChannel.messages.fetch({ limit: 100 }).then(messages => {
      ltpChannel.bulkDelete(messages)
    })
    const ltp = new MessageEmbed()
      .setColor('#000000')
      .setTitle('You\'re looking to play with someone else?')
      .setDescription('Type ``@<Mode> <submode (like 1v1, 2v2 or 3v3)>, <additional comments>`` to start a queue. Please make sure you write the comments between the arguments. All the modes you can use are listed below.')
      .addFields(
        { name: 'Scrims', value: '<@&998242097092104242> Ask the members to play scrims with you! Type what mode (e.g. 1v1 bridge, 2v2 uhc or 4v4 bedwars) in submode.', inline: true },
        { name: 'The Bridge', value: '<@&998242220920553473> Ask the members to play bridge with you! Specify if you wanna play 1v1, 2v2 or 3v3 in submode.', inline: true },
        { name: 'BedWars', value: '<@&998242295268790353> Ask the members to play BedWars with you! Specify if you wanna play 1x8, 2x8, Dream Games etc. in submode. You might wanna say if what strategy you want to use (like iyn) in additional comments.', inline: true },
        { name: 'WoolWars', value: '<@&998242373257678968> Ask the members to play WoolWars with you!', inline: true },
    )
      .setFooter({ text: desc });

    ltpChannel.send({ embeds: [ltp] });
  }
  
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
  refreshAllRanks()

  async function refreshDuelsWins() {
    var totalWins = await verifiedUsers.distinct('stats.duelswins').exec()
    const channel = guild.channels.cache.get('960119541512417350')
    totalWins = totalWins.map((x) => +x);
    var totalWinsFinal = totalWins.reduce((partialSum, a) => partialSum + a, 0);
    channel.setName('Duels Wins: ' + totalWinsFinal.toLocaleString())
    //console.log('Refreshed totalWinsFinal: ' + totalWinsFinal)
  }

  async function refreshDuelsDeaths() {
    var totalDeaths = await verifiedUsers.distinct('stats.bridgewins').exec()
    const channel = guild.channels.cache.get('977677180697989130')
    totalDeaths = totalDeaths.map((x) => +x);
    var totalDeathsFinal = totalDeaths.reduce((partialSum, a) => partialSum + a, 0);
    channel.setName('Bridge Wins: ' + totalDeathsFinal.toLocaleString())
    //console.log('Refreshed totalDeathsFinal: ' + totalDeathsFinal)
  }

  updateAccounts()
  refreshDuelsWins()
  setTimeout(() => {
    refreshDuelsDeaths()
  }, 10000)
  //refreshAllRanks()
}
